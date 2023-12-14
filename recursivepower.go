package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		// resultat := 1
		if power > 1 {
			return nb * RecursivePower(nb, power-1)
		}

		return nb
	}
}

/*
func main() {
	fmt.Println(RecursivePower(4, 3))
}*/
