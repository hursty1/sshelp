package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hursty1/sshelp/internal/filemanager"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a device to start a ssh connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := cmd.Context().Value("config").(filemanager.Config)

		// addNewDeviceToConfig(cfg)
		selectedDevice, err := selectDevice(cfg)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		openConnection(selectedDevice, cfg)
		return nil
	},
}

func openConnection(deviceKey string, config filemanager.Config){
	deviceConfig, ok := config.Devices[deviceKey]
    if !ok {
        log.Fatalf("Device key %s not found in config", deviceKey)
    }
    
	// Construct the SSH command
    sshCommand := fmt.Sprintf("ssh %s@%s -p %s", deviceConfig.Username, deviceConfig.Host, deviceConfig.Port)

    // Execute the command using PowerShell 
    cmd := exec.Command("powershell", "-Command", sshCommand)

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin

    fmt.Printf("Password Reminder: %s \n", deviceConfig.Password)
    if err := cmd.Run(); err != nil {
        log.Fatalf("Failed to execute command: %s", err)
    }
}
func selectDevice(cfg filemanager.Config) (string, error) {
	if len(cfg.Devices) == 0 {
		fmt.Println("No Devices configured.")
		return "", fmt.Errorf("No Devices Configured.")
	}

	keys := make([]string, 0, len(cfg.Devices))
	for k := range cfg.Devices {
		keys = append(keys,k)
	}

	prompt := promptui.Select{
		Label: "Select a device to open connection",
		Items: keys,
		Size: 10,
	}

	_, key, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return key, nil
}