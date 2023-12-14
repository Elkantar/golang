package piscine

func Unmatch(a []int) int {
	for _, res := range a {
		count := 0
		for _, ch := range a {
			if ch == res {
				count++
			}
		}
		if count == 1 || count%2 == 1 {
			return res
		}
	}
	return -1
}

/*
func main() {
	a := []int{1, 1, 2, 3, 4, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}*/
