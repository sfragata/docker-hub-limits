package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/integrii/flaggy"
	"github.com/sfragata/docker-hub-limits/dockerhub"
	"github.com/sfragata/docker-hub-limits/output"
	"github.com/sfragata/docker-hub-limits/utils"
)

// These variables will be replaced by real values when do gorelease
var (
	version = "none"
	date    string
	commit  string
)

func main() {

	info := fmt.Sprintf(
		"%s\nDate: %s\nCommit: %s\nOS: %s\nArch: %s",
		version,
		date,
		commit,
		runtime.GOOS,
		runtime.GOARCH,
	)

	flaggy.SetName("docker-hub-limits")
	flaggy.SetDescription("Utility to check docker download rate limits")

	var repository string
	flaggy.String(&repository, "r", "repository", "Docker repository hosted in hub.docker.com")

	var username string
	flaggy.String(&username, "u", "username", "username registered in hub.docker.com")

	var password string
	flaggy.String(&password, "p", "password", "password registered in hub.docker.com")

	var verbose = false
	flaggy.Bool(&verbose, "v", "verbose", "verbose mode")

	var outputFormatString string
	flaggy.String(&outputFormatString, "o", "output", "output format (json, yaml or xml)")

	flaggy.SetVersion(info)

	flaggy.Parse()

	if utils.IsEmpty(repository) {
		flaggy.ShowHelpAndExit("")
	}

	dockerHubInfo := dockerhub.Info{
		Repository: repository,
		Username:   username,
		Password:   password,
		Verbose:    verbose,
	}

	// Getting token
	err := dockerhub.Token(&dockerHubInfo)

	if err != nil {
		log.Fatalf("Error getting token: %v", err)
	}

	// Getting rate limits
	rateLimits, err := dockerhub.RateLimits(dockerHubInfo)

	if err != nil {
		log.Fatalf("Error getting rate limits: %v", err)
	}

	if utils.IsNotEmpty(outputFormatString) {
		response, err := output.Marshal(rateLimits, output.Type(outputFormatString))
		if err != nil {
			log.Fatalf("Error creating output %s : %v", outputFormatString, err)
		}
		fmt.Println(response)
	} else {
		fmt.Printf("Image: %s | Limit: %d | Remaining: %d\n", dockerHubInfo.Repository, rateLimits.Limit, rateLimits.Remaining)
	}
}
