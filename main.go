package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Devices map[string]DeviceConfig `yaml:"devices"`
}
type DeviceConfig struct {
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Host     string `yaml:"host"`
    Port     string `yaml:"port"`
    Notes    string `yaml:"notes"`
}

type Args struct {
	addNew bool
	listDevices bool
}

var configFilePath string

func main() {

	addNewDevice := flag.Bool("add", false, "Add New Device to the Config")
	listDevices := flag.Bool("list", false, "List Configured Devices")

	flag.Parse()

	configArgs := Args {
		addNew:		*addNewDevice,
		listDevices: *listDevices,
	}

	// Get the directory where the executable is located
    exePath, err := os.Executable()
    if err != nil {
        log.Fatalf("Failed to get executable path: %s", err)
    }
    exeDir := filepath.Dir(exePath)
    configFilePath = filepath.Join(exeDir, "config.yaml")

	// configFilePath := "c:/utils/config.yaml"
    if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
        // log.Fatalf("Failed to read config file: %s", err)
		fmt.Println("Failed to read the file creating a blank one")
		emptyConfig := Config{Devices: make(map[string]DeviceConfig)}
        emptyConfigData, err := yaml.Marshal(&emptyConfig)
		_ = err
		err = os.WriteFile(configFilePath, emptyConfigData, 0644)
        if err != nil {
            log.Fatalf("Failed to create config file: %s", err)
        }
		os.Exit(2)
    }
	// Read the YAML configuration file
    configFile, err := os.ReadFile(configFilePath)
    if err != nil {
        log.Fatalf("Failed to read config file: %s", err)
    }
    // Unmarshal the YAML file into a Config struct
    var config Config
    err = yaml.Unmarshal(configFile, &config)
    if err != nil {
        log.Fatalf("Failed to unmarshal config file: %s", err)
		
    }

	//handle adding new device
	if configArgs.addNew {
		addNewDeviceToConfig(config)
		os.Exit(3)
	}
	if configArgs.listDevices {
		listConfiguredDevices(config)
		os.Exit(1)
	}
	//positional arguments
    if len(os.Args) < 2 && !configArgs.addNew {
        log.Fatalf("Usage: %s <device_key>", os.Args[0])
    }

    deviceKey := os.Args[1]


	openConnection(deviceKey, config)
    
}
func listConfiguredDevices(config Config)  {
	fmt.Println("KEY: Desciption")
	
	for index, device := range config.Devices{
		// fmt.Println(config.Devices[device])
		fmt.Printf("%s: %s \n", index, device.Notes)
	}
}

func addNewDeviceToConfig(config Config) {
	// var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(config)
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
	config.Devices[key] = DeviceConfig{
		Username: username,
		Password: password,
		Host:     hostname,
		Port:     port,
		Notes:    description,
	}
	saveToFile(config)
}


func saveToFile(config Config) {
	// Marshal the config back to YAML
	newConfigData, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}

	// Write the updated config back to the file
	err = os.WriteFile(configFilePath, newConfigData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("New device added to the config.")
	os.Exit(1) //exit application
}

func openConnection(deviceKey string, config Config) {
	// Lookup the device configuration by key
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