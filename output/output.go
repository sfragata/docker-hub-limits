package output

import (
	"fmt"

	"github.com/sfragata/docker-hub-limits/dockerhub"
)

//Type simulate an ENUM
type Type string

const (
	//jsonType json type
	jsonType Type = "json"
	//xmlType xml type
	xmlType = "xml"
	//yamlType yaml type
	yamlType = "yaml"
)

//Marshal get struct and convert to given output type
func Marshal(rateLimits dockerhub.RateLimitsInfo, outputType Type) (string, error) {

	validationError := outputType.validate()
	if validationError != nil {
		return "", validationError
	}

	var response string
	var err error

	if outputType == jsonType {
		response, err = toJSON(rateLimits)
	} else if outputType == xmlType {
		response, err = toXML(rateLimits)
	} else if outputType == yamlType {
		response, err = toYAML(rateLimits)
	}

	if err != nil {
		return "", err
	}

	return response, nil
}

// validate if output is valid (extends OutputType type)
func (outputType Type) validate() error {
	switch outputType {
	case jsonType, xmlType, yamlType:
		return nil
	}
	return fmt.Errorf("Output must be `json`, `yaml` or `xml`")
}
