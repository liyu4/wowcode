package ff

import (
	"errors"
	"fmt"
)

// a array of arrays is know as 2D array
// the two dimensional array in c programming language
// is also know as matrix. A matrix can be represented a table of rows
// and columns.
var matrix = [][]int{
	{2, 4, 5, 7, 8, 9},
	{3, 4, 5, 8, 9, 11},
	{7, 8, 9, 10, 11, 12},
	{8, 11, 12, 13, 14, 16},
}

var errOutOfMatrixRange = errors.New("out of matrix range")
var errNotFoundInColumn = errors.New("not found in column")
var errNotFoundInRow = errors.New("not found in row")

func isExistInMatrix(in int, row int, column int) (bool, error) {
	for {
		if (row < 0 || column < 0) || (row > 3 || column > 5) {
			return false, errOutOfMatrixRange
		}

		ele := matrix[row][column]
		// if in gt ele
		if in == ele {
			fmt.Printf("find postion %v %v value:%v", row, column, matrix[row][column])
			return true, nil
		}
		if in < ele {
			column--
			if column <= 0 {
				return false, errNotFoundInColumn
			}
		} else {
			row++
			if row >= 4 {
				return false, errNotFoundInRow
			}
		}
	}
}
