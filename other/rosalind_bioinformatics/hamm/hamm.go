package hamm

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// hamm accepts a file path that contains two sequences of dna of equal length
// (not exceeding 1 kbp) and returns the hamming distance between the two
// seqences.  The hamming distance is the number of corresponding symbols that
// differ in the two sequences.
// read more here: https://en.wikipedia.org/wiki/Hamming_distance
func hamm(fpath string) (int, error) {

	// Read the entire file from the passed input filepath to string.
	ns, err := ioutil.ReadFile(fpath)
	if err != nil {
		return 0, fmt.Errorf("Cannot read values - %v", err)
	}
	inS := string(ns)
	inS = strings.Trim(inS, "\n")

	// Parse the file contents into the target values; seq_1 (s1) and seq_2 (s2).
	splitS := strings.Split(inS, "\n")
	s1 := splitS[0]
	s2 := splitS[1]

	d := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			d++
		}
	}

	return d, nil
}
