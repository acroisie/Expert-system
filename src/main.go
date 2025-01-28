package main

import (
    "expert-system/src/helpers"
    "expert-system/src/models"
    "expert-system/src/factManager"
    "expert-system/src/rules"
    "fmt"
    "os"
)

func main() {
    problem := models.Problem{}

    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <input file>")
        return
    }
    helpers.ParseFile(os.Args[1], &problem)

    fmt.Println("---------- AST FOR RULES ----------")
    for i, rule := range problem.Rules {
        fmt.Printf("Rule %d:\n", i+1)
        rule.PrintAST()
        fmt.Println()
    }

    factManager.FactList = problem.Facts
    factManager.Display()
    rules.DisplayRules(problem.Rules)
    Algo(problem.Rules)
}

