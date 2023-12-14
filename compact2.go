package piscine

func Compact2(ptr *[]string) int {
	var res []string
	count := 0
	for _, ch := range *ptr {
		if ch != "" {
			res = append(res, ch)
			count++
		}
	}
	*ptr = res
	return count
}
