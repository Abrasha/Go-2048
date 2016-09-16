package main

import "fmt"

//import "math/rand"

const (
	FIELD_HEIGHT = 6
	FIELD_WIDTH = 6
	DIR_TOP = 0
	DIR_RIGHT = 1
	DIR_BOTTOM = 2
	DIR_LEFT = 3
)

type Field struct {
	matrix [FIELD_HEIGHT][FIELD_WIDTH]int
	score  int
	moved  bool
	over   bool
}

func isEdge(col int, row int) bool {
	if col == 0 || col == (FIELD_HEIGHT - 1) || row == 0 || row == (FIELD_WIDTH - 1) {
		return true
	}
	return false
}
func canMoveLeft(field Field, col int, row int) bool {
	if field.matrix[col][row] == field.matrix[col][row - 1] && field.matrix[col][row] != 0 {
		return true
	}
	return false
}

func canMoveRight(field Field, col int, row int) bool {
	if field.matrix[col][row] == field.matrix[col][row + 1] && field.matrix[col][row] != 0 {
		return true
	}
	return false
}
func canMoveUp(field Field, col int, row int) bool {
	if field.matrix[col][row] == field.matrix[col - 1][row] && field.matrix[col][row] != 0 {
		return true
	}
	return false
}
func canMoveDown(field Field, col int, row int) bool {
	if field.matrix[col][row] == field.matrix[col + 1][row] && field.matrix[col][row] != 0 {
		return true
	}
	return false
}

func unmovable(field Field, col int, row int, current int) bool {
	if current == 0 ||
		field.matrix[col][row - 1] == current ||
		field.matrix[col - 1][row] == current ||
		field.matrix[col][row + 1] == current ||
		field.matrix[col + 1][row] == current {
		return false
	}

	return true
}

func show(field Field) {
	for col := 0; col < FIELD_HEIGHT; col++ {
		for row := 0; row < FIELD_WIDTH; row++ {
			if isEdge(col, row) {
				fmt.Print("##\t")
			} else if field.matrix[col][row] == 0 {
				fmt.Print("__\t")
			} else {
				fmt.Printf("%2d\t", field.matrix[col][row])
			}
		}
		fmt.Println()
	}
	fmt.Println("score: ", field.score)
}

func over(field Field) bool {
	for col := 1; col < 5; col++ {
		for row := 1; row < 5; row++ {
			current := field.matrix[col][row]
			if unmovable(field, col, row, current) == false {
				return false
			}
		}
	}
	return true
}

func fill(field *Field) {
	for col := 0; col < 6; col++ {
		for row := 0; row < 6; row++ {
			if col == 0 || col == 5 || row == 0 || row == 5 {
				(*field).matrix[col][row] = -1
			} else {
				//just for tests
				(*field).matrix[col][row] = 2 //replace by random generation
			} //in real project
		}
	}
	// generate here

}

func left(field *Field) {
	for col := 1; col < 5; col++ {
		for row := 2; row < 5; row++ {
			if field.matrix[col][row] > 0 {
				for field.matrix[col][row - 1] == 0 {
					field.matrix[col][row - 1] = field.matrix[col][row]
					field.matrix[col][row] = 0
					row--
					field.moved = true
				}
			}
		}
		for j := 2; j < 5; j++ {
			if canMoveLeft(*field, col, j) {
				field.matrix[col][j - 1] = 2 * field.matrix[col][j - 1]
				field.score += field.matrix[col][j - 1]
				field.moved = true
				for k := j; k < 4; k++ {
					field.matrix[col][k] = field.matrix[col][k + 1]
				}
				field.matrix[col][4] = 0
			}
		}
	}
}
func right(field *Field) {
	for i := 1; i < 5; i++ {
		for j := 3; j > 0; j-- {
			if field.matrix[i][j] > 0 {
				for field.matrix[i][j + 1] == 0 {
					field.matrix[i][j + 1] = field.matrix[i][j]
					field.matrix[i][j] = 0
					j++
					field.moved = true
				}
			}
		}
		for j := 3; j > 0; j-- {
			if canMoveRight(*field, i, j) {
				field.matrix[i][j + 1] = 2 * field.matrix[i][j + 1]
				field.score += field.matrix[i][j + 1]
				field.moved = true
				for k := j; k > 1; k-- {
					field.matrix[i][k] = field.matrix[i][k - 1]
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
				for field.matrix[i - 1][j] == 0 {
					field.matrix[i - 1][j] = field.matrix[i][j]
					field.matrix[i][j] = 0
					i--
					field.moved = true
				}
			}
		}
		for i := 2; i < 5; i++ {
			if canMoveUp(*field, i, j) {
				field.matrix[i - 1][j] = 2 * field.matrix[i - 1][j]
				field.score += field.matrix[i - 1][j]
				field.moved = true
				for k := i; k < 4; k++ {
					field.matrix[k][j] = field.matrix[k + 1][j]
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
				for field.matrix[i + 1][j] == 0 {
					field.matrix[i + 1][j] = field.matrix[i][j]
					field.matrix[i][j] = 0
					i++
					field.moved = true
				}
			}
		}
		for i := 3; i > 0; i-- {
			if canMoveDown(*field, i, j) {
				field.matrix[i + 1][j] = 2 * field.matrix[i + 1][j]
				field.score += field.matrix[i + 1][j]
				field.moved = true
				for k := i; k > 1; k-- {
					field.matrix[k][j] = field.matrix[k - 1][j]
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
