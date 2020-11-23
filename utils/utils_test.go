package utils

import (
	"fmt"
	"strings"
	"testing"
)

// IsEmpty

func TestGivenEmptyValueWhenIsEmptyThenReturnTrue(test *testing.T) {
	value := IsEmpty("")
	if !value {
		test.Errorf("expected: %v actual: %v", true, value)
	}
}

func TestGivenNonEmptyValueWhenIsEmptyThenReturnFalse(test *testing.T) {
	value := IsEmpty("values")
	if value {
		test.Errorf("expected: %v actual: %v", false, value)
	}
}

// IsNotEmpty

func TestGivenEmptyValueWhenIsNotEmptyThenReturnFalse(test *testing.T) {
	value := IsNotEmpty("")
	if value {
		test.Errorf("expected: %v actual: %v", false, value)
	}
}

func TestGivenNonEmptyValueWhenIsNotEmptyThenReturnTrue(test *testing.T) {
	value := IsNotEmpty("values")
	if !value {
		test.Errorf("expected: %v actual: %v", true, value)
	}
}

//StringToInt

func TestGivenValidStringNumberWhenStringToIntThenReturnInt(test *testing.T) {

	value, err := StringToInt("1")
	if err != nil {
		test.Errorf("Should not throw error : %v", err)
	}

	if value != 1 {
		test.Errorf("\nExpected: %v\nActual: %v", 1, value)
	}
}

func TestGivenInvalidStringNumberWhenStringToIntThenReturnError(test *testing.T) {

	expectedSubString := "Error converting"

	value, err := StringToInt("a")
	if err == nil {
		test.Error("Should throw error")
	}
	errorToString := fmt.Sprintf("%v", err)
	if !strings.Contains(errorToString, expectedSubString) {
		test.Errorf("\nExpected: %v\nActual: %v", expectedSubString, errorToString)
	}

	if value != 0 {
		test.Errorf("\nExpected: %v\nActual: %v", 0, value)
	}
}

// ExtractLimits

func TestGivenValidStringWithDelimiterWhenExtractLimitsThenReturnOnlyLimit(test *testing.T) {

	value, err := ExtractLimits("50;ffff")
	if err != nil {
		test.Errorf("Should not throw error : %v", err)
	}

	if value != 50 {
		test.Errorf("\nExpected: %v\nActual: %v", 50, value)
	}
}

func TestGivenValidStringWithoutDelimiterWhenExtractLimitsThenReturnValue(test *testing.T) {

	value, err := ExtractLimits("50")
	if err != nil {
		test.Errorf("Should not throw error : %v", err)
	}

	if value != 50 {
		test.Errorf("\nExpected: %v\nActual: %v", 50, value)
	}
}

func TestGivenEmptyStringWhenExtractLimitsThenReturnZero(test *testing.T) {

	value, err := ExtractLimits("")
	if err != nil {
		test.Errorf("Should not throw error : %v", err)
	}

	if value != 0 {
		test.Errorf("\nExpected: %v\nActual: %v", 0, value)
	}
}
