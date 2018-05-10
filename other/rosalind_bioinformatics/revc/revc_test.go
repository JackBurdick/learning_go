package revc

import (
	"strings"
	"testing"
)

func Test_dna(t *testing.T) {
	var cases = []struct {
		fpath string
		revcS string
	}{
		{
			"../input/revc/test_input_01.txt",
			"ACCGGGTTTT",
		},
	}
	for _, c := range cases {
		revcS, err := revc(c.fpath)
		if err != nil {
			t.Errorf("Error retriving counts: %v", err)
		}
		if strings.ToUpper(revcS) != c.revcS {
			t.Errorf("revc sequence does not match:\nwant:\n%v\n\ngot:\n%v\n\n", c.revcS, strings.ToUpper(revcS))
		}
		//fmt.Printf("file = %v, seq;\n%v\n", c.fpath, strings.ToUpper(revcS))
	}
}
