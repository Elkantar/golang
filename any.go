package piscine

func Any(f func(string) bool, a []string) bool {
	for _, any := range a {
		if f(any) {
			return true
		}
	}
	return false
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
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
*/
