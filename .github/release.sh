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

for BASE in ${MODULES}; do
    if grep -q "_examples" <<<"${BASE}"; then
        continue
    fi

    MODULE_NAME="$(basename "${BASE}")"
    TAG_NAME="${MODULE_NAME}/${NEXT}"

    # go requires tags like cmd/tintgen/v… for submodule paths under cmd/
    if [[ ${BASE} =~ /cmd/[^/]+$ ]]; then
        CMD_PARENT="$(basename "$(dirname "${BASE}")")"
        TAG_NAME="${CMD_PARENT}/${MODULE_NAME}/${NEXT}"
    fi

    echo "--> tagging ${TAG_NAME}"

    if git ls-remote --tags origin | grep -q "/${TAG_NAME}$"; then
        echo "!!> ${TAG_NAME} already tagged on origin"
        continue
    fi
    git tag -m "${TAG_NAME}" "${TAG_NAME}" || exit 1
done

git push origin --tags
