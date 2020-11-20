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

//Info struct
type Info struct {
	Repository string
	Username   string
	Password   string
	token      string
}

type tokenResponse struct {
	Token string `json:"token"`
	// AccessToken string `json:"access_token"`
	// ExpiresIn   int    `json:"expires_in"`
	// IssuedAt    string `json:"issued_at"`
}

//RateLimits struct
// type RateLimits struct {
// 	Limit     int
// 	Remaining int
// }

const tokenTemplateUrl = "https://auth.docker.io/token?service=registry.docker.io&scope=repository:%s:pull"
const rateLimitsTemplateUrl = "https://registry-1.docker.io/v2/%s/manifests/latest"
const bearerToken = "Bearer %s"

//token func token
func token(dockerHubInfo *Info) error {
	urltoken := fmt.Sprintf(tokenTemplateUrl, dockerHubInfo.Repository)
	validURL, err := url.Parse(urltoken)

	if err != nil {
		return err
	}

	request, err := http.NewRequest("GET", validURL.String(), nil)

	if err != nil {
		return err
	}

	if utils.IsNotEmpty(dockerHubInfo.Username) && utils.IsNotEmpty(dockerHubInfo.Password) {
		log.Printf("Using docker hub credentials for user %s \n", dockerHubInfo.Username)
		request.SetBasicAuth(dockerHubInfo.Username, dockerHubInfo.Password)
	} else {
		log.Println("Using anonymous docker hub")
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
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

	tokenResponse := tokenResponse{}

	jsonErr := json.Unmarshal(body, &tokenResponse)

	if jsonErr != nil {
		return jsonErr
	}
	dockerHubInfo.token = tokenResponse.Token

	return nil
}

//RateLimits func token
func RateLimits(dockerHubInfo Info) (int, int, error) {

	token(&dockerHubInfo)

	// fmt.Printf("token: %s \n", dockerHubInfo.token)

	if utils.IsEmpty(dockerHubInfo.token) {
		return 0, 0, fmt.Errorf("Null token")
	}

	urlRateLimits := fmt.Sprintf(rateLimitsTemplateUrl, dockerHubInfo.Repository)
	validURL, err := url.Parse(urlRateLimits)

	if err != nil {
		return 0, 0, err
	}

	request, err := http.NewRequest("HEAD", validURL.String(), nil)

	if err != nil {
		return 0, 0, err
	}

	request.Header.Add("Authorization", fmt.Sprintf(bearerToken, dockerHubInfo.token))

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Do(request)

	if err != nil {
		return 0, 0, err
	}

	if response.StatusCode != 200 {
		return 0, 0, fmt.Errorf("Error: status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	for name, value := range response.Header {
		fmt.Printf("%v: %v\n", name, value)
	}

	rateLimit := response.Header.Get("Ratelimit-Limit")
	remainingLimit := response.Header.Get("Ratelimit-Remaining")
	fmt.Printf("1 - %s | 2 - %s\n", rateLimit, remainingLimit)
	return 0, 0, nil
}
