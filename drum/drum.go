package drum

import (
	"bytes"
	"fmt"
	"strings"
)

// config
// number of 'steps' played by an instrument each song
const numSteps = 16

// Pattern is a high level representation
// of a track pattern contained within a .splice file
type Pattern struct {
	versionString [32]byte
	tempo         float32
	instruments   []Instrument // multiple instruments could exist
}

// Instrument is a high level representation of a
// single instrument in the pattern
type Instrument struct {
	instrumentName []byte
	instrumentID   uint8
	steps          [numSteps]bool // 16 steps max
}

// create string of specified track information
// include: specific track information from struct
// loop instruments in the track and print their steps
func (curTrack *Pattern) String() string {

	// write to buffer then return as buffer.String(), strings are immutable
	var buffer bytes.Buffer
	// track header;
	// Saved with HW Version: 0.909
	// Tempo: 240
	cleanedVersionString := fmt.Sprintf("%s", curTrack.versionString)
	cleanedVersionString = strings.Trim(cleanedVersionString, "\x00")
	curString := fmt.Sprintf("Saved with HW Version: %s\n", cleanedVersionString)
	//curString = strings.Trim(cleanedVersionString, "\n")

	buffer.WriteString(curString)
	//buffer.WriteString(fmt.Sprintf("Saved with HW Version: %s\n", curString))
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
			if step {
				buffer.WriteString("x")
			} else {
				buffer.WriteString("-")
			}
		}
		buffer.WriteString("|\n")
	}
	return buffer.String()
}
