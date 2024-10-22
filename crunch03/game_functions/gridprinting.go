package game_functions

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func EnterValue(s string, pointer *int) {
	for {
		fmt.Printf("Enter " + s + " (must be between 3 and 20): ")
		var userInput string
		fmt.Scanf("%s", &userInput)
		if IsNumber(userInput) {
			value := ToInt(userInput)
			if value >= 3 && value <= 20 {
				*pointer = value
				break
			} else {
				fmt.Println("Error: Invalid input. The value must be between 3 and 20.")
			}
		} else {
			fmt.Println("Error: Please enter a valid number.")
		}
	}
}

func EnterGrid(height, width int) [][]rune {
	fmt.Println("\nEnter the game field (use '#' for live cells and '.' for dead cells):")
	grid := make([][]rune, height)
	for i := 0; i < height; i++ {
		var line string
		fmt.Printf("Row %d: ", i+1)
		fmt.Scanf("%s", &line)
		if len(line) == width && IsValidRow(line) {
			grid[i] = []rune(line)
		} else {
			fmt.Println("Error: Invalid input. The row length must match the width, and only '.' or '#' are allowed.")
			i--
		}
	}
	return grid
}

func PrintGrid(g [][]rune, visited [][]bool) {
	for i, row := range g {
		for j, cell := range row {
			if cell == '#' {
				fmt.Print("× ") // Live cell
			} else if visited[i][j] {
				fmt.Print("∘ ") // Footprint for visited cells
			} else {
				fmt.Print("· ") // Dead cell
			}
		}
		fmt.Println()
	}
}

func GetTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0, err
	}
	width = width / 2
	height--
	return height, width, nil
}
func AdjustGridToFullscreen(grid [][]rune, currentHeight, currentWidth, termHeight, termWidth int) [][]rune {
	// Create a new grid with the terminal dimensions
	newGrid := make([][]rune, termHeight)

	for i := 0; i < termHeight; i++ {
		// Initialize a new row
		newRow := make([]rune, termWidth)

		for j := 0; j < termWidth; j++ {
			if i < currentHeight && j < currentWidth {
				// If within the bounds of the existing grid, copy the value
				newRow[j] = grid[i][j]
			} else {
				// Fill with dead cells ('.') if beyond current grid dimensions
				newRow[j] = '.'
			}
		}
		newGrid[i] = newRow
	}

	return newGrid
}
