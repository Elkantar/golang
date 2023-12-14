package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, letter := range s {
		if alpha(letter) {
			counter++
		}
	}
	return counter
}

func alpha(a rune) bool {
	for x := 'a'; x <= 'z'; x++ {
		if a == x {
			return true
		}
	}
	for x := 'A'; x <= 'Z'; x++ {
		if a == x {
			return true
		}
	}
	return false
}

/*
func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
*/
