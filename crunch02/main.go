package main

import (
	"crunch02/cells"
	"fmt"
)

func main() {
	cells.PrintWords("Choose a mode: \n")
	cells.PrintWords("1. Enter a custom map\n")
	cells.PrintWords("2. Generate a random map\n")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		cells.Custom()
		break
	case 2:
		cells.Random()
		break
	default:
		cells.PrintWords("Error. Incorrect Input")
		break
	}
}
