package game_functions

import "fmt"

func ClearTerminal() {
	// ANSI escape code to clear the screen and move the cursor to the top left
	fmt.Print("\033[H\033[2J")
}
