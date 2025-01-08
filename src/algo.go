package main

import (
    "fmt"
    "expert/models"
)

func main() {
    fmt.Println("Hello, World!")
    rules := models.GetRulesMock()
    facts := models.GetFactsMock()
    models.DisplayRules(rules)
    models.DisplayFacts(facts)
}