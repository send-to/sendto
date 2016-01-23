package main

import (
	//	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	//	"golang.org/x/crypto/openpgp"
	//	"golang.org/x/crypto/openpgp/armor"
	//	"golang.org/x/crypto/openpgp/packet"
	//	"golang.org/x/crypto/ssh/terminal"
)

const (
	v = "0.1"
)

var config map[string]string

func main() {
	command := ""
	args := os.Args[1:] // remove app path from args

	// We expect either a username or a subcommand and then a set of files in args
	if len(args) > 0 {
		command = args[0]
		args = args[1:]
	}

	// Load our configuration
	err := loadConfig()
	if err != nil {
		log.Fatalf("Sorry, an error occurred:\n\t%s", err)
	}

	switch command {
	case "encrypt", "e":
		err = Encrypt(args)
	case "decrypt", "d":
		err = Decrypt(args)
	case "identity", "i":
		err = Identity(args)
	case "version", "v":
		Version()
	case "help", "h":
		Help()
	default:
		// Default action is to send to (if we have a username and files)
		if len(args) > 0 {
			err = SendTo(command, args)
		} else {
			Help()
		}
	}

	if err != nil {
		log.Fatalf("Sorry, an error occurred:\n\t%s", err)
	}
}

// Version prints the version of this app
func Version() {
	fmt.Printf("\n\t-----\n\tSend to client - version:%s\n\t-----\n", v)
}

// Usage returns standard usage as a string
func Usage() string {
	return fmt.Sprintf("\tUsage: sendto kennygrant [files] - send files to the username kennygrant\n")
}

// Help prints the usage and commands
func Help() {
	Version()
	fmt.Printf(Usage())
	fmt.Printf("\t-----\n")
	fmt.Printf("\tCommands:\n")
	fmt.Printf("\tsendto version - display version\n")
	fmt.Printf("\tsendto [username] [files] - encrypt files for a given user\n")
	fmt.Printf("\tsendto encrypt [file] - encrypt a file\n")
	//	fmt.Printf("\tsendto decrypt [file] - decrypt a file\n")
	fmt.Printf("\tsendto identity [name] - sets default sender identity\n\n")
}

// Decrypt files specified, using the user's private key
// TODO: to support decryption we'd need access to private keys, perhaps leave this for hackathon
func Decrypt(args []string) error {
	log.Printf("Sorry, this client does not yet support decrypt")

	return nil
}

// Encrypt the files specified
func Encrypt(args []string) error {

	log.Printf("Sorry, this client does not yet support encryption")
	return nil
}

// SendTo sends files held in args to recipient
func SendTo(recipient string, args []string) error {

	// We expect at least 1 file to send
	if len(args) < 1 {
		return fmt.Errorf("Not enough arguments - %s", Usage())
	}

	// Notify the user that we're starting to send
	files := "file"
	if len(args) > 1 {
		files = "files"
	}
	fmt.Printf("Sending %d %s to %s...\n", len(args), files, recipient)

	// First check if we have our recipient's key on hand
	key, err := loadKey(recipient)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded key:%s", key)

	/*
		// Do something like this?
			// prepare to encrypt our data
			encoded, err := openpgp.Encrypt(out, []*openpgp.Entity{to}, from, hints, nil)
			if err != nil {
				return err
			}

			// Add files to a zip
			zipper := zip.NewWriter(encOut)

			t1, err := zipper.Create("test/test.txt")
			err := zipper.Flush()
			if err != nil {
				return err
			}

			// close the encPipe to finish the process
			err := encOut.Close()
			if err != nil {
				return err
			}
			// Send the file
	*/
	return nil
}

// Identity sets the default sender identity (as opposed to username)
func Identity(args []string) error {
	log.Printf("Sorry, this client does not yet support setting identity")
	return nil
}

// loadKey loads the key associated with this username, if necessary fetching it from the internet
// for now we just look in the ~/.sendto/users folder
// don't use strings perhaps?
func loadKey(recipient string) (string, error) {
	fmt.Printf("Loading key for %s...\n", recipient)

	return "hello key", nil
}

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

	// If no config create a config and save it
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
