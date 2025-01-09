package helpers

import (
	"fmt"
	"bufio"
	"os"
	"strings"
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
	fmt.Println("Initial facts: ", line)
}

func parseQueries(line string) {
	fmt.Println("Queries: ", line)
}

func parseRule(line string) {
	fmt.Println("Rule: ", line)
}