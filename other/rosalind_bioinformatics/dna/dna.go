package dna

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type nucCounts struct {
	aC int
	cC int
	gC int
	tC int
}

// String implements the stringer.
func (nC *nucCounts) String() string {

	// Example output will look similar to `20 12 17 21` which conforms to
	// rosalind's specified output `A, C, G, T`
	return fmt.Sprintf("%v %v %v %v", nC.aC, nC.cC, nC.gC, nC.tC)
}

// dna accepts a file path to a `.txt` file containing a nucleotide sequence and
// returns a structure containing the counts of each nucleotide.
// TODO: input chars should be buffered
func dna(fpath string) (*nucCounts, error) {

	// read the entire file to string from the passed input filepath.
	ns, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, fmt.Errorf("Cannot read values - %v", err)
	}

	counts := &nucCounts{}

	// convert to string, remove '\n', and  convert to lower case letters.
	dnaS := string(ns)
	dnaS = strings.Trim(dnaS, "\n")
	dnaS = strings.ToLower(dnaS)

	// count occurances of each nucleotide in the sequence.
	counts.aC = strings.Count(dnaS, "a")
	counts.cC = strings.Count(dnaS, "c")
	counts.gC = strings.Count(dnaS, "g")
	counts.tC = strings.Count(dnaS, "t")

	return counts, nil
}
