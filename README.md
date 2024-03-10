# pomodoro-timer-go

Pomodoro timer app consiting of a cli and a server.
The server component will allow me to share a single timer with multiple terminals and hosts.

## Notes

- You can use [RunMe](https://github.com/stateful/runme) to use this readme as a playbook.

## Setup

To install the git helpers and project dependencies run the following commands:

```bash { background=false category=setup closeTerminalOnSuccess=false excludeFromRunAll=true interactive=true interpreter=bash name=install-cli promptEnv=true terminalRows=10 }
set -ex

go install github.com/vpayno/pomodoro-timer-go/cmd/pomodoro-cli

pomodoro-cli version
```

## Releases

The `./tag-release` script is used to

- update the [CHANGELOG](./CHANGELOG.md)
- update .version.txt files
- create an annotated tag
- create a GitHub and/or GitLab release

Use this Runme playbook to list the latest 10 releases:

```bash { background=false category=release closeTerminalOnSuccess=false excludeFromRunAll=true interactive=true interpreter=bash name=releases-list promptEnv=true terminalRows=10 }
printf "\n"
printf "Latest releases:\n"
printf "\n"
git tag --list -n1 | tail
printf "\n"
```

Use this Runme playbook to list the unreleased commits:

```bash { background=false category=release closeTerminalOnSuccess=false excludeFromRunAll=true interactive=true interpreter=bash name=releases-unreleased-commits promptEnv=true terminalRows=10 }
printf "\n"
printf "Unreleased commits since %s:\n" "$(git tag --list -n0 | tail -n1)"
printf "\n"
git log --pretty=format:'%h -%d %s (%cr) <%an>' --abbrev-commit --decorate "$(git tag --list -n0 | tail -n1)"..
printf "\n"
```

Use this Runme playbook to tag a new release.

```bash { background=false category=release closeTerminalOnSuccess=false excludeFromRunAll=true interactive=true interpreter=bash name=release-create promptEnv=true terminalRows=20 }
export TAG_VER="x.y.z"
export TAG_TITLE="short description"

reset
./tag-release "${TAG_VER}" "${TAG_TITLE}"
```
