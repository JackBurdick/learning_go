package rna

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func rna(fPath string) (string, error) {

	// read entire file into memory
	ns, err := ioutil.ReadFile(fPath)
	if err != nil {
		return "", fmt.Errorf("Cannot read values - %v", err)
	}

	// convert to string, remove '\n', and  convert to lower case letters.
	dnaS := string(ns)
	dnaS = strings.Trim(dnaS, "\n")
	dnaS = strings.ToLower(dnaS)

	// replace all "t" nucleotides with "u" nucleotides.
	rnaS := strings.Replace(dnaS, "t", "u", -1)

	return rnaS, nil
}
