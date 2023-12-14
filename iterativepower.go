package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		resultat := 1
		for i := 0; i < power; i++ {
			resultat = resultat * nb
		}
		/*for x := nb - 1; x >= 1; x-- {
			fmt.Print("power = ")
			fmt.Println(power)
			power = power * x
		}*/
		return resultat
	}
}

/*
func main() {
	fmt.Println(IterativePower(-2, 5))
}*/
