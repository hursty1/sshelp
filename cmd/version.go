package cmd

import (
	"fmt"

	version "github.com/hursty1/ssh_tool"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version of sshelp",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sshelp version:", version.Version)
	},
}
