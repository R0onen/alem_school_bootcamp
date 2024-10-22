package bootcamp

import "github.com/alem-platform/ap"

// WrapperPrintStr prints the string before calling the original function.
func WrapperPrintStr(fn func(str *string)) func(str *string) {
	return func(str *string) {
		for _, r := range *str {
			ap.PutRune(r) // Print each character
		}
		ap.PutRune('\n') // Print a newline
		fn(str)          // Call the original function
	}
}

// WrapperRot1 applies the ROT1 transformation to the string before calling the original function.
func WrapperRot1(fn func(str *string)) func(str *string) {
	return func(str *string) {
		runes := []rune(*str)
		for i := range runes {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = ((runes[i]-'a'+1)%26 + 'a')
			}
		}
		*str = string(runes) // Convert back to string
		fn(str)              // Call the original function
	}
}

// WrapperRot13 applies the ROT13 transformation to the string before calling the original function.
func WrapperRot13(fn func(str *string)) func(str *string) {
	return func(str *string) {
		runes := []rune(*str)
		for i := range runes {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = ((runes[i]-'a'+13)%26 + 'a')
			}
		}
		*str = string(runes) // Convert back to string
		fn(str)              // Call the original function
	}
}

// WrapperReverseStr reverses the string before calling the original function.
func WrapperReverseStr(fn func(str *string)) func(str *string) {
	return func(str *string) {
		runes := []rune(*str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		*str = string(runes) // Convert back to string
		fn(str)              // Call the original function
	}
}

// WrapFunctions takes a slice of wrapper functions and returns a single function.
func WrapFunctions(decs []func(fn func(str *string)) func(str *string)) func(str *string) {
	return func(str *string) {
		wrappedFn := func(str *string) {}
		for i := len(decs) - 1; i >= 0; i-- {
			wrappedFn = decs[i](wrappedFn)
		}
		wrappedFn(str) // Call the wrapped function
	}
}
