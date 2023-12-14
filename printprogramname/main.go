package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	runes := []rune(name)
	for ch := range runes {
		if name[ch] == '.' || name[ch] == '/' {
			ch++
		} else {
			z01.PrintRune(runes[ch])
		}
	}
	z01.PrintRune('\n')
}
