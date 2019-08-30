package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func TestParseInput(t *testing.T) {
	// Fail if more than one cmd line argument provided
	_, err := parseInput([]string{"1", "2", "3"})
	if err == nil {
		t.Errorf("More than 1 arg should result in failure.")
	}
	// Fail if no argument provided
	_, err = parseInput([]string{})
	if err == nil {
		t.Errorf("No args should result in failure.")
	}
	// should pass if 1 argument provided
	_, err = parseInput([]string{"pleasepass"})
	if err != nil {
		t.Errorf("One argument should not result in error")
	}
}

func TestGetParamValue(t *testing.T) {
	paramValue := "Yay"
	paramType := "String"
	paramName := "Test"
	input := secretsmanager.GetSecretValueOutput{
		// Information about a parameter.
		Parameter: &secretsmanager.Parameter{
			Name:  &paramName,
			Type:  &paramType,
			Value: &paramValue}}
	output := getParamValue(&input)
	if output != paramValue {
		t.Errorf("Function should return param value")
	}
}

func TestGetTestParamValue(t *testing.T) {
	paramValue := "TEST_PARAM_VALUE"
	value, err := getTestParamValue(paramValue)
	if err != nil || value != paramValue {
		t.Errorf("Function should return the same param value")
	}

	paramValue = "Something"
	_, err = getTestParamValue(paramValue)
	if err == nil {
		t.Errorf("Function should return error if test-mode value is not passed")
	}
}
