package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ssm"
)

func TestParseInput(t *testing.T) {
	// Fail if more than one cmd line argument provided
	_, err := parseInput([]string{"1", "2", "3"})
	if err == nil {
		t.Errorf("More than 2 args should result in failure.")
	}
	// Fail if no argument provided
	_, err = parseInput([]string{"1"})
	if err == nil {
		t.Errorf("No args should result in failure.")
	}
	// should pass if 1 argument provided
	_, err = parseInput([]string{"please", "pass"})
	if err != nil {
		t.Errorf("One argument should not result in error")
	}
}

func TestGetParamValue(t *testing.T) {
	paramValue := "Yay"
	paramType := "String"
	paramName := "Test"
	input := ssm.GetParameterOutput{
		// Information about a parameter.
		Parameter: &ssm.Parameter{
			Name:  &paramName,
			Type:  &paramType,
			Value: &paramValue}}
	output string = getParamValue(&input)
	if output != paramValue {
		t.Errorf("Function should return param value")
	}
}
