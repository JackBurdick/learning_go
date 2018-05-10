package dna

import (
	"fmt"
	"testing"
)

func Test_dna(t *testing.T) {
	var cases = []struct {
		fpath string
		aC    int
		cC    int
		gC    int
		tC    int
	}{
		{
			"../input/dna/test_input_01.txt",
			20,
			12,
			17,
			21,
		},
		{
			"../input/dna/test_input_02.txt",
			236,
			221,
			207,
			219,
		},
	}
	for _, c := range cases {
		rs, err := dna(c.fpath)
		if err != nil {
			t.Errorf("Error retriving counts: %v", err)
		}
		if rs.aC != c.aC {
			t.Errorf("A counts do not match - got %v want: %v\n", rs.aC, c.aC)
		}
		if rs.cC != c.cC {
			t.Errorf("C counts do not match - got %v want: %v\n", rs.cC, c.cC)
		}
		if rs.gC != c.gC {
			t.Errorf("G counts do not match - got %v want: %v\n", rs.gC, c.gC)
		}
		if rs.tC != c.tC {
			t.Errorf("T counts do not match - got %v want: %v\n", rs.tC, c.tC)
		}
		fmt.Printf("file = %v, counts = %v\n", c.fpath, rs)
	}
}
