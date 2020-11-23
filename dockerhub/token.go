package dockerhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/sfragata/docker-hub-limits/utils"
)

//Token func token
func Token(dockerHubInfo *Info) error {

	// setting URL
	if utils.IsEmpty(dockerHubInfo.authURL) {
		dockerHubInfo.authURL = tokenTemplateURL
	}

	urltoken := fmt.Sprintf(dockerHubInfo.authURL, dockerHubInfo.Repository)
	validURL, err := url.Parse(urltoken)

	if err != nil {
		return err
	}

	request, err := http.NewRequest("GET", validURL.String(), nil)

	if err != nil {
		return err
	}

	if utils.IsNotEmpty(dockerHubInfo.Username) && utils.IsNotEmpty(dockerHubInfo.Password) {
		if dockerHubInfo.Verbose {
			log.Printf("Using docker hub credentials for user %s \n", dockerHubInfo.Username)
		}
		request.SetBasicAuth(dockerHubInfo.Username, dockerHubInfo.Password)
	} else {
		if dockerHubInfo.Verbose {
			log.Println("Using anonymous docker hub")
		}
	}

	client := &http.Client{
		Timeout: httpTimeout * time.Second,
	}

	if dockerHubInfo.Verbose {
		log.Printf("Getting token (%s)\n", validURL.String())
	}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("Error: status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if dockerHubInfo.Verbose {
		log.Printf("response: token (%s)\n", body)
	}

	tokenResponse := tokenResponse{}

	err = json.Unmarshal(body, &tokenResponse)

	if err != nil {
		return fmt.Errorf("Invalid JSON: %v", err)
	}
	dockerHubInfo.token = tokenResponse.Token

	return nil
}
