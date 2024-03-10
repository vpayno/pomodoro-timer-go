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
