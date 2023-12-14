package main

import (
	"fmt"
	"os"
)

func main() {
	var runes []rune
	var res string
	motclé := []string{"01", "galaxy 01", "galaxy"}
	count := 0
	for i := 1; i < len(os.Args); i++ {
		runes = []rune(os.Args[i])
		res = ""
		for j := 0; j < len(runes); j++ {
			res = res + string(runes[j])
		}
		for _, ch := range motclé {
			if res == ch {
				fmt.Println("Alert!!!")
				count++
			}
		}
		if count == 1 {
			break
		}
	}
}
