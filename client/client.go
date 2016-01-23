package main

import (
	"fmt"
	"log"
	"os"
)

const (
	v = "0.1"
)

func main() {
	command := ""
	args := os.Args[1:] // remove app path from args

	// We expect either a username or a subcommand and then a set of files in args
	if len(args) > 0 {
		command = args[0]
		args = args[1:]
	}

	var err error
	switch command {
	case "encrypt", "e":
		err = encrypt(args)
	case "decrypt", "d":
		err = decrypt(args)
	case "identity", "i":
		err = identity(args)
	case "version", "v":
		version()
	case "help", "h":
		help()
	default:
		// Default action is to send to (if we have a username and files)
		if len(args) > 0 {
			err = sendTo(command, args)
		} else {
			help()
		}
	}

	if err != nil {
		log.Fatalf("Sorry, an error occurred:\n\t%s", err)
	}
}

func version() {
	fmt.Printf("\n\t-----\n\tSend to client - version:%s\n\t-----\n", v)
}

func usage() string {
	return fmt.Sprintf("\tUsage: sendto kennygrant [files] - send files to the username kennygrant\n")
}

func help() {
	version()
	fmt.Printf(usage())
	fmt.Printf("\t-----\n")
	fmt.Printf("\tCommands:\n")
	fmt.Printf("\tsendto version - display version\n")
	fmt.Printf("\tsendto [username] [files] - encrypt files for a given user\n")
	fmt.Printf("\tsendto encrypt [file] - encrypt a file\n")
	//	fmt.Printf("\tsendto decrypt [file] - decrypt a file\n")
	fmt.Printf("\tsendto identity [name] - sets default sender identity\n\n")
}

// decrypt files specified, using the user's private key
// TODO: to support decryption we'd need access to private keys, perhaps leave this for hackathon
func decrypt(args []string) error {
	log.Printf("Sorry, this client does not yet support decrypt")

	return nil
}

// encrypt the paths specified
func encrypt(args []string) error {

	log.Printf("Sorry, this client does not yet support encryption")
	return nil
}

// The main path - send files held in args to recipient
func sendTo(recipient string, args []string) error {
	fmt.Printf("\nSending %d files to %s...\n\n", len(args), recipient)

	if len(args) < 1 {
		return fmt.Errorf("Not enough arguments - %s", usage())
	}

	return nil
}

func identity(args []string) error {
	log.Printf("Sorry, this client does not yet support setting identity")
	return nil
}
