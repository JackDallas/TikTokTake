package httptiktok

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	structs "github.com/JackDallas/TikTokTake/internal/structs"
)

//MakeTikTokRequest : Makes a request to TikTok with our preset headers
func MakeTikTokRequest(method string, url string) (*http.Response, error) {
	client := &http.Client{}

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

	return client.Do(req)
}

// GetTikTokJSONPage : Returns a ListJSON object (URL must have signature already!)
// TODO: Return only needed info, make more generic
func GetTikTokJSONPage(jsonURL string, refererURL string) (structs.ListJSON, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", jsonURL, nil)
	if err != nil {
		return structs.ListJSON{}, err
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

	resp, err := client.Do(req)

	if err != nil {
		return structs.ListJSON{}, err
	}

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
