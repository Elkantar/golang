package piscine

func Map(f func(int) bool, a []int) []bool {
	maps := make([]bool, len(a))
	for Key, reach := range a {
		maps[Key] = f(reach)
	}
	return maps
}

/*
func IsPrime(nbr int) bool {
	rep := true
	if nbr == 2 || nbr == 3 || nbr == 5 || nbr == 7 {
		rep = true
	} else if nbr == 1 || nbr == 0 {
		rep = false
	} else {
		for i := 2; i < 40; i++ {
			if nbr%i == 0 {
				rep = false
			}
		}
	}
	return rep
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
*/
