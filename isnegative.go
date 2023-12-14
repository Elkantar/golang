package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	valtrue := 'T'
	valfalse := 'F'
	if nb >= 0 {
		z01.PrintRune(valfalse)
		z01.PrintRune(10)
	} else {
		z01.PrintRune(valtrue)
		z01.PrintRune(10)
	}
}
