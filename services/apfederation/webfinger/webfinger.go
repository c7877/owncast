package webfinger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Webfinger struct{}

var temporaryGlobalInstance *Webfinger

func Get() *Webfinger {
	if temporaryGlobalInstance == nil {
		temporaryGlobalInstance = New()
	}
	return temporaryGlobalInstance
}

func New() *Webfinger {
	return &Webfinger{}
}

// GetWebfingerLinks will return webfinger data for an account.
func (w *Webfinger) GetWebfingerLinks(account string) ([]map[string]interface{}, error) {
	type webfingerResponse struct {
		Links []map[string]interface{} `json:"links"`
	}

	account = strings.TrimLeft(account, "@") // remove any leading @
	accountComponents := strings.Split(account, "@")
	fediverseServer := accountComponents[1]

	// HTTPS is required.
	requestURL, err := url.Parse("https://" + fediverseServer)
	if err != nil {
		return nil, fmt.Errorf("unable to parse fediverse server host %s", fediverseServer)
	}

	requestURL.Path = "/.well-known/webfinger"
	query := requestURL.Query()
	query.Add("resource", fmt.Sprintf("acct:%s", account))
	requestURL.RawQuery = query.Encode()

	// Do not support redirects.
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	response, err := client.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var links webfingerResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&links); err != nil {
		return nil, err
	}

	return links.Links, nil
}
