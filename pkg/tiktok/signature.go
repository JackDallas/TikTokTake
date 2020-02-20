package tiktok

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"

	log "github.com/sirupsen/logrus"
)

// GenerateSignature : Generates a TikTok Signature for use in requests
func GenerateSignature(url string, tac string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		log.Error(err)
		return "", err
	}

	dir := path.Dir(ex)
	jsLoc := path.Join(dir, "/libs/tiktok-signature/browser.js")
	//TODO replace this
	cmd := exec.Command("node", jsLoc, url, tac)

	var errbuf bytes.Buffer
	cmd.Stderr = &errbuf

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(errbuf.String())
		return "", err
	}

	return string(out), nil
}
