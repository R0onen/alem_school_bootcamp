package cells

import (
	"math/rand"
)

func Replace(h int, w int, field [][]rune) {
	hnew := rand.Intn(len(field))
	wnew := rand.Intn(len(field[0]))
	if field[hnew][wnew] != '*' {
		field[hnew][wnew] = '*'
		field[h][w] = '.'
	} else {
		Replace(h, w, field)
	}
}
