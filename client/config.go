package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// Config holds our client config
var Config map[string]string

// LoadConfig reads or creates our config file at ~/.sendto/config
func LoadConfig() error {

	// Create our files folder to store files to send
	err := createFolder("files")
	if err != nil {
		return err
	}

	// Create our users folder to store public keys downloaded
	err = createFolder("users")
	if err != nil {
		return err
	}

	// Load our config file (or create a new one with one entry - sender identity)
	// First check it exists
	file, err := ioutil.ReadFile(configFilePath())
	if err == nil {
		err = json.Unmarshal(file, &Config)
		if err != nil {
			return err
		}
	}

	// If no config exists create a config and save it
	if len(Config) == 0 {
		err = setupConfig()
		if err != nil {
			return err
		}

		err = SaveConfig()
		if err != nil {
			return err
		}

	}

	return nil
}

// SaveConfig saves our config out to a file at ~/.sendto/config
func SaveConfig() error {
	// Write out a json file representing our config map
	configJSON, err := json.MarshalIndent(Config, "", "\t")
	if err != nil {
		return err
	}

	// Write the config json file
	err = ioutil.WriteFile(configFilePath(), configJSON, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func setupConfig() error {
	Config = make(map[string]string, 0)

	// Get hold of the current user details
	u, err := user.Current()
	if err != nil {
		return err
	}
	fmt.Printf("Setting default sender identity to:%s\n", u.Name)
	Config["sender"] = u.Name
	return nil
}

func configPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Error loading config: %s\n", err)
		return ""
	}
	return filepath.Join(usr.HomeDir, ".sendto")
}

func configFilePath() string {
	return filepath.Join(configPath(), "config.json")
}

func createFolder(name string) error {
	p := filepath.Join(configPath(), name)
	return os.MkdirAll(p, os.ModeDir|os.ModePerm)
}

// fileExists returns true if this file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
