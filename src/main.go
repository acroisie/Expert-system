package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"expert/factManager"
	"expert/rules"
)

func main() {
	fmt.Println("--- Expert System ---\n")
	ruleList := rules.GetRulesMock()
	ruleList = rules.RulesConditionalOperatorFormatter(ruleList)
	factManager.FactList = factManager.GetFactsMock()
	factManager.Display()
	subRuleList := ruleList
	algo(subRuleList)

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
