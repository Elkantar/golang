package piscine

func ForEach(f func(int), a []int) {
	for _, reach := range a {
		f(reach)
	}
}

/*
func PrintNbr(n int) {
	if n < 0 {
		fmt.Print('-')
	} else if n >= 1 && n <= 5 {
		fmt.Print(n)
	} else if n == 6 {
		fmt.Print(n)
		fmt.Print("\n")
	}
}

/*
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}*/
