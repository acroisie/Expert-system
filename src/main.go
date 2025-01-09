package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"expert/models"
)

func main() {
	fmt.Println("--- Expert System ---\n")
	fmt.Println("--- RULES ---")
	rules := models.GetRulesMock()
    models.DisplayRules(rules)
	fmt.Println("\n--- FACTS ---")
    facts := models.GetFactsMock()
    models.DisplayFacts(facts)
	fmt.Println("\n--- RULES AFTER CONDITIONAL OPERATOR FORMATTER ---")
	rules = models.RulesConditionalOperatorFormatter(rules)
	models.DisplayRules(rules)
	
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
