package helpers

import (
    "fmt"
	"os"
	"bufio"
	"strings"
    "expert-system/src/factManager"
	"expert-system/src/v"
	"expert-system/src/models"
)

func ReadFactsFromFile(filename string) ([]factManager.Fact, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var facts []factManager.Fact
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		letter := rune(parts[0][0])
		valueStr := strings.TrimSpace(parts[1])

		var value v.Value
		if valueStr == "TRUE" {
			value = v.TRUE
		} else if valueStr == "FALSE" {
			value = v.FALSE
		} else {
			return nil, fmt.Errorf("invalid value in line: %s", line)
		}

		facts = append(facts, factManager.Fact{
			Letter:  letter,
			Value:   value,
			Initial: false,
			Reason:  factManager.Reason{Msg: ""},
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return facts, nil
}

func TestFactList(filePath string, queries []models.Query, displayLogs bool) bool {
	testFacts, testErr := ReadFactsFromFile(filePath)
	if testErr != nil {
		fmt.Println(testErr)
		return false
	}

	if displayLogs {
		fmt.Println("\n\n---------- TEST ----------\n")
		fmt.Println("Letter = Expected - Result")
	}

	different := false
	for _, query := range queries {
		resultFact, err := factManager.GetFactReferenceByLetter(query.Letter)
		if err != nil {
			fmt.Printf("ResultFact error: %s\n", err)
			return false
		}
		testFact, testErr := factManager.GetFactReferenceByLetterExtern(query.Letter, testFacts)
		if testErr != nil {
			fmt.Printf("TestFact error: %s\n", testErr)
			return false
		}
		if displayLogs {
			fmt.Printf("%c = %s - %s\n", query.Letter, testFact.Value, resultFact.Value)
		}
		if testFact.Value != resultFact.Value {
			different = true
		}
	}
	if different {
		if displayLogs {
			fmt.Println("Test failed")
		}
		return false
	}
	if displayLogs {
		fmt.Println("Test passed!")
	}
	return true
}