package subs

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// subs accepts a file path that contains two sequences (s and t) of dna (not
// exceeding 1 kbp) and returns the locations of t as a substring of s.
func subs(fpath string) ([]int, error) {

	// Read the entire file from the passed input filepath to string.
	ns, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, fmt.Errorf("Cannot read values - %v", err)
	}
	inS := string(ns)
	inS = strings.Trim(inS, "\n")

	// Parse the file contents into the target values; seq_1 (s1) and seq_2 (s2).
	splitS := strings.Split(inS, "\n")
	s := splitS[0]
	s = strings.Trim(s, "\n")
	t := splitS[1]
	t = strings.Trim(t, "\n")

	// Store the last index of the string (s) in a var (lsi) so we don't have to
	// compute it each iteration. li is representative of the last index in the
	// substring.  These two numbers will be compared to ensure the substring
	// can fit in the remaining string.
	lsi := len(s)
	li := 0
	var ans []int
	for i := range s {
		li = i + len(t)
		if li <= lsi {
			if s[i:i+len(t)] == t {
				// note: the required answer format is 1-based array indexing
				ans = append(ans, i+1)
			}
		} else {
			break
		}
	}

	return ans, nil
}
