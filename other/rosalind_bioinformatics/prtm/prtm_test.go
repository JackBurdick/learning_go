package prtm

import (
	"fmt"
	"testing"
)

func Test_prtm(t *testing.T) {
	var cases = []struct {
		fpath string
		val   float64
	}{
		{
			"../input/prtm/test_input_01.txt",
			821.392,
		},
		{
			"../input/prtm/test_input_02.txt",
			108886.721,
		},
	}
	for _, c := range cases {
		val, err := prtm(c.fpath)
		if err != nil {
			t.Errorf("Error retriving mass: %v", err)
		}

		// Truncate the decimal for comparison.
		gotM := fmt.Sprintf("%.3f", val)
		wantM := fmt.Sprintf("%.3f", c.val)
		if gotM != wantM {
			t.Errorf("mass values do not match - got %v want: %v\n", gotM, wantM)
		}
		fmt.Printf("file = %v, total mass = %v\n", c.fpath, val)
	}
}
