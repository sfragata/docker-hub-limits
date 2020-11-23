package dockerhub

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGivenValidInfosWhenRateLimitsThenReturnValidValues(test *testing.T) {

	expectedToken := "1234"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		expectedBearerToken := fmt.Sprintf(bearerToken, expectedToken)
		if authHeader != expectedBearerToken {
			test.Errorf("\nExpected: %s\nActual: %s", expectedBearerToken, authHeader)
		}
		w.Header().Set("Ratelimit-Limit", "50;w=111")
		w.Header().Set("Ratelimit-Remaining", "38;w=111")
	}))
	defer server.Close()

	dockerhub := Info{
		Repository:   "repo",
		rateLimitURL: server.URL + "/%s",
		token:        expectedToken,
	}

	rateLimits, err := RateLimits(dockerhub)

	if err != nil {
		test.Errorf("Error: %v", err)
	}
	if rateLimits.Limit != 50 {
		test.Errorf("Limits: \nExpected: %d\nActual: %d", 50, rateLimits.Limit)
	}

	if rateLimits.Remaining != 38 {
		test.Errorf("Remaining: \nExpected: %d\nActual: %d", 38, rateLimits.Remaining)
	}

}

func TestGivenInvalidLimitsWhenRateLimitsThenReturnZero(test *testing.T) {

	expectedToken := "1234"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Ratelimit-Limit", "aaa")
		w.Header().Set("Ratelimit-Remaining", "bbb")
	}))
	defer server.Close()

	dockerhub := Info{
		Repository:   "repo",
		rateLimitURL: server.URL + "/%s",
		token:        expectedToken,
	}

	rateLimits, err := RateLimits(dockerhub)

	if err != nil {
		test.Errorf("Error: %v", err)
	}
	if rateLimits.Limit != 0 {
		test.Errorf("Limits: \nExpected: %d\nActual: %d", 50, rateLimits.Limit)
	}

	if rateLimits.Remaining != 0 {
		test.Errorf("Remaining: \nExpected: %d\nActual: %d", 38, rateLimits.Remaining)
	}

}

func TestGivenServerErrorWhenRateLimitsThenReturnError(test *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer server.Close()

	dockerhub := Info{
		Repository:   "repo",
		rateLimitURL: server.URL + "/%s",
		token:        "1234",
	}

	_, err := RateLimits(dockerhub)

	if err == nil {
		test.Error("Should throw error")
	}

}
