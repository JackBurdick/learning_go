package subs

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_subs(t *testing.T) {
	var cases = []struct {
		fpath string
		locs  []int
	}{
		{
			"../input/subs/test_input_01.txt",
			[]int{2, 4, 10},
		},
		{
			"../input/subs/test_input_02.txt",
			[]int{8, 28, 176, 203, 233, 256, 263, 359, 366, 373, 380, 397, 404, 461, 471, 486, 493, 512, 519, 597, 629, 656, 663, 677, 686, 709, 740, 747, 783, 802, 809, 824, 855},
		},
	}
	for _, c := range cases {
		locs, err := subs(c.fpath)
		if err != nil {
			t.Errorf("Error retriving locations: %v", err)
		}
		if !reflect.DeepEqual(locs, c.locs) {
			t.Errorf("locations do not match - got %v want: %v\n", locs, c.locs)
		}
		fmt.Printf("file = %v, locations = %v\n", c.fpath, locs)
	}
}
