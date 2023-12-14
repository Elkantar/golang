package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	length := 0

	for index := range params { // Lis la taille du "mot"
		length = index
	}

	for ch := 1; ch <= length; ch++ { // ch = 1 pour commencé au premier caractère puis incrémente jusqu'a la taille de l'index
		for _, i := range params[ch] {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
