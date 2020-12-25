package dockerhub

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/sfragata/docker-hub-limits/utils"
)

//RateLimits func token
func RateLimits(dockerHubInfo Info) (RateLimitsInfo, error) {

	if utils.IsEmpty(dockerHubInfo.token) {
		return RateLimitsInfo{}, fmt.Errorf("Null token")
	}

	// setting URL
	if utils.IsEmpty(dockerHubInfo.rateLimitURL) {
		dockerHubInfo.rateLimitURL = rateLimitsTemplateURL
	}

	urlRateLimits := fmt.Sprintf(dockerHubInfo.rateLimitURL, dockerHubInfo.Repository)
	validURL, err := url.Parse(urlRateLimits)

	if err != nil {
		return RateLimitsInfo{}, err
	}

	request, err := http.NewRequest("HEAD", validURL.String(), nil)

	if err != nil {
		return RateLimitsInfo{}, err
	}

	request.Header.Add("Authorization", fmt.Sprintf(bearerToken, dockerHubInfo.token))

	if dockerHubInfo.Verbose {
		log.Printf("Getting rate limits (%s)\n", validURL.String())
	}

	client := &http.Client{
		Timeout: httpTimeout * time.Second,
	}

	response, err := client.Do(request)

	if err != nil {
		return RateLimitsInfo{}, err
	}

	if response.StatusCode != 200 {
		return RateLimitsInfo{}, fmt.Errorf("Error: status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	if dockerHubInfo.Verbose {
		log.Println("response:")
		for name, value := range response.Header {
			fmt.Printf("\t%v: %v\n", name, value)
		}
	}

	rateLimit, err := utils.ExtractLimits(response.Header.Get("Ratelimit-Limit"))
	if err != nil {
		rateLimit = 0
		fmt.Printf("Error, couldn't extract rate limit [%s]\n", response.Header.Get("Ratelimit-Limit"))
	}

	remainingLimit, err := utils.ExtractLimits(response.Header.Get("Ratelimit-Remaining"))

	if err != nil {
		remainingLimit = 0
		fmt.Printf("Error, couldn't extract remaining limit [%s]\n", response.Header.Get("Ratelimit-Remaining"))
	}

	return RateLimitsInfo{Limit: rateLimit, Remaining: remainingLimit, ImageName: dockerHubInfo.Repository}, nil
}
