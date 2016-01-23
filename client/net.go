package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error %d downloading file at %s", resp.StatusCode, url)
	}

	// If response OK, create our file with downloaded response body
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	n, err := io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes to file %s\n", n, filePath)

	return nil
}

// PostData sends data to the server
func PostData(recipient string, filePath string, sender string) error {

	// Prepare a new multipart form writer
	var formData bytes.Buffer
	w := multipart.NewWriter(&formData)

	// Add the encrypted file
	err := addFile(w, filePath)
	if err != nil {
		return err
	}

	// Add sender identity
	err = addField(w, "from", sender)
	if err != nil {
		return err
	}
	// Add recipient identity
	err = addField(w, "to", recipient)
	if err != nil {
		return err
	}
	// Close the writer
	w.Close()

	// Now post the form
	//url := "https://sendto.click/files/upload"
	url := "http://localhost:3000/files/upload"
	req, err := http.NewRequest("POST", url, &formData)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error %d posting file to %s", resp.StatusCode, url)
	}

	fmt.Printf("POSTED: %s %v", url, resp)

	return nil

}

// addFile adds a file to this multipart form
func addFile(w *multipart.Writer, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	fw, err := w.CreateFormFile("file", filePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, f)
	if err != nil {
		return err
	}
	return nil
}

// addField adds a field to the multipart writer
func addField(w *multipart.Writer, k, v string) error {
	fw, err := w.CreateFormField(k)
	if err != nil {
		return err
	}
	_, err = fw.Write([]byte(v))
	if err != nil {
		return err
	}
	return nil
}

/*
formData := url.Values{
"from": []string{sender},
}

resp, err := http.PostForm(u, formData)
if err != nil {
return err
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
return fmt.Errorf("error %d posting file to %s", resp.StatusCode, u)
}

return nil
*/
