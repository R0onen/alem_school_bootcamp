package cells

import "fmt"

func FindBombs(h int, w int, field [][]rune, count int) {
	if IsAllMines(field) {
		PrintError()
		return
	}
	unrevealed_cells := h*w - 4
	Renderfirsttime(h, w, field)
	countM := 0
	PrintWords("Total bombs in the field: ")
	PrintWords(Num(count))
	PrintWords("\n")

	var guessH, guessW int

	for {
		sum := CountAdjacentBombs(h, w, field, guessH, guessW)
		PrintWords("Enter coordinates to guess (or -1 -1 to quit): ")
		_, err := fmt.Scanf("%d %d", &guessH, &guessW)
		countM++
		if err != nil {
			PrintWords("Error reading input. Please try again.\n")
			continue
		}
		if guessH == -1 && guessW == -1 {
			break
		}

		guessH--
		guessW--
		if countM == 1 && field[guessH][guessW] == '*' {
			Replace(guessH, guessW, field)
		} else if field[guessH][guessW] == '*' {
			Allbombs(h, w, field, &sum)
			PrintWords("Game Over!\n")
			PrintStatistics(h, w, count, countM)
			break
		}

		// Проверка, открыта ли ячейка
		if field[guessH][guessW] != '.' {
			PrintWords("This cell is already opened. Try again.\n")
			continue
		}

		// Подсчет соседних бомб
		if sum > 0 {
			field[guessH][guessW] = rune('0' + sum) // Устанавливаем количество бомб
		} else {
			field[guessH][guessW] = ' ' // Устанавливаем пустую ячейку
			// Рекурсивное открытие соседних ячеек
			OpenAdjacentCells(h, w, field, guessH, guessW, &unrevealed_cells)
		}
		unrevealed_cells--
		if unrevealed_cells <= count {
			Allbombs(h, w, field, &sum)
			PrintWords("You Win!\n")
			PrintStatistics(h, w, count, countM)
			break
		}

		Render(h, w, field, &sum)
		PrintWords(Num(unrevealed_cells))
	}
}
