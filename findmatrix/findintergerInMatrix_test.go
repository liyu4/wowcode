package ff

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

//  TestFindInterger test case need convert nomarl case, unnomarl case, bad case
func TestFindInterger(t *testing.T) {
	findTest := []struct {
		in     int
		row    int
		column int
		exp    bool
		errExp error
	}{
		// nomarl case
		{10, 0, 5, true, nil},
		// not found case cos input is larger than any number that in matrix
		{100, 0, 5, false, errNotFoundInRow},
		// not found case because input is smaller than  any numebr that in matrix
		{1, 0, 5, false, errNotFoundInColumn},
		// out of matrix range
		{1, -1, 9, false, errOutOfMatrixRange},
		{1, 0, -1, false, errOutOfMatrixRange},
		{1, 10, 0, false, errOutOfMatrixRange},
		{1, 0, 11, false, errOutOfMatrixRange},
	}

	for _, ft := range findTest {
		isFound, err := isExistInMatrix(ft.in, ft.row, ft.column)
		assert.Equal(t, err, ft.errExp)
		assert.Equal(t, isFound, ft.exp)
	}
}
