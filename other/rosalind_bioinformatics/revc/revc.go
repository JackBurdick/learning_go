package revc

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// revc accepts a file path to a .txt file of nucleotides and returns a string
// of the reverse complement of the included nucleotides.
func revc(fpath string) (string, error) {

	// read the entire file to string from the passed input filepath.
	ns, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", fmt.Errorf("Cannot read values - %v", err)
	}

	// convert to string, remove '\n', and  convert to lower case letters.
	dnaS := string(ns)
	dnaS = strings.Trim(dnaS, "\n")
	dnaS = strings.ToLower(dnaS)

	// reverse DNA strand.
	chars := []rune(dnaS)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	dnaS = string(chars)

	// complement DNA strand.
	chars = []rune(dnaS)
	for i := 0; i < len(chars); i++ {
		switch chars[i] {
		case 'a':
			chars[i] = 't'
		case 't':
			chars[i] = 'a'
		case 'c':
			chars[i] = 'g'
		case 'g':
			chars[i] = 'c'
		}
	}
	dnaS = string(chars)

	return dnaS, nil
}
