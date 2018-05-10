package iprb

import (
	"fmt"
	"testing"
)

func Test_iprb(t *testing.T) {
	var cases = []struct {
		fpath string
		prob  float64
	}{
		{
			"../input/iprb/test_input_01.txt",
			0.7833333333333333,
		},
		{
			"../input/iprb/test_input_02.txt",
			0.6837329337329338,
		},
	}
	for _, c := range cases {
		prob, err := iprb(c.fpath)
		if err != nil {
			t.Errorf("Error retriving difference: %v", err)
		}
		if prob != c.prob {
			t.Errorf("probability values do not match - got %v want: %v\n", prob, c.prob)
		}
		fmt.Printf("file = %v, prob = %v\n", c.fpath, prob)
	}
}
