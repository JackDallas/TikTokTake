package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		scrapeUserVideos(args[0])
	} else {
		fmt.Println("Need 1 argument, url of account")
		os.Exit(1)
	}
}

func scrapeUserVideos(url string) {
	meta, err := getPageDetails(url)
	if err != nil {
		log.Fatal(err)
	}

	jsonURL := fmt.Sprintf("https://m.tiktok.com/share/item/list?secUid=%s&id=%s&type=1&count=30&minCursor=0&maxCursor=0&shareUid=&lang=", meta.Props.PageProps.UserData.SecUID, meta.Props.PageProps.UserData.UserID)
	// the webtoken is unused for now but will be implemented later
	signature, err := generateSignature(jsonURL, "window."+meta.Query.Webtoken)
	if err != nil {
		log.Fatal(err)
	}

	jsonURL = jsonURL + "&_signature=" + signature
	jsonURL = strings.TrimSpace(jsonURL)

	json, err := getPageJSON(jsonURL)
	if err != nil {
		log.Fatal(err)
	}

	var urlList []string
	for _, item := range json.Body.ItemListData {
		if len(item.ItemInfos.Video.Urls) > 0 {
			urlList = append(urlList, item.ItemInfos.Video.Urls[0])
		}
	}

	fmt.Println(urlList)
}

func generateSignature(url string, tac string) (string, error) {
	cmd := exec.Command("node", "libs/tiktok-signature/browser.js", url, tac)

	var errbuf bytes.Buffer
	cmd.Stderr = &errbuf

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(errbuf.String())
		return "", err
	}

	return string(out), nil
}

func getPageJSON(jsonURL string) (ListJSON, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", jsonURL, nil)
	if err != nil {
		return ListJSON{}, err
	}

	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	req.Header.Add("origin", "https://www.tiktok.com")
	req.Header.Add("referer", "https://www.tiktok.com/@one_eared_uno")

	resp, err := client.Do(req)

	if err != nil {
		return ListJSON{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ListJSON{}, err
		}

		var listJSON ListJSON
		json.Unmarshal(bodyBytes, &listJSON)

		return listJSON, nil
	}
	return ListJSON{}, errors.New("Error: Non OK status code")
}

func getPageDetails(url string) (TikTokMeta, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TikTokMeta{}, err
	}

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	resp, err := client.Do(req)

	if err != nil {
		return TikTokMeta{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return TikTokMeta{}, err
		}
		bodyString := string(bodyBytes)

		var re = regexp.MustCompile(`(?m)<script id="__NEXT_DATA__" type="application/json" crossorigin="anonymous">.+?</script>`)
		results := re.FindAllString(bodyString, -1)

		if len(results) > 0 {
			metaString := results[0]
			metaString = strings.Replace(metaString, "<script id=\"__NEXT_DATA__\" type=\"application/json\" crossorigin=\"anonymous\">", "", -1)
			metaString = strings.Replace(metaString, "</script>", "", -1)

			var meta TikTokMeta
			json.Unmarshal([]byte(metaString), &meta)
			return meta, nil
		}
		return TikTokMeta{}, errors.New("No meta")
	}

	return TikTokMeta{}, errors.New("Error: Non OK status code")
}
