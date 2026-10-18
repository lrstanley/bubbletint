#!/usr/bin/env bash

# release helper: tags the repo and bumps downstream go.mod requires
set -euo pipefail

export BASE_MODULE="github.com/lrstanley/bubbletint/v2"
# optional nested-module paths bumped to ${NEXT} when found in submodule go.mod
BUMP_REQUIRES_EXTRA=(
    "github.com/lrstanley/bubbletint/chromatint/v2"
)

if ! command -v svu >/dev/null 2>&1; then
    echo "svu is not installed"
    exit 1
fi

if [ "$(git rev-parse --show-toplevel)" != "$PWD" ]; then
    echo "--> must be run from the root of the repository"
    exit 1
fi

export CURRENT
CURRENT="$(svu current --tag.pattern "v*")"
export NEXT
NEXT="$1"
if [ -z "$NEXT" ]; then
    NEXT="$(svu patch --tag.pattern "v*")"
fi

if [ "$CURRENT" == "$NEXT" ]; then
    echo "!!> ${CURRENT} is already the latest version"
fi

echo "--> ${CURRENT} -> ${NEXT}"
echo "press enter to continue"
read -r

if ! git tag -l | grep -q "^${NEXT}$"; then
    echo "--> tagging ${NEXT}"
    git tag -m "$NEXT" "$NEXT"
fi

if ! git ls-remote --tags origin | grep -q "/${NEXT}$"; then
    echo "--> pushing ${NEXT}"
    git push origin "$NEXT"
fi

MODULES="$(find "$PWD" -mindepth 2 -maxdepth 3 -name "go.mod" -exec dirname "{}" \;)"

for BASE in ${MODULES}; do
    GOMOD="${BASE}/go.mod"
    UPDATED=false
    for MOD in "${BASE_MODULE}" "${BUMP_REQUIRES_EXTRA[@]}"; do
        if grep -Fq "${MOD}" "${GOMOD}"; then
            echo "--> bump ${MOD} in ${GOMOD}"
            go mod edit -require="${MOD}@${NEXT}" "${GOMOD}" || exit 1
            UPDATED=true
        fi
    done
    if [ "${UPDATED}" = false ]; then
        echo "--> skip ${BASE} (no base or extra bump require)"
        continue
    fi
    pushd "${BASE}" || exit 1
    go mod tidy || exit 1
    echo "--> staging ${BASE}/go.mod and ${BASE}/go.sum"
    git add go.mod go.sum || exit 1
    popd || exit 1
done

if ! git diff --cached --quiet >/dev/null; then
    echo "--> committing and pushing go.mod changes"
    git commit -m "chore(release): bump sub-modules to ${NEXT}"
    git push origin
fi

# submodule tags only for paths in bump_requires_extra (not every nested go.mod dir)
_repo_pkg_root="${BASE_MODULE%/*}"

for MOD in "${BUMP_REQUIRES_EXTRA[@]}"; do
    rel="${MOD#${_repo_pkg_root}/}"

    case "${MOD}" in
        "${_repo_pkg_root}"/*) ;;
        *)
            echo "!!> bump module ${MOD} is not under ${_repo_pkg_root}; skip tagging"
            continue
            ;;
    esac


    if [[ "${rel}" =~ ^(.+)/v[0-9]+$ ]]; then
        tag_prefix="${BASH_REMATCH[1]}"
    else
        tag_prefix="${rel}"
    fi

    tag_name="${tag_prefix}/${NEXT}"
    submodule_root="${PWD}/${tag_prefix}"

    if [ ! -f "${submodule_root}/go.mod" ]; then
        echo "!!> no go.mod at ${tag_prefix}; skip tagging ${tag_name}"
        continue
    fi

    echo "--> tagging ${tag_name}"

    if git ls-remote --tags origin | grep -q "/${tag_name}$"; then
        echo "!!> ${tag_name} already tagged on origin"
        continue
    fi
    git tag -m "${tag_name}" "${tag_name}" || exit 1
done

git push origin --tags
