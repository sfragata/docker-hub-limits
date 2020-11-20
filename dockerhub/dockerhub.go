package dockerhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/sfragata/docker-hub-limits/utils"
)

//Info struct
type Info struct {
	Repository string
	Username   string
	Password   string
	token      string
	Verbose    bool
}

type tokenResponse struct {
	Token string `json:"token"`
	// AccessToken string `json:"access_token"`
	// ExpiresIn   int    `json:"expires_in"`
	// IssuedAt    string `json:"issued_at"`
}

//RateLimitsInfo struct
type RateLimitsInfo struct {
	Limit     int
	Remaining int
}

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
		Timeout: 5 * time.Second,
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

	jsonErr := json.Unmarshal(body, &tokenResponse)

	if jsonErr != nil {
		return jsonErr
	}
	dockerHubInfo.token = tokenResponse.Token

	return nil
}

//RateLimits func token
func RateLimits(dockerHubInfo Info) (*RateLimitsInfo, error) {

	token(&dockerHubInfo)

	if utils.IsEmpty(dockerHubInfo.token) {
		return nil, fmt.Errorf("Null token")
	}

	urlRateLimits := fmt.Sprintf(rateLimitsTemplateUrl, dockerHubInfo.Repository)
	validURL, err := url.Parse(urlRateLimits)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("HEAD", validURL.String(), nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf(bearerToken, dockerHubInfo.token))

	if dockerHubInfo.Verbose {
		log.Printf("Getting rate limits (%s)\n", validURL.String())
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error: status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	if dockerHubInfo.Verbose {
		log.Println("response:")
		for name, value := range response.Header {
			fmt.Printf("\t%v: %v\n", name, value)
		}
	}

	rateLimit, err := extractLimits(response.Header.Get("Ratelimit-Limit"))
	if err != nil {
		rateLimit = 0
		fmt.Printf("Error, couldn't extract rate limit [%s]\n", response.Header.Get("Ratelimit-Limit"))
	}

	remainingLimit, err := extractLimits(response.Header.Get("Ratelimit-Remaining"))

	if err != nil {
		remainingLimit = 0
		fmt.Printf("Error, couldn't extract remaining limit [%s]\n", response.Header.Get("Ratelimit-Remaining"))
	}

	return &RateLimitsInfo{Limit: rateLimit, Remaining: remainingLimit}, nil
}

func extractLimits(value string) (int, error) {
	if utils.IsEmpty(value) {
		return 0, nil
	}

	if strings.Contains(value, ";") {
		values := strings.Split(value, ";")
		if len(values) > 0 {
			return stringToInt(values[0])
		}
	}
	return stringToInt(value)

}

func stringToInt(value string) (int, error) {
	integer, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("Error converting %s to integer", value)
	}
	return integer, nil
}
