package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Checks if only 1 argument is provided
// Returns the first command line argument
func parseInput(args []string) (string, error) {
	if len(args) == 1 {
		return "", fmt.Errorf("No argument provided, exiting")
	} else if len(args) > 2 {
		return "", fmt.Errorf("Too many arguments provided only 1 argument is supported, exiting")
	}
	return args[1], nil
}

// Creates an AWS session
// Retrieves and decrypts a given parameter
func retriveParam(name string) (*ssm.GetParameterOutput, error) {
	sess := session.Must(session.NewSession())
	svc := ssm.New(sess)

	decrypt := true

	out, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: &decrypt})
	return out, err
}

// return value from
func getParamValue(paramOutput *ssm.GetParameterOutput) string {
	return *paramOutput.Parameter.Value
}

func main() {
	paramName, err := parseInput(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	out, err := retriveParam(paramName)
	if err != nil {
		log.Println("There was an error fetching/decrypting the parameter:", paramName)
		log.Fatal(err.Error())
	} else {
		fmt.Println(getParamValue(out))
	}
}
