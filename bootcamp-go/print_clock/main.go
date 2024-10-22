package main

import "github.com/alem-platform/ap"

func main() {
	h1 := '0'
	trans := 0

	for i := 0; i < 3; i++ {
		h2 := '0'
		for j := 0; j < 10-trans; j++ {
			m3 := '0'
			for k := 0; k < 6; k++ {
				m4 := '0'
				for l := 0; l < 10; l++ {
					ap.PutRune(h1)
					ap.PutRune(h2)
					ap.PutRune(':')
					ap.PutRune(m3)
					ap.PutRune(m4)
					ap.PutRune('\n')
					m4++
				}
				m3++
			}
			if h1 == '2' {
				trans = 6
			}
			h2++
		}
		h1++
	}
}
