package bootcamp

func ChainFunctions(funcs []func(*string)) func(*string) {
	return func(str *string) {
		for _, fn := range funcs {
			fn(str)
		}
	}
}
