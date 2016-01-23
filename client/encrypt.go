package main

import (
	//	"archive/zip"
	"fmt"
	"io/ioutil"
	"path/filepath"

	//	"golang.org/x/crypto/openpgp"
	//	"golang.org/x/crypto/openpgp/armor"
	//	"golang.org/x/crypto/openpgp/packet"
	//	"golang.org/x/crypto/ssh/terminal"
)

// LoadKey loads the key associated with this username,
// first by loooking at ~/.sendto/users/recipient/key.pub
// or if that fails by fetching it from the internet and saving at that location
func LoadKey(recipient string) (string, error) {
	fmt.Printf("Loading key for %s...\n", recipient)

	// For the moment as a test, use keybase.io, should be using our server
	keyURL := fmt.Sprintf("https://keybase.io/%s/key.asc", recipient)
	keyPath := filepath.Join(configPath(), "users", recipient, "key.pub")

	// Check if the key file exists at ~/.sendto/users/recipient/key.pub
	if !fileExists(keyPath) {
		// Make the enclosing dir
		createFolder(filepath.Join("users", recipient))

		// Fetch the key from our server
		err := DownloadData(keyURL, keyPath)
		if err != nil {
			return "", err
		}

		// Print the key for the user as we have fetched it for the first time?
		fmt.Printf("Fetched key for user:%s from:%s\n", recipient, keyURL)
	}

	// Load the key into memory - we might not need to do this
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return "", err
	}

	return string(key), nil
}

// EncryptFiles zips and encrypts our arguments (files or folders) using a public key
func EncryptFiles(args []string, key string) (string, error) {

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
	return "", nil
}

// DecryptFiles decrypts and unzips a file using a private key
// and returns the path of the resulting file/folder on success
func DecryptFiles(p string, key string) (string, error) {

	return "", nil
}
