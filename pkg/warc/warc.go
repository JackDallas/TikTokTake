package warc

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

//WARC :
type WARC struct {
	SaveDirName string
	WARCName    string
}

//MakeRequest : Makes a http.Request but with the wget binary set to warc all files
func (warc *WARC) MakeRequest(req *http.Request) error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}

	dir := path.Dir(ex)
	saveDir := path.Join(dir, "archive", warc.SaveDirName)
	savePath := path.Join(saveDir, warc.WARCName)

	os.MkdirAll(saveDir, 0740)

	cmd := exec.Command(
		"wget",
		"--delete-after",
		fmt.Sprintf("--warc-file=%s", savePath),
		buildWgetHeaders(req),
		req.URL.String())

	var out bytes.Buffer
	cmd.Stdout = &out

	var outErr bytes.Buffer
	cmd.Stderr = &outErr

	err = cmd.Run()
	if err != nil {
		log.Errorf("Error running wget: %s, stderr: %s", err.Error(), outErr.String())
		return err
	}

	return nil
}

func buildWgetHeaders(req *http.Request) string {
	if len(req.Header) == 0 {
		return ""
	}

	var headers []string
	for k, v := range req.Header {
		headers = append(headers, fmt.Sprintf("--header='%s: %s'", k, v))
	}
	return strings.Join(headers, " ")
}
