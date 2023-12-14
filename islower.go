package piscine

func IsLower(s string) bool {
	for _, ch := range s {
		if !IsRunelower(ch) {
			return false
		}
	}
	return true
}

func IsRunelower(c rune) bool {
	return c >= 'a' && c <= 'z'
}

/*
func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}*/
