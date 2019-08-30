package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// Checks if only 1 argument is provided
// Returns the first command line argument
func parseInput(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("No argument provided, exiting")
	} else if len(args) > 1 {
		return "", fmt.Errorf("Too many arguments provided only 1 argument is supported, exiting")
	}
	return args[0], nil
}

func retrieveEnv() (string, error) {
	filebyte, err := ioutil.ReadFile("/opt/environment")
	out := string(filebyte[:])
	out = strings.TrimSuffix(out, "\n")
	return out, err
}

// Creates an AWS session
// Retrieves and decrypts a given parameter
func retrieveParam(paramName string, getEnvOutput string) (*secretsmanager.GetSecretValueOutput, error) {
	secretsmanagerPath := getEnvOutput + paramName
	sess := session.Must(session.NewSession())
	svc := secretsmanager.New(sess)
	out, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId:           &secretsmanagerPath})
	return out, err
}

// return value from
func getParamValue(paramOutput *secretsmanager.GetSecretValueOutput) string {
	return *paramOutput.SecretString
}

// Return test param value
func getTestParamValue(param string) (string, error) {
	if param == "TEST_PARAM_VALUE" {
		return param, nil
	} else {
		return "", fmt.Errorf("Wrong value for test-mode, should be: TEST_PARAM_VALUE")
	}
}

func main() {
	// Bypass AWS calls in test mode
	testMode := flag.Bool("test-mode", false, "Enable test mode")
	flag.Parse()

	paramName, err := parseInput(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	// test mode
	if *testMode {
		out, err := getTestParamValue(paramName)
		if err != nil {
			log.Println("There was an error fetching/decrypting the parameter:", paramName)
			log.Fatal(err.Error())
		} else {
			fmt.Println(out)
		}
	} else {
		getEnvOutput, err := retrieveEnv()
		out, err := retrieveParam(paramName, getEnvOutput)
		if err != nil {
			log.Println("There was an error fetching/decrypting the parameter:", paramName)
			log.Fatal(err.Error())
		} else {
			fmt.Println(getParamValue(out))
		}
	}
}
