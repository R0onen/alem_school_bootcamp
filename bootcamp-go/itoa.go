package bootcamp

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := n < 0 // True or False
	if isNegative {
		n = -n
	}

	var result []rune

	for n > 0 {
		remainder := n % 10
		result = append(result, rune(remainder+'0'))
		n /= 10
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
