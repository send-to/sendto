package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
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

	// Use the home dir as a default sender name
	name := path.Base(homePath())

	u, err := user.Current()
	// Recover gracefully from lack of user.Current() on other platforms
	// only set if this is supported
	if err == nil && u.Name != "" {
		name = u.Name
	}

	// Set up a default config, pointing to sendto.click
	Config["sender"] = name
	Config["keyserver"] = "https://sendto.click/users/%s/key"
	Config["server"] = "https://sendto.click"

	fmt.Printf("Setting default config:%v\n", Config)
	return nil
}

func configPath() string {
	return filepath.Join(homePath(), ".sendto")
}

func homePath() string {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}
	return home
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
