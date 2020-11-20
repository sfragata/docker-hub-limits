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
	dockerRepo := flag.String("docker-repo", "", "Docker repository hosted in dockerhub.io")
	username := flag.String("username", "", "username registered in dockerhub.io")
	password := flag.String("password", "", "password registered in dockerhub.io")

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
	}

	resp1, resp2, _ := dockerhub.RateLimits(dockerHubInfo)

	fmt.Println(resp1)
	fmt.Println(resp2)
}
