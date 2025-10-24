package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hursty1/sshelp/internal/filemanager"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add New device to the config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := cmd.Context().Value("config").(filemanager.Config)

		addNewDeviceToConfig(cfg)

		err := cfg.Save()
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}


func addNewDeviceToConfig(config filemanager.Config) {
	// var err error
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println(config)
	fmt.Println("Please enter a Device Name")
	key, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// key = key[:len(key)-1] 
	key = strings.TrimSpace(key)
	//check if key exists already
    if _, ok := config.Devices[key]; ok {
        fmt.Println("Device key already exists in the config.")
        return
    } 
	fmt.Println("Enter a username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	username = strings.TrimSpace(username)

	fmt.Println("Enter a hostname: ")
	hostname, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	hostname = strings.TrimSpace(hostname)
	fmt.Println("Enter a password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	password = strings.TrimSpace(password)
	fmt.Println("Enter a port: ")
	port, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	port = strings.TrimSpace(port)

	fmt.Println("Enter a Description or Note: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	description = strings.TrimSpace(description)

	// Add the new device to the config
	
	newDevice := filemanager.DeviceConfig{
		Username: username,
		Password: password,
		Host:     hostname,
		Port:     port,
		Notes:    description,
	}

	config.AddDevice(key, newDevice)

}