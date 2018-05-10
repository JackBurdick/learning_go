package sudoku

import (
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	var cases = []struct {
		fpath string
		nWant int
	}{
		{
			"./input/puzzle_01.txt",
			81,
		},
		{
			"./input/puzzle_02.txt",
			81,
		},
		{
			"./input/puzzle_03.txt",
			81,
		},
		{
			"./input/puzzle_04.txt",
			81,
		},
	}
	for _, c := range cases {
		sS, err := solveSudoku(c.fpath)
		if err != nil {
			t.Errorf("Error retriving solved Sudoku: %v", err)
		}

		// count number of solved solutions in returned puzzle
		nGot := 0
		for _, vals := range sS {
			if len(vals) == 1 {
				nGot++
			}
		}

		if nGot != c.nWant {
			t.Errorf("solved Sudoku does not match target\n")
		}
	}
}
