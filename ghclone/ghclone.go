package ghclone

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Repo struct {
	Name string
	FullName string `json:"full_name"`
	URL string `json:"html_url"`
	SSHURL string `json:"ssh_url"`
	HTTPSURL string `json:"clone_url"`
	SizeKB int64 `json:"size"`
}

func FetchRepos(accessToken string, private bool) ([]Repo, error) {
	privateString := ""
	if !private {
		privateString = "&type=public"
	}

	next := fmt.Sprintf("https://api.github.com/user/repos?access_token=%s%s", accessToken, privateString)
	repos := make([]Repo, 0)
	for next != "" {
		resp, err := http.Get(next)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		var newRepos []Repo
		err = json.Unmarshal(body, &newRepos)
		if err != nil {
			return nil, err
		}
		repos = append(repos, newRepos...)

		next, err = nextLink(resp.Header)
		if err != nil {
			return nil, err
		}
	}

	return repos, nil
}

func nextLink(header http.Header) (string, error) {
	links := strings.Split(header.Get("Link"), ",")
	next := ""
	for _, link := range links {
		parts := strings.Split(link, ";")
		if len(parts) != 2 {
			return "", errors.New("link parts not equal to 2")
		}

		if strings.TrimSpace(parts[1]) == "rel=\"next\"" {
			parts[0] = strings.TrimSpace(parts[0])
			if len(parts[0]) == 0 {
				return "", errors.New("empty next link")
			}
			next = parts[0][1:len(parts[0]) - 1]
		}
	}
	return next, nil
}
