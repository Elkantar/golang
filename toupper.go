package piscine

func ToUpper(s string) string {
	sentence := []rune(s)
	for ch := range sentence {
		if sentence[ch] >= 'a' && sentence[ch] <= 'z' {
			sentence[ch] = sentence[ch] - 32
		}
	}
	return string(sentence)
}

/*
func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}*/
