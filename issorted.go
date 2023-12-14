package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	if len(tab) > 1 {
		if f(tab[0], tab[1]) >= 0 {
			for i := 0; i < len(tab)-1; i++ {
				if f(tab[i], tab[i+1]) < 0 {
					return false
				}
			}
		}
		if f(tab[0], tab[1]) <= 0 {
			for i := 0; i < len(tab)-1; i++ {
				if f(tab[i], tab[i+1]) > 0 {
					return false
				}
			}
		}
	}

	return true
}

/*
func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
*/
