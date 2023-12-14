package piscine

func NRune(s string, n int) rune {
	sentence := []rune(s)
	counter := 1

	for letter := range sentence {
		counter = letter
	}
	if n-1 >= 0 && n-1 <= counter {
		return sentence[n-1]
	}
	return 0
}

/*
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}*/
