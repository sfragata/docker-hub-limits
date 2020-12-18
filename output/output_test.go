package output

import (
	"testing"

	"github.com/sfragata/docker-hub-limits/dockerhub"
)

var dummyRateLimist = dockerhub.RateLimitsInfo{
	ImageName: "dummy_image",
	Limit:     10,
	Remaining: 5,
}

func TestGivenValidOutputTypeWhenValidateThenReturnNoError(test *testing.T) {
	validType := Type("json")
	err := validType.validate()
	if err != nil {
		test.Errorf("Error: %v", err)
	}

}

func TestGivenInvalidOutputTypeWhenValidateThenReturnError(test *testing.T) {
	validType := Type("fake")
	err := validType.validate()
	if err == nil {
		test.Error("Should throw error")
	}

}

func TestGivenEmptyOutputTypeWhenValidateThenReturnError(test *testing.T) {
	var validType Type
	err := validType.validate()
	if err == nil {
		test.Error("Should throw error")
	}

}
