package piscine // main

func IsPrintable(s string) bool {
	sentence := []rune(s)
	for _, ch := range sentence {
		if isprint(ch) {
			return false
		}
	}

	return true
}

func isprint(c rune) bool {
	return c >= 0 && c <= 31 // table ascii
}

/*
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
	fmt.Println(IsPrintable("\t\f\v\b\v\b\b\v")) // false
	fmt.Println(IsPrintable("^z9U=L:^x<><t"))    // true
}*/
