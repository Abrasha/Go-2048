package main

import "fmt"

//import "math/rand"

type Field struct {
	matrix [6][6]int
	score  int
	moved  bool
	over   bool
}

func show(field Field) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if i == 0 || i == 5 || j == 0 || j == 5 {
				fmt.Print("#\t")
			} else if field.matrix[i][j] == 0 {
				fmt.Print("\t")
			} else {
				fmt.Printf("%d\t", field.matrix[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println("score: ", field.score)
}

func over(field Field) bool {
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			current := field.matrix[i][j]
			if current == 0 ||
				field.matrix[i][j-1] == current ||
				field.matrix[i-1][j] == current ||
				field.matrix[i][j+1] == current ||
				field.matrix[i+1][j] == current {
				return false
			}
		}
	}
	return true
}

func fill(field *Field) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if i == 0 || i == 5 || j == 0 || j == 5 {
				(*field).matrix[i][j] = -1
			} else { //just for tests
				(*field).matrix[i][j] = 2 //replace by random generation
			} //in real project
		}
	}
	// generate here

}

func left(field *Field) {
	for i := 1; i < 5; i++ {
		for j := 2; j < 5; j++ {
			if field.matrix[i][j] > 0 {
				for field.matrix[i][j-1] == 0 {
					field.matrix[i][j-1] = field.matrix[i][j]
					field.matrix[i][j] = 0
					j--
					field.moved = true
				}
			}
		}
		for j := 2; j < 5; j++ {
			if field.matrix[i][j] == field.matrix[i][j-1] && field.matrix[i][j] != 0 {
				field.matrix[i][j-1] = 2 * field.matrix[i][j-1]
				field.score += field.matrix[i][j-1]
				field.moved = true
				for k := j; k < 4; k++ {
					field.matrix[i][k] = field.matrix[i][k+1]
				}
				field.matrix[i][4] = 0
			}
		}
	}
}
func right(field *Field) {
	for i := 1; i < 5; i++ {
		for j := 3; j > 0; j-- {
			if field.matrix[i][j] > 0 {
				for field.matrix[i][j+1] == 0 {
					field.matrix[i][j+1] = field.matrix[i][j]
					field.matrix[i][j] = 0
					j++
					field.moved = true
				}
			}
		}
		for j := 3; j > 0; j-- {
			if field.matrix[i][j] == field.matrix[i][j+1] && field.matrix[i][j] != 0 {
				field.matrix[i][j+1] = 2 * field.matrix[i][j+1]
				field.score += field.matrix[i][j+1]
				field.moved = true
				for k := j; k > 1; k-- {
					field.matrix[i][k] = field.matrix[i][k-1]
				}
				field.matrix[i][1] = 0
			}
		}
	}
}
func up(field *Field) {
	for j := 1; j < 5; j++ {
		for i := 2; i < 5; i++ {
			if field.matrix[i][j] > 0 {
				for field.matrix[i-1][j] == 0 {
					field.matrix[i-1][j] = field.matrix[i][j]
					field.matrix[i][j] = 0
					i--
					field.moved = true
				}
			}
		}
		for i := 2; i < 5; i++ {
			if field.matrix[i][j] == field.matrix[i-1][j] && field.matrix[i][j] != 0 {
				field.matrix[i-1][j] = 2 * field.matrix[i-1][j]
				field.score += field.matrix[i-1][j]
				field.moved = true
				for k := i; k < 4; k++ {
					field.matrix[k][j] = field.matrix[k+1][j]
				}
				field.matrix[4][j] = 0
			}
		}
	}
}

func down(field *Field) {
	for j := 1; j < 5; j++ {
		for i := 3; i > 0; i-- {
			if field.matrix[i][j] > 0 {
				for field.matrix[i+1][j] == 0 {
					field.matrix[i+1][j] = field.matrix[i][j]
					field.matrix[i][j] = 0
					i++
					field.moved = true
				}
			}
		}
		for i := 3; i > 0; i-- {
			if field.matrix[i][j] == field.matrix[i+1][j] && field.matrix[i][j] != 0 {
				field.matrix[i+1][j] = 2 * field.matrix[i+1][j]
				field.score += field.matrix[i+1][j]
				field.moved = true
				for k := i; k > 1; k-- {
					field.matrix[k][j] = field.matrix[k-1][j]
				}
				field.matrix[1][j] = 0
			}
		}
	}
}
func main() {
	var x Field
	fill(&x)
	show(x)
	right(&x)
	show(x)

}
