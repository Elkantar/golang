package piscine

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

/*
func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
*/
