package cosineSimilarity

import (
	"testing"
)

func Test_cosineSimilarity(t *testing.T) {
	var cases = []struct {
		fpath string
		nWant int
	}{
		{
			"../small_input/",
			1,
		},
	}
	for _, c := range cases {
		nGot, err := cosineSimilarity(c.fpath)
		if err != nil {
			t.Errorf("Unable to read directory: %v", err)
		}

		if nGot != c.nWant {
			t.Errorf("whoops\n")
		}
	}
}
