package main

import (
	"expert-system/src/helpers"
	"expert-system/src/menu"
	"expert-system/src/models"
	"os"
)

func main() {
	problem := models.Problem{}

	if len(os.Args) != 2 {
		menu.LaunchMenu(&problem)
		return
	}
	
	helpers.ParseFile(os.Args[1], &problem)
	menu.LaunchMenu(&problem)
}
