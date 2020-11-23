package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sfragata/docker-hub-limits/dockerhub"
	"github.com/sfragata/docker-hub-limits/utils"
)

func main() {
	dockerRepo := flag.String("docker-repo", "", "Docker repository hosted in hub.docker.com")
	username := flag.String("username", "", "username registered in hub.docker.com")
	password := flag.String("password", "", "password registered in hub.docker.com")
	verbose := flag.Bool("verbose", false, "verbose mode")

	flag.Parse()

	if utils.IsEmpty(*dockerRepo) {
		log.Println("docker-repo is mandatory")
		flag.PrintDefaults()
		os.Exit(1)
	}

	dockerHubInfo := dockerhub.Info{
		Repository: *dockerRepo,
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
	fmt.Printf("Limit: %d \nRemaining: %d \n", rateLimits.Limit, rateLimits.Remaining)
}
