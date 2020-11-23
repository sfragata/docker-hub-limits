package utils

import (
	"fmt"
	"strconv"
	"strings"
)

//IsEmpty check is string is empty
func IsEmpty(parameter string) bool {
	return len(strings.TrimSpace(parameter)) == 0
}

//IsNotEmpty check is string is empty
func IsNotEmpty(parameter string) bool {
	return !IsEmpty(parameter)
}

//ExtractLimits split string by ;
func ExtractLimits(value string) (int, error) {
	if IsEmpty(value) {
		return 0, nil
	}

	if strings.Contains(value, ";") {
		values := strings.Split(value, ";")
		if len(values) > 0 {
			return StringToInt(values[0])
		}
	}
	return StringToInt(value)

}

//StringToInt convert string to int
func StringToInt(value string) (int, error) {
	integer, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("Error converting %s to integer", value)
	}
	return integer, nil
}
