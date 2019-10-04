#!/bin/bash
(
  all_commited=$(git diff --exit-code && git diff --cached --exit-code)

  if [ -n "$all_commited" ] || [ "$(git rev-parse @{u})" != "$(git rev-parse HEAD)" ]; then
      echo "Not all changes were commited and pushed"
      exit 1
  fi

  new_tag=$(grep -e '^## [0-9\.]*' CHANGELOG.md | awk '{print $2}')
  release_title="v${new_tag} release"

  read -r -p "Release title: ${release_title}? [y/N] " response
  if [[ "$response" =~ ^([yY][eE][sS]|[yY])+$ ]]; then
    git tag -a "v${new_tag}" -m "${release_title}"
    git push origin "v#${new_tag}"
  fi
)
