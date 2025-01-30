package main

import (
    "expert-system/src/helpers"
    "expert-system/src/models"
    "expert-system/src/factManager"
    "expert-system/src/algo"
    "expert-system/src/rules"
    "fmt"
    "os"
)

func main() {
    problem := models.Problem{}

    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <input file>")
        return
    }
    helpers.ParseFile(os.Args[1], &problem)

    displayLogs := true
    if len(os.Args) >= 4 {
        displayLogs = false
    }

    factManager.FactList = problem.Facts

    if displayLogs {
        fmt.Println("---------- AST FOR RULES ----------")
        for i, rule := range problem.Rules {
            fmt.Printf("Rule %d:\n", i+1)
            rule.PrintAST()
            fmt.Println()
        }
        // factManager.Display()
        // rules.DisplayRules(problem.Rules)
    }

    formattedRules := rules.RulesConditionalOperatorFormatter(problem.Rules)
    result, reasoningLogs := algo.Algo(formattedRules)
    if result {
        if false {
            fmt.Printf("\n\nREASONING LOGS:\n")
            for _, log := range reasoningLogs {
                fmt.Printf(log)
            }
        }
        
        factManager.SortFactListByAlphabet(factManager.FactList)
        if displayLogs {
            fmt.Printf("\n\nFACT LIST RESULT:\n")
            factManager.Display()
        }

        if len(os.Args) >= 3 {
            res := helpers.TestFactList(os.Args[2], problem.Queries, displayLogs)
            if !res {
                os.Exit(1)
            }
        }
        os.Exit(0)
    }
}
