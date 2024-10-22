package cells

import "fmt"

func Custom() {
	PrintWords("Enter height and length:")
	var h, w int
	for {
		fmt.Scanf("%d %d", &h, &w)
		if !HWHandler(h, w) {
			PrintError()
			continue
		}
		break // Exit the loop if input was successful
	}
	field := make([][]rune, h)
	count := 0
	for {
		for i := 0; i < h; i++ {
			field[i] = make([]rune, w)
		}
		for row := 0; row < h; row++ {
			var str string
			fmt.Scanf("%s", &str)
			if !FieldInputHandler(str) {
				PrintError()
				continue
			}
			for j := 0; j < w; j++ {
				field[row][j] = rune(str[j])
			}
		}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if field[i][j] == '*' {
					count++
				}
			}
		}
		if !FieldHandlerBombcount(count) {
			PrintError()
			continue
		}
		break
	}
	FindBombs(h, w, field, count)
}
