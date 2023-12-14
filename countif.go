package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, countif := range tab {
		if f(countif) {
			count++
		}
	}
	return count
}

/*
func IsNumeric(s string) bool {
	for _, ch := range s {
		if !IsRuneNumbers(ch) {
			return false
		}
	}

	return true
}

func IsRuneNumbers(c rune) bool {
	return c >= '0' && c <= '9'
}

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
*/
