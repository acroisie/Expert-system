package helpers

import (
	"bufio"
	"expert-system/src/factManager"
	"expert-system/src/models"
	"expert-system/src/parser"
	"expert-system/src/rules"
	"expert-system/src/v"
	"fmt"
	"os"
	"strings"
)

func ParseFile(inputFile string, problem *models.Problem) {
	allLetters := make(map[rune]bool)
	initialFacts := make(map[rune]bool)

	problem.Rules = []rules.Rule{}
	problem.Queries = []models.Query{}
	problem.Facts = []factManager.Fact{}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", inputFile, err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if fileInfo.Size() == 0 {
		fmt.Println("Error: file is empty")
		os.Exit(1)
	}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "="):
			parseInitialFacts(line[1:], initialFacts)

		case strings.HasPrefix(line, "?"):
			parseQueries(line[1:], &problem.Queries, allLetters)

		default:
			r, lettersInRule, err := parseRule(line)
			if err != nil {
				fmt.Println("Error in parseRule:", err)
				os.Exit(1)
			}
			problem.Rules = append(problem.Rules, *r)
			for letter := range lettersInRule {
				allLetters[letter] = true
			}
		}
	}

	buildFacts(problem, allLetters, initialFacts)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file '%s': %v\n", inputFile, err)
		os.Exit(1)
	}
}

func parseInitialFacts(line string, initialFacts map[rune]bool) {
	trimmed := strings.Split(line, " ")[0]
	for _, letter := range trimmed {
		if letter < 'A' || letter > 'Z' {
			fmt.Printf("Error: invalid initial fact '%c'\n", letter)
			os.Exit(1)
		}
		initialFacts[letter] = true
	}
}

func parseQueries(line string, queries *[]models.Query, allLetters map[rune]bool) {
	trimmed := strings.Split(line, " ")[0]
	for _, letter := range trimmed {
		if letter < 'A' || letter > 'Z' {
			fmt.Printf("Error: invalid query '%c'\n", letter)
			os.Exit(1)
		}
		*queries = append(*queries, models.Query{Letter: letter})
		allLetters[letter] = true
	}
}

func parseRule(line string) (*rules.Rule, map[rune]bool, error) {
	line = strings.Split(line, "#")[0]
	line = strings.ReplaceAll(line, " ", "")

	p := parser.NewParser(line)
	r, err := p.ParseRule()
	if err != nil {
		return nil, nil, err
	}

	lettersInRule := collectLetters(r)
	return r, lettersInRule, nil
}

func collectLetters(rule *rules.Rule) map[rune]bool {
	letters := make(map[rune]bool)

	var traverseExpressionGroup func(*rules.ExpressionGroup)
	traverseExpressionGroup = func(eg *rules.ExpressionGroup) {
		if eg == nil {
			return
		}
		if eg.LeftVariable != nil {
			letters[eg.LeftVariable.Letter] = true
		}
		if eg.RightVariable != nil {
			letters[eg.RightVariable.Letter] = true
		}
		if eg.LeftExpressionGroup != nil {
			traverseExpressionGroup(eg.LeftExpressionGroup)
		}
		if eg.RightExpressionGroup != nil {
			traverseExpressionGroup(eg.RightExpressionGroup)
		}
	}

	if rule.LeftVariable != nil {
		letters[rule.LeftVariable.Letter] = true
	}
	if rule.RightVariable != nil {
		letters[rule.RightVariable.Letter] = true
	}

	traverseExpressionGroup(rule.LeftExpressionGroup)
	traverseExpressionGroup(rule.RightExpressionGroup)

	return letters
}

func buildFacts(problem *models.Problem, allLetters, initialFacts map[rune]bool) {
	for letter := range allLetters {
		val := v.UNKNOWN
		init := false

		if initialFacts[letter] {
			val = v.TRUE
			init = true
		}

		problem.Facts = append(problem.Facts, factManager.Fact{
			Letter:  letter,
			Value:   val,
			Initial: init,
			Reason:  factManager.Reason{Msg: ""},
		})
	}
}
