package main

import (
	//	"archive/zip"
	"fmt"

	//	"golang.org/x/crypto/openpgp"
	//	"golang.org/x/crypto/openpgp/armor"
	//	"golang.org/x/crypto/openpgp/packet"
	//	"golang.org/x/crypto/ssh/terminal"
)

// loadKey loads the key associated with this username, if necessary fetching it from the internet
// for now we just look in the ~/.sendto/users folder
func loadKey(recipient string) (string, error) {
	fmt.Printf("Loading key for %s...\n", recipient)

	return "hello key", nil
}

func encryptFiles() {
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
}
