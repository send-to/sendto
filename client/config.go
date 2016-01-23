package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var config map[string]string

// loadConfig reads or creates our config file at ~/.sendto/config
func loadConfig() error {

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
		err = json.Unmarshal(file, &config)
		if err != nil {
			return err
		}
	}

	// If no config exists create a config and save it
	if len(config) == 0 {
		err = setupConfig()
		if err != nil {
			return err
		}

		err = saveConfig()
		if err != nil {
			return err
		}

	}

	return nil
}

// saves our config out to a file at ~/.sendto/config
func saveConfig() error {
	// Write out a json file representing our config map
	configJSON, err := json.MarshalIndent(config, "", "\t")
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
	config = make(map[string]string, 0)

	// Get hold of the current user details
	u, err := user.Current()
	if err != nil {
		return err
	}
	fmt.Printf("Setting default sender identity to:%s\n", u.Name)
	config["sender"] = u.Name
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
