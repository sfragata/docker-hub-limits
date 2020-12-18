# docker-hub-limits

Utility to check [docker download rate limits](https://docs.docker.com/docker-hub/download-rate-limit/)

![Golang CI](https://github.com/sfragata/docker-hub-limits/workflows/Golang%20CI/badge.svg)

## Usage

```
Usage of docker-hub-limits:
  -output string
    	output format (json, yaml or xml)
  -password string
    	password registered in hub.docker.com
  -repository string
    	Docker repository hosted in hub.docker.com
  -username string
    	username registered in hub.docker.com
  -verbose
    	verbose mode
```