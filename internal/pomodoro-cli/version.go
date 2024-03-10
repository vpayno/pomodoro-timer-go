/*
Copyright © 2024 Victor Payno
*/
package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type appInfo struct {
	name      string
	version   string
	gitHash   string
	buildTime string
}

var metadata = appInfo{
	name:    "pomodoro-cli",
	version: "0.0.0",
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show application version.",
	Long:  `Show application version.`,

	Run: func(_ *cobra.Command, _ []string) {
		showVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolP("verbose", "v", false, "Show verbose version")
}

// SetVersion is used my the main package to pass version information to the app package.
func SetVersion(b []byte) {
	slice := strings.Split(string(b), "\n")
	slice = slice[:len(slice)-1]

	if slice[0] != "" {
		metadata.version = slice[0]
	}

	if len(slice) > 1 {
		if slice[1] != "" {
			metadata.gitHash = slice[1]
		}

		if slice[2] != "" {
			metadata.buildTime = slice[2]
		}
	}
}

// Shows application version
func showVersion() {
	fmt.Println()
	fmt.Printf("%s Version: %s\n\n", metadata.name, metadata.version)

	if metadata.gitHash != "" {
		fmt.Printf("\tgit hash: %s\n", metadata.gitHash)
	}

	if metadata.buildTime != "" {
		fmt.Printf(" build time: %s\n", metadata.buildTime)
	}
}
