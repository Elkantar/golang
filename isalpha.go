package piscine // main

func IsAlpha(s string) bool {
	for _, ch := range s {
		if !IsRuneAlpha(ch) && !IsRuneUp(ch) && !IsRuneNumber(ch) {
			return false
		} else if IsRuneAlpha(ch) && IsRuneUp(ch) && IsRuneNumber(ch) {
			return true
		}
	}

	return true
}

func IsRuneAlpha(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func IsRuneUp(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func IsRuneNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

/*
func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
	fmt.Println(IsAlpha("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")) // true
	fmt.Println(IsAlpha("e5?FH)N1qf:zc"))                                                  // false
}*/
