package gc

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type result struct {
	ID  string
	gcC float64
}

// TODO: this is a messy implementation.. Currently passing, but refactor needed.

func gc(fPath string) (result, error) {

	// Read file.
	fContent, err := ioutil.ReadFile(fPath)
	if err != nil {
		var empty result
		return empty, fmt.Errorf("Cannot read values - %v", err)
	}
	fSt := string(fContent)
	fSt = strings.Trim(fSt, "\n")

	// Split on start of fasta sequence, remove first, blank, entry.
	split := strings.Split(fSt, ">")
	split = append(split[:0], split[0+1:]...)

	var contentS []result
	for _, entry := range split {
		var content result

		entryS := strings.SplitN(entry, "\n", 2)
		id := entryS[0]
		id = strings.Trim(id, "\n")
		id = strings.Trim(id, "\r")
		seq := entryS[1]

		gcC := 0
		tC := 0
		for _, c := range seq {
			ch := string(c)
			ch = strings.ToUpper(ch)

			// Count gc content.
			if ch == "C" || ch == "G" {
				gcC++
			}

			// Count target nucleotides.
			if ch == "A" || ch == "C" || ch == "G" || ch == "T" {
				tC++
			}

		}
		gcP := (float64(gcC) / float64(tC)) * 100.0
		content.ID = id
		content.gcC = gcP

		contentS = append(contentS, content)
	}

	// Loop results to find highest GC content. This could be included in the
	// the original loop if performance was important
	var largest result
	lP := 0.0
	var lID int
	for i, v := range contentS {
		if v.gcC > lP {
			lP = v.gcC
			lID = i
		}
	}
	largest = contentS[lID]

	return largest, nil
}
