package prot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

// prot accepts a file path to a `.txt` file containing a string of nucleotides
// and returns a protein string encoded by the given nucleotide sequence.
func prot(fpath string) (string, error) {

	// dnaToProt contains a mapping of codon to amino acid.
	dnaToProt := map[string]string{
		"UUU": "F", "UUC": "F", "UUA": "L", "UUG": "L",
		"UCU": "S", "UCC": "S", "UCA": "S", "UCG": "S",
		"UAU": "Y", "UAC": "Y", "UAA": "STOP", "UAG": "STOP",
		"UGU": "C", "UGC": "C", "UGA": "STOP", "UGG": "W",
		"CUU": "L", "CUC": "L", "CUA": "L", "CUG": "L",
		"CCU": "P", "CCC": "P", "CCA": "P", "CCG": "P",
		"CAU": "H", "CAC": "H", "CAA": "Q", "CAG": "Q",
		"CGU": "R", "CGC": "R", "CGA": "R", "CGG": "R",
		"AUU": "I", "AUC": "I", "AUA": "I", "AUG": "M",
		"ACU": "T", "ACC": "T", "ACA": "T", "ACG": "T",
		"AAU": "N", "AAC": "N", "AAA": "K", "AAG": "K",
		"AGU": "S", "AGC": "S", "AGA": "R", "AGG": "R",
		"GUU": "V", "GUC": "V", "GUA": "V", "GUG": "V",
		"GCU": "A", "GCC": "A", "GCA": "A", "GCG": "A",
		"GAU": "D", "GAC": "D", "GAA": "E", "GAG": "E",
		"GGU": "G", "GGC": "G", "GGA": "G", "GGG": "G",
	}

	// read the entire file from the passed input filepath to string.
	ns, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", fmt.Errorf("Cannot read values - %v", err)
	}

	// convert to string, remove '\n', and  convert to upper case letters.
	mrnaS := string(ns)
	mrnaS = strings.Trim(mrnaS, "\n")
	mrnaS = strings.ToUpper(mrnaS)

	seqL := len(mrnaS)
	numCodon := seqL / 3

	// read three chars, convert to amino acid, write to buffer
	var buf bytes.Buffer
	for i := 0; i < numCodon; i++ {

		// Capture 3 base pairs to encode to amino acid.
		si := i * 3
		ei := si + 3
		codon := mrnaS[si:ei]

		aA, ok := dnaToProt[codon]
		if !ok {
			return "", fmt.Errorf("codon does not transcribe - %v", codon)
		}

		// Terminate transcription since a stop codon was encoded.
		if aA == "STOP" {
			return buf.String(), nil
		}

		buf.WriteString(aA)
	}

	// NOTE: stop codon was never received.  Depending on the implementation and
	// requirements, this logic may have to change to include an error or return
	// an empty string.
	return buf.String(), nil
}
