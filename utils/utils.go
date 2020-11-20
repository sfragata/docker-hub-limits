package utils

import "strings"

//IsEmpty check is string is empty
func IsEmpty(parameter string) bool {
	return len(strings.TrimSpace(parameter)) == 0
}

//IsNotEmpty check is string is empty
func IsNotEmpty(parameter string) bool {
	return !IsEmpty(parameter)
}
