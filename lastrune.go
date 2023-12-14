package piscine

func LastRune(s string) rune {
	sentence := []rune(s)
	counter := 0
	for letter := range sentence {
		counter = letter + 1
	}
	return sentence[counter-1]
}

/*
func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}*/
