package bootcamp

func NthRune(s string, n int) rune {
 if n <= 0 || n > len(s) {
  return -1
 }
 runes := []rune(s)
 return runes[n-1]
}

func TestNthRune(fn func(s string, n int) rune) bool {
 tests := []struct {
  inputString string
  n           int
  expected    rune
 }{
  {"hello", 1, 'h'},
  {"hello", 5, 'o'},
  {"hello", 3, 'l'},
  {"", 1, -1},
  {"hello", 0, -1},
  {"hello", 6, -1},
  {"go", 2, 'o'},
  {"Привет", 3, 'и'},
 }

 for _, test := range tests {
  result := fn(test.inputString, test.n)
  if result != test.expected {
   return false
  }
 }
 return true
}

func ZeroRune(s string, n int) rune {
 return '0'
}

func FirstRune(s string, n int) rune {
 if len(s) > 0 {
  return rune(s[0])
 }
 return -1
}