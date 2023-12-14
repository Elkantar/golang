package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func SplitWhiteSpaces(s string) []string {
	var out []string = []string{}
	var current string = ""

	for _, char := range s {
		if char == '\n' || char == '\t' || char == ' ' {
			if current != "" {
				out = append(out, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}

	if current != "" {
		out = append(out, current)
	}

	return out
}

func main() {
	arg := os.Args

	length := 0

	for range arg {
		length++
	}
	strings := SplitWhiteSpaces(arg[length-1])
	taille := len(strings)
	for i := 0; i < taille; i++ {
		if strings[taille-1] == "" {
			taille -= 1
		} else {
			break
		}
	}

	PrintStr(strings[taille-1])
	z01.PrintRune('\n')
}
