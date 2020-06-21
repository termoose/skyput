package dappdappgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Post(skyHash string) error {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	form := url.Values{}
	form.Add("skylink", skyHash)

	req, _ := http.NewRequest("POST", portalUrl, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("post dappdappgo: %v", err)
	}

	var responseData Response
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&responseData)

	if responseData.Error != "" {
		return fmt.Errorf("🚫 DappDappGo: %s", responseData.Msg)
	}

	return nil
}
