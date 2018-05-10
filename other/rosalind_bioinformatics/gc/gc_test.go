package gc

import (
	"fmt"
	"testing"
)

func Test_gc(t *testing.T) {
	var cases = []struct {
		fPath string
		ID    string
		gcF   float64
	}{
		{
			"../input/gc/test_input_01.txt",
			"Rosalind_0808",
			60.919540,
		},
		{
			"../input/gc/test_input_02.txt",
			"Rosalind_2824",
			50.672646,
		},
	}
	for _, c := range cases {
		got, err := gc(c.fPath)
		if err != nil {
			t.Errorf("Error retriving gc result: %v", err)
		}
		if got.ID != c.ID {
			t.Errorf("IDs do not match- got %s want: %s", got.ID, c.ID)
		}

		// Truncate the decimal for comparison.
		gotP := fmt.Sprintf("%.6f", got.gcC)
		wantP := fmt.Sprintf("%.6f", c.gcF)
		if gotP != wantP {
			t.Errorf("GC content does not match - got %v want: %v", gotP, wantP)
		}
	}
}
