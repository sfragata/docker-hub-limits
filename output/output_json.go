package output

import (
	"encoding/json"

	"github.com/sfragata/docker-hub-limits/dockerhub"
)

type jsonRateLimitsResponse struct {
	Imagename string `json:"image_name"`
	Limit     int    `json:"limit"`
	Remaining int    `json:"remaining"`
}

func toJSON(rateLimits dockerhub.RateLimitsInfo) (string, error) {

	jsonResponse := jsonRateLimitsResponse{
		Imagename: rateLimits.ImageName,
		Limit:     rateLimits.Limit,
		Remaining: rateLimits.Remaining,
	}

	resp, err := json.Marshal(jsonResponse)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
