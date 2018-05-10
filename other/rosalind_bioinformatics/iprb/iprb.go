package iprb

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// iprb accepts a file path to a `.txt` file containing three values
// representing a population containing k homozygous dominant, m heterozygous,
// and n homozygous recessive individuals and returns the probability that two
// randomly selected mating organisms will produce an individual possessing a
// dominant allele. Note: assume any two organisms can mate
func iprb(fpath string) (float64, error) {

	// read the entire file from the passed input filepath to string.
	inV, err := ioutil.ReadFile(fpath)
	if err != nil {
		return 0, fmt.Errorf("Cannot read values - %v", err)
	}
	inS := string(inV)
	inS = strings.Trim(inS, "\n")

	// parse the file contents into the target values; months (mo) and litter
	// size (lS)
	splitS := strings.Split(inS, " ")
	RR, _ := strconv.Atoi(splitS[0])
	Rr, _ := strconv.Atoi(splitS[1])
	rr, _ := strconv.Atoi(splitS[2])

	// calculate probability of each allele.
	// NOTE: RR is not considered since any mating pair containing one RR would
	// results in a 100% chance of the individual containing a dominant allele
	t := RR + Rr + rr
	pRr := float64(Rr) / float64(t)
	prr := float64(rr) / float64(t)

	// subtract variations from intial probability of 100%.
	prb := 1.0
	// probability of both parents being homozygous recessive.
	prb -= prr * ((float64(rr) - 1) / (float64(t) - 1))
	// probability of homozygous recessive mating with heterozygous, x0.5 for
	// the recessive trait, x2 for both cases.
	prb -= prr * float64(2) * (float64(Rr) / (float64(t) - float64(1)) * float64(0.50))
	// probability of the heterozygous mating with heterozygous, x(0.5*0.5).
	prb -= pRr * (float64(Rr-1) / float64(t-1)) * 0.25

	return prb, nil
}
