package prtm

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// createMap creates a map of monoiostopic masses of amino acids.  The values
// are obtained from: http://rosalind.info/glossary/monoisotopic-mass-table/
func createMap() map[string]float64 {
	mM := map[string]float64{
		"A": 71.03711,
		"C": 103.00919,
		"D": 115.02694,
		"E": 129.04259,
		"F": 147.06841,
		"G": 57.02146,
		"H": 137.05891,
		"I": 113.08406,
		"K": 128.09496,
		"L": 113.08406,
		"M": 131.04049,
		"N": 114.04293,
		"P": 97.05276,
		"Q": 128.05858,
		"R": 156.10111,
		"S": 87.03203,
		"T": 101.04768,
		"V": 99.06841,
		"W": 186.07931,
		"Y": 163.06333,
	}

	return mM
}

func prtm(fPath string) (float64, error) {
	var mass float64

	// Read in and clean content from file.
	fC, err := ioutil.ReadFile(fPath)
	if err != nil {
		return mass, fmt.Errorf("Cannot read values - %v", err)
	}
	fS := string(fC)
	fS = strings.Trim(fS, "\n")

	mM := createMap()

	// Calculate total mass by summing individual component masses.
	for _, c := range fS {
		m, ok := mM[string(c)]
		if !ok {
			return mass, fmt.Errorf("mass of %v unknown", string(c))
		}
		mass += m
	}

	return mass, nil
}
