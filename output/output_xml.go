package output

import (
	"encoding/xml"

	"github.com/sfragata/docker-hub-limits/dockerhub"
)

type xmlRateLimitsResponse struct {
	Imagename string `xml:"image_name"`
	Limit     int    `xml:"limit"`
	Remaining int    `xml:"remaining"`
}

func toXML(rateLimits dockerhub.RateLimitsInfo) (string, error) {

	xmlResponse := xmlRateLimitsResponse{
		Imagename: rateLimits.ImageName,
		Limit:     rateLimits.Limit,
		Remaining: rateLimits.Remaining,
	}

	resp, err := xml.Marshal(xmlResponse)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
