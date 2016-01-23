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

	// We expect a subcommand and a set of files in args
	if len(args) > 0 {
		command = args[1]
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
		version(args)
	case "help", "h":
		help(args)
	default:
		help(args)
	}

	if err != nil {
		log.Fatalf("Sorry, an error occurred:\n\t%s", err)
	}
}

func version(args []string) error {
	fmt.Printf("\n\t-----\n\tSend to client - version:%s\n\t-----\n", v)
	return nil
}

func help(args []string) error {
	version(args)
	fmt.Printf("\tsendto version - display version\n")
	fmt.Printf("\tsendto [username] [files] - encrypt files for a given user\n")
	fmt.Printf("\tsendto decrypt [file] - decrypt a file\n")
	fmt.Printf("\tsendto identity [name] - sets default sender identity\n\n")
	return nil
}

func decrypt(args []string) error {
	log.Printf("Sorry, this client does not yet support decrypt")
	return nil
}

func encrypt(args []string) error {
	log.Printf("Sorry, this client does not yet support encryption")
	return nil
}

func identity(args []string) error {
	log.Printf("Sorry, this client does not yet support setting identity")
	return nil
}
