package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"

    "expert-system/src/helpers"
    "expert-system/src/models"
    "expert-system/src/menu"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <input file>")
        return
    }

    problem := models.Problem{}
    helpers.ParseFile(os.Args[1], &problem)

    initialModel := menu.InitMainModel(&problem)

    p := tea.NewProgram(initialModel)
    if err := p.Start(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}
