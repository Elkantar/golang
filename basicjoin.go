package piscine

func BasicJoin(elems []string) string {
	concat := ""
	compteur := 0
	for range elems {
		compteur = compteur + 1
	}

	for i := range elems {
		if i == (compteur - 1) {
			concat += elems[i]
		} else {
			concat += elems[i]
		}
	}
	return concat
}

/*
func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}*/
