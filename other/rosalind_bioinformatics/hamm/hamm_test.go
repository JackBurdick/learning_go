package hamm

import (
	"fmt"
	"testing"
)

func Test_hamm(t *testing.T) {
	var cases = []struct {
		fpath string
		diff  int
	}{
		{
			"../input/hamm/test_input_01.txt",
			462,
		},
		{
			"../input/hamm/test_input_02.txt",
			7,
		},
	}
	for _, c := range cases {
		diff, err := hamm(c.fpath)
		if err != nil {
			t.Errorf("Error retriving difference: %v", err)
		}
		if diff != c.diff {
			t.Errorf("difference values do not match - got %v want: %v\n", diff, c.diff)
		}
		fmt.Printf("file = %v, counts = %v\n", c.fpath, diff)
	}
}
