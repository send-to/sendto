package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadData retrieves a file from the server and save contents at filePath
func DownloadData(url string, filePath string) error {

	// Fetch the file
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// If response OK, create our file with downloaded response body
	if resp.StatusCode == http.StatusOK {
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		n, err := io.Copy(f, resp.Body)
		if err != nil {
			return err
		}
		fmt.Printf("Wrote %d bytes to file %s\n", n, filePath)

	} else {
		return fmt.Errorf("error %d downloading file at %s", resp.StatusCode, url)
	}

	return nil
}

// PostData sends data to the server
func PostData(url string, filePath string) error {

	return nil
}
