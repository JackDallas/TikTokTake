package tiktok

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	structs "github.com/JackDallas/TikTokTake/internal/structs"

	log "github.com/sirupsen/logrus"
)

//User : A TikTok user account
type User struct {
	Username       string
	Metadata       structs.TikTokMeta
	ProfileRequest *http.Request
}

//NewUser : Creates and validates a TikTok user account
func NewUser(username string) (User, error) {
	userMeta, req, err := getUserDetails(username)
	if err != nil {
		log.Errorf("Could not find user %s: %s", username, err.Error())
		return User{}, err
	}
	return User{username, userMeta, req}, nil
}

//TODO: see if this can be used on all HTML pages
func getUserDetails(username string) (structs.TikTokMeta, *http.Request, error) {
	if len(username) == 0 {
		return structs.TikTokMeta{}, nil, errors.New("empty username")
	}

	if username[0] == '@' {
		username = username[1:]
		log.Debug("Username starts with @ stripping to become : " + username)
	}

	client := &http.Client{}
	req, err := NewHTMLRequest("GET", "https://tiktok.com/@"+username)
	if err != nil {
		return structs.TikTokMeta{}, nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return structs.TikTokMeta{}, nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return structs.TikTokMeta{}, nil, err
		}

		re := regexp.MustCompile(`(?m)<script id="__NEXT_DATA__" type="application/json" crossorigin="anonymous">.+?</script>`)
		results := re.FindAllString(string(bodyBytes), -1)

		if len(results) > 0 {
			metaString := results[0]
			metaString = strings.Replace(metaString, "<script id=\"__NEXT_DATA__\" type=\"application/json\" crossorigin=\"anonymous\">", "", -1)
			metaString = strings.Replace(metaString, "</script>", "", -1)

			var meta structs.TikTokMeta
			json.Unmarshal([]byte(metaString), &meta)
			return meta, req, nil
		}
		return structs.TikTokMeta{}, nil, errors.New("No metadata found")
	}

	return structs.TikTokMeta{}, nil, errors.New("Non OK status code, User may not exist: " + err.Error())
}

// GetVideos : returns a list of video urls for a user
func (user *User) GetVideos() ([]*http.Request, error) {
	client := &http.Client{}
	var urlList []*http.Request
	referer := "https://tiktok.com/@" + user.Username

	maxCursor := "0"
	running := true

	for running {
		jsonURL := fmt.Sprintf("https://m.tiktok.com/share/item/list?secUid=%s&id=%s&type=1&count=30&minCursor=0&maxCursor=%s&shareUid=&lang=",
			user.Metadata.Props.PageProps.UserData.SecUID,
			user.Metadata.Props.PageProps.UserData.UserID,
			maxCursor)

		// the webtoken is unused for now but will be implemented later
		signature, err := GenerateSignature(jsonURL, "window."+user.Metadata.Query.Webtoken)
		if err != nil {
			return []*http.Request{}, err
		}

		jsonURL = jsonURL + "&_signature=" + signature
		jsonURL = strings.TrimSpace(jsonURL)

		req, err := NewJSONRequest(jsonURL, referer)
		if err != nil {
			log.Debug(err.Error())
			running = false
			break
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Debug(err.Error())
			running = false
			break
		}

		json, err := DecodeJSONRequest(resp)
		if err != nil {
			log.Debug(err.Error())
			running = false
			break
		}

		for _, item := range json.Body.ItemListData {
			if len(item.ItemInfos.Video.Urls) > 0 {
				nextURL, error := NewVideoRequest(item.ItemInfos.Video.Urls[0], referer)
				if err != nil {
					log.Errorf("Error making video request: %s", error.Error())
				} else {
					urlList = append(urlList, nextURL)
				}
			}
		}

		// Handle paging

		if len(json.Body.MaxCursor) == 0 {
			running = false
		} else if maxCursor != json.Body.MaxCursor {
			maxCursor = json.Body.MaxCursor
			log.Debug("Next max cursor: " + maxCursor)
		} else {
			running = false
		}
	}

	return urlList, nil
}
