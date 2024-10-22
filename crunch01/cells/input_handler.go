package cells

// handling for height and width
func HandleHW(h, w int) bool {
	ok := true

	if h < 0 || w < 0 {
		ok = false
	}

	return ok
}

// handling for height and width
func HandleField(w int, str string) bool {
	ok := true

	if w != len(str) {
		ok = false
	}
	for i := 0; i < w; i++ {
		if rune(str[i]) < '0' || rune(str[i]) > '3' {
			ok = false
		}
	}
	return ok
}

func HandleString(w int, str string) bool {
	ok := true

	if w != len(str) {
		ok = false
	}

	return ok
}

// handle array
