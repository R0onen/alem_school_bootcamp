package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args[1:]
	argCount := len(args)
	var message, breed string

	// Parse the command-line arguments
	if argCount == 1 && args[0] != "-b" {
		message = args[0]
		breed = "simple"
	} else if argCount == 3 && args[0] == "-b" {
		message = args[2]
		breed = args[1]
	} else if argCount == 2 {
		breed = args[1]
	} else {
		printString("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n") // Added newline here
		return
	}

	// Display the message and dog based on breed
	printDogSay(message, breed)
	ap.PutRune('\n')
}

// Prints a text bubble around the message
func printTextBubble(message string) {
	maxLineLength, lineCount := calculateMaxLineWidth(message)

	// Adjust the bubble width
	if maxLineLength+2 < 6 {
		maxLineLength = 4
	}

	// Print the top border of the bubble
	for i := 0; i < maxLineLength+2; i++ {
		if i == 0 {
			ap.PutRune(' ')
		}
		ap.PutRune('_')
	}
	ap.PutRune('\n')

	// Split message into lines
	lines := splitByNewline(message)

	// Print the message inside the bubble
	if lineCount+1 == 1 {
		printString("< ")
		printString(message)
		printString(" >")
	} else {
		printString("/ ")
		printString(lines[0])
		for i := 0; i < maxLineLength-len(lines[0]); i++ {
			ap.PutRune(' ')
		}
		printString(" \\")
		ap.PutRune('\n')

		// Print each line of the message
		for _, line := range lines[1 : len(lines)-1] {
			printString("| ")
			printString(line)
			for i := 0; i < maxLineLength-len(line); i++ {
				ap.PutRune(' ')
			}
			printString(" |")
			ap.PutRune('\n')
		}

		// Print the last line
		printString("\\ ")
		printString(lines[len(lines)-1])
		for i := 0; i < maxLineLength-len(lines[len(lines)-1]); i++ {
			ap.PutRune(' ')
		}
		printString(" /")
	}

	// Print the bottom border of the bubble
	ap.PutRune('\n')
	for i := 0; i < maxLineLength+2; i++ {
		if i == 0 {
			ap.PutRune(' ')
		}
		ap.PutRune('-')
	}
	ap.PutRune('\n')
}

// Displays the dog message with the chosen breed
func printDogSay(message, breed string) {
	switch breed {
	case "simple":
		printTextBubble(message)
		printString("  \\\n" +
			"^..^      /\n" +
			"/_/\\_____/\n" +
			"   /\\   /\\\n" +
			"  /  \\ /  \\")
	case "maltipoo":
		printTextBubble(message)
		printString(
			"  \\\n" +
				"   \\ __    __\n" +
				"   o-''))_____\\\\\n" +
				"   \"--__/ * * * )\n" +
				"   c_c__/-c____/")
	case "tazy":
		printTextBubble(message)
		printString(
			"  \\                __\n" +
				"   \\___________   /  \\\n" +
				"               \\ / ..|\\\n" +
				"                (_\\  |_)\n" +
				"                /  \\@'\n" +
				"               /     \\\n" +
				"           _  /  `   |\n" +
				"          \\\\ /  \\  | _\\\n" +
				"           \\   /_ || \\\\_\n" +
				"            \\____)|_) \\_)")
	case "cur":
		printTextBubble(message)
		printString(
			"   \\\n" +
				"    \\ D\\___/\\\n" +
				"     \\ (0_o)\n" +
				"        (V)\n" +
				"----oOo--U--oOo------\n" +
				"_______|_______|_____")
	default:
		printString("usage: dogsay [-b cur|maltipoo|simple|tazy] text") 
		return
	}
}

// Calculates the max width of a line in the message and the number of lines
func calculateMaxLineWidth(message string) (int, int) {
	maxWidth := 0
	lineWidth := 0
	lineCount := 0

	for i, char := range message {
		if char == '\n' {
			lineCount++
			if lineWidth > maxWidth {
				maxWidth = lineWidth
			}
			lineWidth = 0
		} else {
			lineWidth++
		}

		// Update max width for the last line
		if i == len(message)-1 && char != '\n' && lineWidth > maxWidth {
			maxWidth = lineWidth
		}
	}

	return maxWidth, lineCount
}

// Helper function to print a string
func printString(message string) {
	runes := []rune(message)
	for i := 0; i < len(runes); i++ {
		ap.PutRune(runes[i])
	}
}

// Splits the message by newline characters
func splitByNewline(message string) []string {
	if message == "" {
		return []string{}
	}
	var lines []string
	temp := ""
	for _, char := range message {
		if char == '\n' {
			lines = append(lines, temp)
			temp = ""
		} else {
			temp += string(char)
		}
	}
	if temp != "" || message[len(message)-1] == '\n' {
		lines = append(lines, temp)
	}
	return lines
}
