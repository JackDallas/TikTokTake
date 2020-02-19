package download

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

// Video :
func Video(url, fileName, saveLoc string) error {
	// Create the directory to store the video
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dir := path.Dir(ex)
	saveDir := path.Join(dir, "archive", saveLoc)
	os.MkdirAll(saveDir, 0740)

	saveLoc = path.Join(saveDir, fileName+".mp4")

	// Download the video
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(saveLoc, bodyBytes, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
