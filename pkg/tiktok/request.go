package tiktok

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	structs "github.com/JackDallas/TikTokTake/internal/structs"
)

//NewHTMLRequest : Creates a request to a TikTok HTML endpoint with preset tiktok headers
func NewHTMLRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	return req, nil
}

//NewJSONRequest : Creates a request to a TikTok JSON endpoint (URL must have signature already!)
func NewJSONRequest(jsonURL string, refererURL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", jsonURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	req.Header.Add("origin", "https://www.tiktok.com")
	req.Header.Add("referer", refererURL)

	return req, nil
}

//NewVideoRequest : Creates a request to a TikTok video endpoint
func NewVideoRequest(url string, refererURL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "video/webm,video/ogg,video/*;q=0.9,application/ogg;q=0.7,audio/*;q=0.6,*/*;q=0.5")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("connection", "keep-alive")
	req.Header.Add("host", req.URL.Host)
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("range", "bytes=0-")
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	req.Header.Add("referer", refererURL)

	return req, nil
}

//DecodeJSONRequest : Decodes a JSON response from TikTok into a struct
func DecodeJSONRequest(resp *http.Response) (structs.ListJSON, error) {
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return structs.ListJSON{}, err
		}

		var listJSON structs.ListJSON
		json.Unmarshal(bodyBytes, &listJSON)

		return listJSON, nil
	}
	return structs.ListJSON{}, errors.New("Error: Non OK status code")
}
