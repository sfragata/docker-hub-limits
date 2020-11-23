package dockerhub

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sfragata/docker-hub-limits/utils"
)

const validJSON = "{\"token\":\"1234\",\"access_token\":\"56768\",\"expires_in\":300,\"issued_at\":\"2020-11-23T03:11:56.701790516Z\"}"

func TestGivenValidJsonResponseWhenTokenThenReturnValidToken(test *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validJSON)
	}))
	defer server.Close()

	dockerhub := Info{
		Repository: "repo",
		authURL:    server.URL + "/%s",
	}

	err := Token(&dockerhub)

	if err != nil {
		test.Errorf("Error: %v", err)
	}

	if dockerhub.token != "1234" {
		test.Errorf("\nExpected token: %s \nactual : %s", "1234", dockerhub.token)
	}
}

func TestGivenServerErrorWhenTokenThenReturnError(test *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer server.Close()

	dockerhub := Info{
		Repository: "repo",
		authURL:    server.URL + "/%s",
	}

	err := Token(&dockerhub)

	if err == nil {
		test.Error("Should throw error")
	}

}

func TestGivenValidJsonResponseWhenTokenWithUsernamePasswordThenReturnValidToken(test *testing.T) {

	username := "foo"
	password := "bar"

	encodedAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validJSON)

		authHeader := r.Header.Get("Authorization")
		if utils.IsEmpty(authHeader) {
			test.Error("Should have 'Authorization' header")
		}
		if authHeader != "Basic "+encodedAuth {
			test.Errorf("\nExpected: %s\nActual: %s", encodedAuth, authHeader)
		}
	}))
	defer server.Close()

	dockerhub := Info{
		Repository: "repo",
		authURL:    server.URL + "/%s",
		Username:   username,
		Password:   password,
	}

	err := Token(&dockerhub)

	if err != nil {
		test.Errorf("Error: %v", err)
	}
}

func TestGivenInvalidJsonResponseWhenTokenThenReturnError(test *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "invalidJSON")
	}))
	defer server.Close()

	dockerhub := Info{
		Repository: "repo",
		authURL:    server.URL + "/%s",
	}

	err := Token(&dockerhub)

	if err == nil {
		test.Errorf("Error: %v", err)
	}
	errorToString := fmt.Sprintf("%v", err)
	if !strings.Contains(errorToString, "Invalid JSON") {
		test.Errorf("\nShould contain: %s\nActual: %s", "Invalid JSON", errorToString)
	}

}
