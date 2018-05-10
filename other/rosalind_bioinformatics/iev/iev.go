package iev

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func iev(fpath string) (float64, error) {
	// read the entire file from the passed input filepath to string.
	vals, err := ioutil.ReadFile(fpath)
	if err != nil {
		return 0.0, fmt.Errorf("Cannot read values - %v", err)
	}
	valS := string(vals)
	valS = strings.Trim(valS, "\n")

	// parse the file contents and assign to the target values.
	vSplit := strings.Split(valS, " ")
	AAAAi, _ := strconv.Atoi(vSplit[0])
	AAAA := float64(AAAAi) * 2 * 1
	AAAai, _ := strconv.Atoi(vSplit[1])
	AAAa := float64(AAAai) * 2 * 1
	AAaai, _ := strconv.Atoi(vSplit[2])
	AAaa := float64(AAaai) * 2 * 1
	AaAai, _ := strconv.Atoi(vSplit[3])
	AaAa := float64(AaAai) * 2 * 0.75
	Aaaai, _ := strconv.Atoi(vSplit[4])
	Aaaa := float64(Aaaai) * 2 * 0.5
	aaaai, _ := strconv.Atoi(vSplit[5])
	aaaa := float64(aaaai) * 0

	// calculate expected.
	expected := float64(0.0)
	expected += AAAA
	expected += AAAa
	expected += AAaa
	expected += AaAa
	expected += Aaaa
	expected += aaaa

	return expected, nil
}
