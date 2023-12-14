package piscine

func Capitalize(s string) string {
	ar := []rune(s)

	lettre := true

	for i := 0; i < len(s); i++ {
		if AlphaNumeric(ar[i]) == true && lettre {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] = 'A' - 'a' + ar[i]
			}
			lettre = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] = 'a' - 'A' + ar[i]
		} else if AlphaNumeric(ar[i]) == false {
			lettre = true
		}
	}
	return string(ar)
}

func AlphaNumeric(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

/*
func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}*/
