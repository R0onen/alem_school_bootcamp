package main

import (
	"fmt"
	"math/rand"

	"github.com/alem-platform/ap"
)

func main() {
	num := rand.Intn(100)
	for {
		Print("Guess number: ")
		var guess int
		fmt.Scanf("%d", &guess)
		if guess < num {
			Print("Higher\n")
		}
		if guess > num {
			Print("Lower\n")
		}
		if guess == num {
			Print("Match, you win!\n")
			return
		}
	}
}

func Print(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}
