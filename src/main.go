package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"expert/models"
)

func main() {
	rules := models.GetRulesMock()
    models.DisplayRules(rules)
    facts := models.GetFactsMock()
    models.DisplayFacts(facts)
	return 
	if (len(os.Args) != 2) {
		fmt.Println("Usage: go run src/main.go <filename>")
		return
	}

	inputFile := os.Args[1]
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Println("File content: ", string(content))
}