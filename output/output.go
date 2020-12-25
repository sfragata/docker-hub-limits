package output

import (
	"fmt"
	"strings"

	"github.com/sfragata/docker-hub-limits/dockerhub"
)

//Type simulate an ENUM
type Type string

var outputMap = map[Type]outputable{
	jsonType: jsonOutput{},
	xmlType:  xmlOutput{},
	yamlType: yamlOutput{},
}

type outputable interface {
	toString(rateLimits dockerhub.RateLimitsInfo) (string, error)
}

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

	output := outputMap[outputType]

	if output == nil {
		return "", fmt.Errorf("Invalid output %s, valid values are: %s", outputType, printKeys())
	}

	response, err := output.toString(rateLimits)

	if err != nil {
		return "", err
	}

	return response, nil
}

func printKeys() string {
	var keys []string
	for k := range outputMap {
		keys = append(keys, string(k))
	}

	return strings.Join(keys, ", ")
}
