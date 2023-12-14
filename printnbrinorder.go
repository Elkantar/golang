package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var valeur []int
		chiffre := 0
		compteur := 0
		var unité int
		for n != 0 {
			chiffre = n % 10
			n = n / 10
			valeur = append(valeur, chiffre)
		}
		for compteur2 := range valeur {
			compteur = compteur2 + 1
		}
		for i := 0; i < compteur-1; i++ {
			for j := 0; j < compteur-1; j++ {
				if valeur[j] > valeur[j+1] {
					unité = valeur[j]
					valeur[j] = valeur[j+1]
					valeur[j+1] = unité
				}
			}
		}
		for i := 0; i < compteur; i++ {
			z01.PrintRune(rune(valeur[i] + 48))
		}
	}
}

/*
func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}*/
