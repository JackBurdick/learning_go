package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	var cases = []struct {
		target int
		arr    []int
		ans    []int
	}{
		{
			// first + next
			9,
			[]int{2, 7, 11, 15},
			[]int{0, 1},
		},
		{
			// last + prev
			11,
			[]int{3, 3, 3, 3, 3, 2, 9},
			[]int{5, 6},
		},
		{
			// first + last
			11,
			[]int{2, 3, 3, 3, 3, 3, 9},
			[]int{0, 6},
		},
		{
			// middle + middle
			8,
			[]int{2, 3, 4, 4, 3, 2, 9},
			[]int{2, 3},
		},
		{
			// only two
			11,
			[]int{2, 9},
			[]int{0, 1},
		},
	}
	for _, c := range cases {
		got, err := twoSum(c.arr, c.target)
		if err != nil {
			t.Errorf("Error during testing: %v", err)
		}
		if !reflect.DeepEqual(got, c.ans) {
			t.Errorf("twoSum(%v, %v) = %v; want %v", c.arr, c.target, got, c.ans)
		}
	}
}
