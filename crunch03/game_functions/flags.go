package game_functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func PrintHelp() {
	fmt.Println("Options:")
	fmt.Println("  -help         Show the help message and exit")
	fmt.Println("  -verbose      Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	fmt.Println("  -delay-ms=    Set the animation speed in milliseconds. Default is 2500 milliseconds")
	fmt.Println("  -file=X       Load the initial grid from a specified file")
	fmt.Println("  -edges-portal Enable portal edges where cells that exit the grid appear on the opposite side")
	fmt.Println("  -random=WxH   Generate a random grid of the specified width (W) and height (H)")
	fmt.Println("  -fullscreen   Adjust the grid to fit the terminal size with empty cells")
	fmt.Println("  -footprints   Add traces of visited cells, displayed as '∘'")
	fmt.Println("  -colored      Add color to live cells and traces if footprints are enabled")
}

func ReadGridFromFile(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	var rowLength int
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		if line == "" { // If line is empty it continues ignoring the line
			continue
		}
		if rowLength == 0 {
			rowLength = len(line)
		} else if len(line) != rowLength {
			return nil, fmt.Errorf("Invalid row length.")
		}

		if !IsValidRow(line) {
			return nil, fmt.Errorf("Invalid characters. Only '.' or '#' are expected.")
		}
		row := []rune(line)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if lineCount == 0 {
		return nil, fmt.Errorf("The file is empty.")
	}

	return grid, nil
}

func PrintColoredGrid(grid [][]rune, visited [][]bool, age [][]int) {
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '#' { // Assuming '█' represents live cells
				if age[i][j] < 3 {
					fmt.Print("\033[32m", "× ", "\033[0m") // Green for young cells
				} else if age[i][j] < 6 {
					fmt.Print("\033[33m", "× ", "\033[0m") // Yellow for middle-aged cells
				} else {
					fmt.Print("\033[31m", "× ", "\033[0m") // Red for old cells
				}
			} else if visited[i][j] {
				fmt.Print("\033[34m", "∘ ", "\033[0m") // Blue for footprints
			} else {
				fmt.Print("· ")
			}
		}
		fmt.Println()
	}
}

func GenerateRandomGrid(size string) ([][]rune, error) {
	parts := strings.Split(size, "x")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid format for random size")
	}

	height := ToInt(parts[0])
	width := ToInt(parts[1])

	if width < 3 || width > 21 || height < 3 || height > 21 {
		return nil, fmt.Errorf("Size must be between 3 and 20")
	}

	rand.Seed(time.Now().UnixNano()) // Random number generator

	grid := make([][]rune, height) // Grid with random cells
	for i := 0; i < height; i++ {
		grid[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			if rand.Intn(2) == 0 {
				grid[i][j] = '#'
			} else {
				grid[i][j] = '.'
			}
		}
	}
	return grid, nil
}
