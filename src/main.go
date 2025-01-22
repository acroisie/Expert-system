package main

import (
	"expert-system/src/helpers"
	"expert-system/src/models"
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

	fmt.Println(problem.Facts)
	fmt.Println(problem.Queries)
	fmt.Println(problem.Rules)
}
