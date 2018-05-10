package removeDuplicates

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	var cases = []struct {
		raw  []int
		nLen int
	}{
		{
			// duplicate values
			[]int{1, 2, 2},
			2,
		},
	}
	for _, c := range cases {
		nLen, err := removeDuplicates(c.nLen)
		if err != nil {
			t.Errorf("Error removing duplicates: %v\n", err)
		}
		if nLen != c.nLen {
			t.Errorf("Lengths do not match - got %v want: %v\n", nLen, c.nLen)
		}
	}
}
