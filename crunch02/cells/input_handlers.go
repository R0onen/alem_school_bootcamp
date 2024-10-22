package cells

func HWHandler(h int, w int) bool {
	if IsNumeric(Num(h)) && IsNumeric(Num(w)) {
		return h >= 3 && w >= 3 && h <= 35 && w <= 35
	} else {
		return false
	}
}

func FieldHandlerBombcount(count int) bool {
	return count >= 2
}

func FieldInputHandler(str string) bool {
	ok := false
	for i := 0; i < len(str); i++ {
		if str[i] == '.' || str[i] == '*' {
			ok = true
		} else {
			ok = false
		}
	}
	return ok
}
