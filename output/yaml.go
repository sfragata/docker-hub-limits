package output

import (
	"gopkg.in/yaml.v2"

	"github.com/sfragata/docker-hub-limits/dockerhub"
)

type yamlOutput struct{}
type yamlRateLimitsResponse struct {
	Imagename string `yaml:"image_name"`
	Limit     int    `yaml:"limit"`
	Remaining int    `yaml:"remaining"`
}

func (out yamlOutput) toString(rateLimits dockerhub.RateLimitsInfo) (string, error) {

	yamlResponse := yamlRateLimitsResponse{
		Imagename: rateLimits.ImageName,
		Limit:     rateLimits.Limit,
		Remaining: rateLimits.Remaining,
	}

	resp, err := yaml.Marshal(yamlResponse)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
