package helpers

import (
	"fmt"
	"bufio"
	"os"
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
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error: ", err)
    }
}