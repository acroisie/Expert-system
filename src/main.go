package main

import (
	"fmt"
    "os"
    "expert-system/src/helpers"
	"expert-system/src/models"
    "expert-system/src/factManager"
    "expert-system/src/rules"
)

func main() {
    fmt.Println("--- Expert System ---\n")
	ruleList := rules.GetRulesMock()
	ruleList = rules.RulesConditionalOperatorFormatter(ruleList)
	factManager.FactList = factManager.GetFactsMock()
	subRuleList := ruleList
	Algo(subRuleList)
    return 
	problem := models.Problem{}
	
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <input file>")
        return
    }
    helpers.ParseFile(os.Args[1], &problem)

	fmt.Println(problem.Facts)
    fmt.Println(problem.Queries)
}