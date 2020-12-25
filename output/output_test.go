package output

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sfragata/docker-hub-limits/dockerhub"
	"github.com/sfragata/docker-hub-limits/utils"
)

var dummyRateLimist = dockerhub.RateLimitsInfo{
	ImageName: "dummy_image",
	Limit:     10,
	Remaining: 5,
}

func TestGivenValidStructWhenToStringThenReturnJsonString(test *testing.T) {

	jsonOutput := jsonOutput{}
	response, err := jsonOutput.toString(dummyRateLimist)

	if err != nil {
		test.Errorf("Error: %v", err)
	}

	if utils.IsEmpty(response) {
		test.Error("Error: response is empty")
	}

	if !strings.Contains(response, dummyRateLimist.ImageName) {
		test.Errorf("Error: response '%s' should contains '%s' as image name", response, dummyRateLimist.ImageName)
	}

	if !strings.Contains(response, strconv.Itoa(dummyRateLimist.Limit)) {
		test.Errorf("Error: response '%s' should contains '%d' as limit", response, dummyRateLimist.Limit)
	}

	if !strings.Contains(response, strconv.Itoa(dummyRateLimist.Remaining)) {
		test.Errorf("Error: response '%s' should contains '%d' as limit", response, dummyRateLimist.Remaining)
	}

}

func TestGivenValidStructWhenToStringThenReturnXmlString(test *testing.T) {

	xmlOutput := xmlOutput{}
	response, err := xmlOutput.toString(dummyRateLimist)

	if err != nil {
		test.Errorf("Error: %v", err)
	}

	if utils.IsEmpty(response) {
		test.Error("Error: response is empty")
	}

	if !strings.Contains(response, dummyRateLimist.ImageName) {
		test.Errorf("Error: response '%s' should contains '%s' as image name", response, dummyRateLimist.ImageName)
	}

	if !strings.Contains(response, strconv.Itoa(dummyRateLimist.Limit)) {
		test.Errorf("Error: response '%s' should contains '%d' as limit", response, dummyRateLimist.Limit)
	}

	if !strings.Contains(response, strconv.Itoa(dummyRateLimist.Remaining)) {
		test.Errorf("Error: response '%s' should contains '%d' as limit", response, dummyRateLimist.Remaining)
	}

}

func TestGivenValidStructWhenToStringThenReturnYamlString(test *testing.T) {

	yamlOutput := yamlOutput{}
	response, err := yamlOutput.toString(dummyRateLimist)

	if err != nil {
		test.Errorf("Error: %v", err)
	}

	if utils.IsEmpty(response) {
		test.Error("Error: response is empty")
	}

	if !strings.Contains(response, dummyRateLimist.ImageName) {
		test.Errorf("Error: response '%s' should contains '%s' as image name", response, dummyRateLimist.ImageName)
	}

	if !strings.Contains(response, strconv.Itoa(dummyRateLimist.Limit)) {
		test.Errorf("Error: response '%s' should contains '%d' as limit", response, dummyRateLimist.Limit)
	}

	if !strings.Contains(response, strconv.Itoa(dummyRateLimist.Remaining)) {
		test.Errorf("Error: response '%s' should contains '%d' as limit", response, dummyRateLimist.Remaining)
	}

}
