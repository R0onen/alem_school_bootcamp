package cells


func Num(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := num < 0 // True or False
	if isNegative {
		num = -num
	}

	var result []rune

	for num > 0 {
		remainder := num % 10
		result = append(result, rune(remainder + '0'))
		num /= 10
	}

	if isNegative {
		result = append(result, '-')
	}

	
	start := 0
	end := len(result) - 1
	for start < end {

		temp := result[start]
		result[start] = result[end]
		result[end] = temp

		start++
		end--
	}

	return string(result)
}
