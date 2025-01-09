package main

import (
	"fmt"
    "os"
    "expert-system/src/helpers"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <input file>")
        return
    }
    helpers.ParseFile(os.Args[1])
}
