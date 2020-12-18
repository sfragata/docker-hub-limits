package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sfragata/docker-hub-limits/dockerhub"
	"github.com/sfragata/docker-hub-limits/output"
	"github.com/sfragata/docker-hub-limits/utils"
)

func main() {
	repository := flag.String("repository", "", "Docker repository hosted in hub.docker.com")
	username := flag.String("username", "", "username registered in hub.docker.com")
	password := flag.String("password", "", "password registered in hub.docker.com")
	verbose := flag.Bool("verbose", false, "verbose mode")
	outputFormatString := flag.String("output", "", "output format (json, yaml or xml)")

	flag.Parse()

	if utils.IsEmpty(*repository) {
		log.Println("docker-repo is mandatory")
		flag.PrintDefaults()
		os.Exit(1)
	}

	dockerHubInfo := dockerhub.Info{
		Repository: *repository,
		Username:   *username,
		Password:   *password,
		Verbose:    *verbose,
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

	if utils.IsNotEmpty(*outputFormatString) {
		response, err := output.Marshal(*rateLimits, output.Type(*outputFormatString))
		if err != nil {
			log.Fatalf("Error creating output %s : %v", *outputFormatString, err)
		}
		fmt.Println(response)
	} else {
		fmt.Printf("Image: %s | Limit: %d | Remaining: %d\n", dockerHubInfo.Repository, rateLimits.Limit, rateLimits.Remaining)
	}
}
