package piscine

func UltimateDivMod(a *int, b *int) {
	tempA := *a / *b
	tempB := *a % *b
	*a = tempA
	*b = tempB
}

/*
func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}*/
