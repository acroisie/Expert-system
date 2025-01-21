package helpers

import (
	"bufio"
	"expert-system/src/factManager"
	"expert-system/src/models"
	"expert-system/src/parser"
	"expert-system/src/v"
	"fmt"
	"os"
	"strings"
)

func ParseFile(inputFile string, problem *models.Problem) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "="):
			parseInitialFacts(line[1:], problem)
		case strings.HasPrefix(line, "?"):
			parseQueries(line[1:], problem)
		default:
			parseRule(line, problem)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func parseInitialFacts(line string, problem *models.Problem) {
	buff := strings.Split(line, " ")
	initialFacts := buff[0]
	// fmt.Println("Initial facts: ", initialFacts)

	for _, letter := range initialFacts {
		if letter < 'A' || letter > 'Z' {
			fmt.Println("Error: Invalid initial fact")
			os.Exit(1)
		}

		problem.Facts = append(problem.Facts, factManager.Fact{
			Letter: letter,
			Value:  v.TRUE,
			Initial: true,
			Reason: factManager.Reason{Msg: "Initial fact"},
		})
	}
}

func parseQueries(line string, problem *models.Problem) {
	buff := strings.Split(line, " ")
	queries := buff[0]

	for _, letter := range queries {
		if letter < 'A' || letter > 'Z' {
			fmt.Println("Error: Invalid query")
			os.Exit(1)
		}

		query := models.Query{
			Letter: letter,
		}
		problem.Queries = append(problem.Queries, query)
	}
}

func parseRule(line string, problem *models.Problem) {
	line = strings.ReplaceAll(line, " ", "")
	buff := strings.Split(line, "#")
	rule := buff[0]

	p := parser.NewParser(rule)

	r, err := p.ParseRule()

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	problem.Rules = append(problem.Rules, *r)
}
