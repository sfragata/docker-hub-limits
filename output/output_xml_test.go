package output

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sfragata/docker-hub-limits/utils"
)

func TestGivenValidStructWhenToXMLThenReturnJsonString(test *testing.T) {

	response, err := toXML(dummyRateLimist)

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
