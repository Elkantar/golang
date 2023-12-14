package piscine

func IsUpper(s string) bool {
	for _, ch := range []rune(s) {
		if !IsRuneUpper(ch) {
			return false
		}
	}

	return true
}

func IsRuneUpper(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

/*
func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}*/
