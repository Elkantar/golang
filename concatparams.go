package piscine

func ConcatParams(args []string) string {
	concat := ""
	compteur := 0
	for range args {
		compteur = compteur + 1
	}

	for i := range args {
		if i == (compteur - 1) {
			concat = concat + args[i]
		} else {
			concat = concat + args[i] + "\n"
		}
	}
	return concat
}

/*
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}*/
