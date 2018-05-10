package iev

import (
	"fmt"
	"testing"
)

func Test_iev(t *testing.T) {
	var cases = []struct {
		fpath string
		want  float64
	}{
		{
			"../input/iev/test_input_01.txt",
			3.5,
		},
		{
			"../input/iev/test_input_02.txt",
			158162.0,
		},
	}
	for _, c := range cases {
		got, err := iev(c.fpath)
		if err != nil {
			t.Errorf("Error retriving iev result: %v", err)
		}
		if got != c.want {
			t.Errorf("values do not match - got %v want: %v\n", got, c.want)
		}
		fmt.Printf("file = %v, prob = %v\n", c.fpath, got)
	}
}
