package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	taille := 0

	for index := range params {
		taille = index
	}

	for ch := taille; ch >= 1; ch-- {
		for _, i := range params[ch] {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
