package main

import (
	"crunch03/game_functions"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	height, width int
	grid          [][]rune
	err           error
)

var (
	verbose     = flag.Bool("verbose", false, "Display detailed information about the map size, number of ticks, speed, and map name.")
	delayMs     = flag.Int("delay-ms", 2500, "Set the animation speed in milliseconds. Default is 2500 ms.")
	help        = flag.Bool("help", false, "Show help information.")
	file        = flag.String("file", "", "Read the initial grid from a specified file.")
	random      = flag.String("random", "", "Generates a random grid of the specified width (W) and height (H).")
	fullscreen  = flag.Bool("fullscreen", false, "Adjusts the grid to fit the terminal size with empty cells.")
	footprints  = flag.Bool("footprints", false, "Adds traces of visited cells using ∘.")
	colored     = flag.Bool("colored", false, "Adds color to live cells and traces if footprints are implemented.")
	edgesPortal = flag.Bool("edges-portal", false, "Enables portal edges where cells that exit the grid appear on the opposite side.")
)

func main() {
	flag.Parse()

	// Check for conflicts
	if *file != "" && *random != "" {
		fmt.Println("Conflicting flags: 'file' and 'random'. Using 'file'.")
		*random = ""
	}

	if *help {
		game_functions.PrintHelp()
		os.Exit(0)
	}

	if *file != "" {
		grid, err = game_functions.ReadGridFromFile(*file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
		height = len(grid)
		if height > 0 {
			width = len(grid[0])
		}
	} else if *random != "" {
		grid, err = game_functions.GenerateRandomGrid(*random)
		if err != nil {
			fmt.Printf("Error generating random grid: %v\n", err)
			os.Exit(1)
		}
		height = len(grid)
		if height > 0 {
			width = len(grid[0])
		}
	} else {
		// Get grid dimensions and values from user input
		game_functions.EnterValue("height", &height)
		game_functions.EnterValue("width", &width)
		grid = game_functions.EnterGrid(height, width)
	}
	if *fullscreen {
		termHeight, termWidth, err := game_functions.GetTerminalSize()
		if err != nil {
			fmt.Printf("Error getting terminal size: %v\n", err)
			os.Exit(1)
		}
		grid = game_functions.AdjustGridToFullscreen(grid, height, width, termHeight, termWidth)
		height, width = termHeight, termWidth
	}
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}
	numOfTicks := 0
	age := make([][]int, height)
	for i := range age {
		age[i] = make([]int, width)
	}

	for {
		if numOfTicks != 0 {
			game_functions.ClearTerminal()
			numOfLiveCells := game_functions.CountLiveCells(grid)

			// Mark cells as visited
			if *footprints {
				for i := range grid {
					for j := range grid[i] {
						if grid[i][j] == '#' {
							visited[i][j] = true
						}
					}
				}
			}

			// Calculate next grid based on edges portal if enabled
			nextGrid := game_functions.NextTick(grid, height, width, *edgesPortal, age)

			// Check for changes in the grid
			if game_functions.GridsAreEqual(grid, nextGrid, height, width) {
				game_functions.PrintGrid(grid, visited)
				fmt.Println("\nNo more changes, terminating program.")
				break
			}

			grid = nextGrid

			if *verbose {
				fmt.Println("\nTick:", numOfTicks)
				fmt.Println("Grid Size:", height, "x", width)
				fmt.Println("Live cells:", numOfLiveCells)
				fmt.Println("DelayMs:", *delayMs, "ms")
			}
			if *colored {
				game_functions.PrintColoredGrid(grid, visited, age)
			} else {
				game_functions.PrintGrid(grid, visited)
			}
			if game_functions.NoLiveCells(grid) {
				fmt.Println("\nAll cells are dead, terminating program.")
				break
			}
		} else {
			game_functions.ClearTerminal()
			numOfLiveCells := game_functions.CountLiveCells(grid)
			if *verbose {
				fmt.Println("\nTick:", numOfTicks)
				fmt.Println("Grid Size:", height, "x", width)
				fmt.Println("Live cells:", numOfLiveCells)
				fmt.Println("DelayMs:", *delayMs, "ms")
			}
			if *colored {
				game_functions.PrintColoredGrid(grid, visited, age)
			} else {
				game_functions.PrintGrid(grid, visited)
			}
		}
		numOfTicks++
		// Sleep for a minimum of 250 ms or the calculated delay
		if *delayMs < 250 {
			fmt.Print("Error too small value of delay")
		} else {
			time.Sleep(time.Duration(*delayMs) * time.Millisecond)
		}
	}
}
