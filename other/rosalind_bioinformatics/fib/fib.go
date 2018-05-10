package fib

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// calcFib accepts the number of months (mos), the litter size of each rabbit
// (lSz) and the reproduction age of rabbits (rAge) and recuursively calculates
// the number of total rabbits assuming the accepted parameters.
func calcFib(mos int, lS int, rAge int) int {
	if mos <= rAge {
		// rabbit is not of breeding age.
		return 1
	}

	return calcFib(mos-1, lS, rAge) + (lS * calcFib(mos-2, lS, rAge))
}

// fib accepts a file path to a `.txt` file containing two values (months and
// offspring produced by a pair of breeding rabbits) and returns the number of
// rabits in total given the input parameters.
func fib(fpath string) (int, error) {

	// read the entire file from the passed input filepath to string.
	ns, err := ioutil.ReadFile(fpath)
	if err != nil {
		return 0, fmt.Errorf("Cannot read values - %v", err)
	}
	inS := string(ns)
	inS = strings.Trim(inS, "\n")

	// parse the file contents into the target values; months (mo) and litter
	// size (lS)
	splitS := strings.Split(inS, " ")
	mo, _ := strconv.Atoi(splitS[0])
	lS, _ := strconv.Atoi(splitS[1])

	result := calcFib(mo, lS, 2)

	return result, nil
}
