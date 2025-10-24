package cmd

import (
	"fmt"
	"log"
	"main/internal/filemanager"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a device config",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := cmd.Context().Value("config").(filemanager.Config)

		//
		deleteDeviceCommand(cfg)

		err := cfg.Save()
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}



func deleteDeviceCommand(cfg filemanager.Config) error {
	//
	if len(cfg.Devices) == 0 {
		fmt.Println("No Devices configured.")
		return fmt.Errorf("No Devices Configured.")
	}

	keys := make([]string, 0, len(cfg.Devices))
	for k := range cfg.Devices {
		keys = append(keys,k)
	}

	prompt := promptui.Select{
		Label: "Select a device to delete",
		Items: keys,
		Size: 10,
	}

	_, key, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	fmt.Printf("Deleting: %s\n", key)
	delete(cfg.Devices, key)
	if err := cfg.Save(); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("Device %s was deleted\n", key)
	
	return nil
}