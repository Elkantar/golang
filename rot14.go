package piscine

func Rot14(s string) string {
	Rune := []rune(s)
	var res string

	for i := 0; i < len(Rune); i++ {
		if Rune[i] >= 'a' && Rune[i] <= 'z' {
			if Rune[i] >= 'm' {
				Rune[i] = Rune[i] - 12
			} else {
				Rune[i] = Rune[i] + 14
			}
		} else if Rune[i] >= 'A' && Rune[i] <= 'Z' {
			if Rune[i] >= 'M' {
				Rune[i] = Rune[i] - 12
			} else {
				Rune[i] = Rune[i] + 14
			}
		}
		res += string(Rune[i])
	}
	return res
}

/*
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
*/
