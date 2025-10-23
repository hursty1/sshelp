package filemanager

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
var appDir = "sshelp"


func create_or_get_file_path() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Failed to get users config directory %v", err)
		return "", err
	}
	appDir := filepath.Join(configDir, appDir)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		log.Fatalf("Failed to create config Directory: %v", err)
		return "", err
	}
	configFilePath := filepath.Join(appDir, "config.yaml")
	
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		fmt.Println("Config not found. Creating a blank one at:", configFilePath)
		emptyConfig := Config{Devices: make(map[string]DeviceConfig)}
		emptyData, _ := yaml.Marshal(&emptyConfig)
		err = os.WriteFile(configFilePath, emptyData, 0644)
		if err != nil {
			log.Fatalf("Failed to create config file: %v", err)
		}
		os.Exit(2)
	}
	return configFilePath, nil
}

func (c *Config) Read() error {
	configFilePath, err := create_or_get_file_path()
	if err != nil {
		return err
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}


	err = yaml.Unmarshal(configFile, c)
	if err != nil {
		log.Fatalf("Unable to unmarshal configuration file: %v", err)
	}
	return nil
}

func (c *Config) Save() error {
	path, err := create_or_get_file_path()
    if err != nil {
        return err
    }
	data, err := yaml.Marshal(c)
    if err != nil {
        return err
    }
	return os.WriteFile(path, data, 0644)
}

func (c *Config) AddDevice(key string, device DeviceConfig) error {

	if _, ok := c.Devices[key]; ok {
		fmt.Printf("Device Key: %s, already exists", key)
		return fmt.Errorf("Device Key Already Exists")
	}
	
	c.Devices[key] = device

	return nil
}

//update

//delete