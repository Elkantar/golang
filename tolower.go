package piscine

func ToLower(s string) string {
	sentence := []rune(s)
	for ch := range sentence {
		if sentence[ch] >= 'A' && sentence[ch] <= 'Z' {
			sentence[ch] = sentence[ch] + 32
		}
	}
	return string(sentence)
}

/*
func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}*/
