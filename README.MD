# docker-hub-limits

Utility to check [docker download rate limits](https://docs.docker.com/docker-hub/download-rate-limit/)

![Golang CI](https://github.com/sfragata/docker-hub-limits/workflows/Golang%20CI/badge.svg)

## Usage

```
docker-hub-limits - Utility to check docker download rate limits

  Flags:
       --version      Displays the program version string.
    -h --help         Displays help with available flag, subcommand, and positional value parameters.
    -r --repository   Docker repository hosted in hub.docker.com
    -u --username     username registered in hub.docker.com
    -p --password     password registered in hub.docker.com
    -v --verbose      verbose mode
    -o --output       output format (json, yaml or xml)
```