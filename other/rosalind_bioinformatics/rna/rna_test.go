package rna

import (
	"strings"
	"testing"
)

func Test_dna(t *testing.T) {
	var cases = []struct {
		fpath string
		rnaS  string
	}{
		{
			"../input/rna/test_input_01.txt",
			"GAUGGAACUUGACUACGUAAAUU",
		},
	}
	for _, c := range cases {
		rnaS, err := rna(c.fpath)
		if err != nil {
			t.Errorf("Error retriving counts: %v", err)
		}
		if strings.ToUpper(rnaS) != c.rnaS {
			t.Errorf("rna sequence does not match:\nwant:\n%v\n\ngot:\n%v\n\n", c.rnaS, strings.ToUpper(rnaS))
		}
		//fmt.Printf("file = %v, seq;\n%v\n", c.fpath, strings.ToUpper(rnaS))
	}
}
