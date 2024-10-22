package main

import (
	"crunch01/cells"
	"fmt" // fmt.Scanf
)

func main() {
	var height, width int
	cells.PrintWords("Enter a values for height and width for your map:")
	fmt.Scanf("%d %d", &height, &width) // input height and width of field
	if !cells.HandleHW(height, width) {
		cells.PrintWords("Input Error: not values height and width")
		return
	}
	// initialization slice with heigth and width variables
	field := make([][]rune, height)

	for i := 0; i < height; i++ {
		field[i] = make([]rune, width)
	}

	cells.PrintWords("Draw your map using these values:")
	cells.PrintWords("0 - Wall\n1 - Free cell\n2 - Player\n3 - Award")
	// input types of fields using strings which are divided into runes
	for row := 0; row < height; row++ {
		var str string
		fmt.Scanf("%s ", &str)

		if !cells.HandleField(width, str) {
			cells.PrintWords("Input Error: not correct values for field's types")
			return
		}
		for j := 0; j < width; j++ {
			field[row][j] = rune(str[j])
		}
	}
	cells.PrintWords("Enter new characters to output on the cells. You should write it in that order:")
	cells.PrintWords("1. Wall\n2. Player\n3. Award")
	// input symbols which are going to change in cells
	var feature string
	fmt.Scanf("%s ", &feature)
	if !cells.HandleString(3, feature) {
		cells.PrintWords("Input Error: not correct value for changeable symbols")
		return
	}
	// calling function for creating a whole map
	cells.Render(height, width, field, feature)
}
