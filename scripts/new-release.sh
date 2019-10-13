#!/bin/bash

promtConfirmation() {
  read -r -p "$1" response
  if ! [[ "$response" =~ ^([yY][eE][sS]|[yY])+$ ]]; then
    exit 1
  fi
}

(
  all_commited=$(git diff --exit-code && git diff --cached --exit-code)

  if [ "$(git rev-parse --abbrev-ref HEAD)" != "master" ]; then
    echo "Your current branch is not master."
    promtConfirmation "Are you sure you want to continnue? [y/N] "
  elif [ "$(git rev-parse @{u})" != "$(git rev-parse HEAD)" ]; then
    echo "You did not push all commits."
    promtConfirmation "Are you sure you want to continnue? [y/N] "
  elif [ -n "$all_commited" ]; then
    echo "You did not commit all changes."
    promtConfirmation "Are you sure you want to continnue? [y/N] "
  fi

  new_tag=$(grep -e '^## [0-9\.]*' CHANGELOG.md | awk '{print $2}' | head -n 1)
  release_title="v${new_tag} release"

  promtConfirmation "Release title: ${release_title}? [y/N] " # TODO: change title to description
  git tag -a "v${new_tag}" -m "${release_title}"
  git push origin "v${new_tag}"

)
