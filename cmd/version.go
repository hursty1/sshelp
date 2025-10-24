package cmd

import (
	"fmt"

	version "github.com/hursty1/sshelp/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version of sshelp",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sshelp version:", version.Get())
		latest, err := LatestVersion()
		if err != nil {
			fmt.Println("update check failed:", err)
			// return false
		}
		fmt.Println("Latest Version is:", latest)
	},
}
