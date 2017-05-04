// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"bytes"
	"fmt"
)

func (curTrack *Pattern) String() string {
	// create string of specified track information
	// include: specific track information from struct
	// loop instruments in the track and print their steps

	// write to buffer then return as buffer.String(), strings are immutable
	var buffer bytes.Buffer
	// track header;
	// Saved with HW Version: 0.909
	// Tempo: 240
	buffer.WriteString(fmt.Sprintf("Saved with HW Version: %s\n", curTrack.versionString))
	buffer.WriteString(fmt.Sprintf("Tempo: %v\n", curTrack.tempo))

	// print instrument/step info > (99) Maracas	|x-x-|x-x-|x-x-|x-x-|
	for _, instrument := range curTrack.instruments {
		// identification > (0) SubKick
		buffer.WriteString(fmt.Sprintf("(%v) %s\t", instrument.instrumentID, instrument.instrumentName))
		// steps > |x---|----|x---|----|
		for i, step := range instrument.steps {
			if i%4 == 0 {
				buffer.WriteString("|")
			}
			// per spec. exception: print "E" if unknown
			if step == 1 {
				buffer.WriteString("x")
			} else if step == 0 {
				buffer.WriteString("-")
			} else {
				buffer.WriteString("E")
			}
		}
		buffer.WriteString("|\n")
	}
	return buffer.String()
}
