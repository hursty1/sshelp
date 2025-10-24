package cmd

import (
	"fmt"
	"main/internal/filemanager"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the configured hosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := cmd.Context().Value("config").(filemanager.Config)
	
		listConfiguredDevices(cfg)
		return nil
	},
}

var p = fmt.Println
var pf = fmt.Printf

func listConfiguredDevices(cfg filemanager.Config) {
	if len(cfg.Devices) < 1 {
		p("No devices are currently configured.")
	}
	p("Listing all Configured Devices\n")
	for key, device := range cfg.Devices {
		pf("Key: %v\n", key)
		pf("Host: %v\n", device.Host)
		pf("Notes: %v\n\n", device.Notes)

	}
}