package main

import (
	"fmt"
	"os"
)

// vérification d'erreur conditionnelle
func main() {
	arr := os.Args[1:]
	if checkErrors(arr) {
		fmt.Println("Error")
	} else {
		sudoku := optimizeSudoku(arr)
		if solveSudoku(&sudoku, len(sudoku)) {
			printSudoku(sudoku)
		} else {
			fmt.Println("Error")
		}
	}
}

func checkErrors(arr []string) bool { // si le parametre count != 9, alors une erreur est retournée alors la présence d'une erreur est retournée
	count := '1'
	for i, str := range arr {
		if len(str) != 9 {
			return true
		}
		count++
		for _, ch := range arr[i] {
			count++
			if !(ch >= '1' && ch <= '9' || ch == '.') {
				return true
			}
			count++
		}
	}
	return false
}

func solveSudoku(arr *[][]int, length int) bool { // vérifier 0 dans le sudoku
	isEmpty := true
	row := -1 // je ne connais pas l'emplacement 0
	column := -1
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if (*arr)[i][j] == 0 {
				row = i
				column = j // zéro trouvé
				isEmpty = false
				break
			}
		}
	}
	if isEmpty { // atteint la fin de la fonction et a fonctionné
		return true
	}
	for number := 1; number <= 9; number++ { // si number <= 9
		if isCorrect(*arr, row, column, number) {
			(*arr)[row][column] = number // ajoute le nombre fini
			if solveSudoku(arr, length) {
				return true
			} else {
				(*arr)[row][column] = 0
			}
		}
	}
	return false
}

func optimizeSudoku(arr []string) [][]int { // string convertit en un tableau int à deux dimensions
	sudoku := make([][]int, 9)
	for i := range sudoku {
		sudoku[i] = make([]int, 9)
	}
	for i, str := range arr {
		for j, ch := range str {
			sudoku[i][j] = runeToInt(ch)
		}
	}
	return sudoku
}

func printSudoku(arr [][]int) { // résultat final
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func isCorrect(arr [][]int, row int, column int, num int) bool { // vérifier dans une ligne, dans une colonne et dans un mini-carré les doublons
	return !checkRow(arr, row, num) && !checkColumn(arr, column, num) && !checkSubSudoku(arr, row-(row%3), column-(column%3), num)
}

func checkRow(arr [][]int, row int, num int) bool { // vérification horizontale
	for column := 0; column < len(arr); column++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

func checkColumn(arr [][]int, column int, num int) bool { // vérifier dans la colonne
	for row := 0; row < len(arr); row++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

func checkSubSudoku(arr [][]int, rowIndex int, columnIndex int, num int) bool { // vérifier dans le mini carré
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if arr[rowIndex+row][columnIndex+column] == num {
				return true
			}
		}
	}
	return false
}

func runeToInt(number rune) int { // compteur (retour au type int)
	count := 0
	for i := '1'; i <= number; i++ {
		count++
	}
	return count
}

//  go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
