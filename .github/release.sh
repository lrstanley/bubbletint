#!/usr/bin/env bash

# release helper: tags the repo and bumps downstream go.mod requires
# in dependency order. each module is tagged before the next one is tidied,
# because glamourtint requires the chromatint tag it just published.
set -euo pipefail

# go mod edit and go mod tidy must follow each go.mod, not go.work
export GOWORK=off

# tidy runs immediately after a tag push, before the module proxy has the tag
export GOPRIVATE="github.com/lrstanley/bubbletint"

BASE_MODULE="github.com/lrstanley/bubbletint/v2"

# dependency order. do not include _examples: it uses replace directives
DOWNSTREAM=(
	chromatint
	glamourtint
)

if ! command -v svu >/dev/null 2>&1; then
	echo "svu is not installed"
	exit 1
fi

if [ "$(git rev-parse --show-toplevel)" != "$PWD" ]; then
	echo "must be run from the root of the repository"
	exit 1
fi

CURRENT="$(svu current --tag.pattern "v*")"
NEXT="${1:-}"
if [ -z "$NEXT" ]; then
	NEXT="$(svu patch --tag.pattern "v*")"
fi

if [ "$CURRENT" = "$NEXT" ]; then
	echo "${CURRENT} is already the latest version"
fi

echo "${CURRENT} -> ${NEXT}"
echo "press enter to continue"
read -r

tag_and_push() {
	local tag_name="$1"

	if ! git tag -l | grep -q "^${tag_name}$"; then
		echo "tagging ${tag_name}"
		git tag -m "${tag_name}" "${tag_name}"
	fi

	if ! git ls-remote --tags origin | grep -q "/${tag_name}$"; then
		echo "pushing ${tag_name}"
		git push origin "${tag_name}"
	fi
}

# deps_for prints the module paths dir should require at ${NEXT}.
deps_for() {
	case "$1" in
	chromatint)
		printf '%s\n' "$BASE_MODULE"
		;;
	glamourtint)
		printf '%s\n' \
			"$BASE_MODULE" \
			"github.com/lrstanley/bubbletint/chromatint/v2"
		;;
	*)
		echo "unknown downstream module ${1}" >&2
		return 1
		;;
	esac
}

bump_module() {
	local dir="$1"
	local gomod="${dir}/go.mod"
	local mod

	if [ ! -f "$gomod" ]; then
		echo "no go.mod at ${dir}" >&2
		return 1
	fi

	while IFS= read -r mod; do
		# require lines are indented. the module line is not, so a module
		# path is not treated as a requirement of itself.
		if ! grep -F "${mod} " "$gomod" | grep -Eq '^[[:space:]]'; then
			echo "${dir} does not require ${mod}" >&2
			return 1
		fi
		echo "bump ${mod} in ${gomod}"
		go mod edit -require="${mod}@${NEXT}" "$gomod"
	done < <(deps_for "$dir")

	pushd "$dir" >/dev/null
	go mod tidy
	popd >/dev/null

	git add "${dir}/go.mod" "${dir}/go.sum"
	if git diff --cached --quiet; then
		echo "no go.mod changes for ${dir}"
		return 0
	fi

	echo "committing ${dir}"
	git commit -m "chore(release): bump ${dir} to ${NEXT}"
	git push origin HEAD
}

tag_and_push "$NEXT"

for dir in "${DOWNSTREAM[@]}"; do
	bump_module "$dir"
	tag_and_push "${dir}/${NEXT}"
done
