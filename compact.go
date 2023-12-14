package piscine

func Compact(ptr *[]string) int {
	var res []string
	count1 := 0
	for _, ch := range *ptr {
		if ch != "" {
			res = append(res, ch)
			count1++
		}
	}
	*ptr = res
	return count1
}
