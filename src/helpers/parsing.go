package helpers

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"expert-system/src/models"
)

func ParseFile(inputFile string) {
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
			parseInitialFacts(line[1:])
		case strings.HasPrefix(line, "?"):
			parseQueries(line[1:])
		default:
			parseRule(line)
		}
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error: ", err)
    }
}

func parseInitialFacts(line string) {
	buff := strings.Split(line, " ")
	initialFacts := buff[0]
	fmt.Println("Initial facts: ", initialFacts)

	for _, letter := range initialFacts {
		if letter < 'A' || letter > 'Z' {
			fmt.Println("Error: Invalid initial fact")
			os.Exit(1)
		}

		fact := models.Fact{
			Letter: letter,
			Value: models.TRUE,
			Initial: true,
			Reason: models.Reason{Msg: "Initial fact"},
		}
		fmt.Println(fact) // Append to facts and store it in a global variable
	}
}

func parseQueries(line string) {
	// fmt.Println("Queries: ", line)
}

func parseRule(line string) {
	// fmt.Println("Rule: ", line)
}