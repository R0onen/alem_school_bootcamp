package bootcamp

func GetMathOperation(op string) *func(int, int) int {
	var fn func(int, int) int

	switch op {
	case "add":
		fn = func(a, b int) int { return a + b }
	case "subtract":
		fn = func(a, b int) int { return a - b }
	case "multiply":
		fn = func(a, b int) int { return a * b }
	case "divide":
		fn = func(a, b int) int {
			if b == 0 {
				return 0 // Handle division by zero
			}
			return a / b
		}
	default:
		return nil
	}

	return &fn
}
