/*
Copyright © 2024 Victor Payno
*/
package main

import (
	"fmt"

	_ "embed"

	cli "github.com/vpayno/pomodoro-timer-go/internal/pomodoro-cli"
)

//go:generate bash ../../scripts/go-generate-helper-git-version-info
//go:embed .version.txt
var version []byte

func init() {
	cli.SetVersion(version)
}

func main() {
	fmt.Println("in cli main()")
	cli.Execute()
}
