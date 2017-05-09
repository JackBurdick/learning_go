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
	// instrumentName is a human readable name
	instrumentName []byte
	// instrumentID is a unique instrument id
	instrumentID uint8
	steps        [numSteps]bool
}

// create string of track information according to spec
// e.g.;
//	Saved with HW Version: 0.708-alpha
//	Tempo: 999
//	(1) Kick	|x---|----|x---|----|
//	(2) HiHat	|x-x-|x-x-|x-x-|x-x-|
func (curTrack *Pattern) String() string {
	// write formatted information to buffer then return as string

	var buffer bytes.Buffer

	// format and write cleaned track header to buffer
	cleanedVersionString := fmt.Sprintf("%s", curTrack.versionString)
	cleanedVersionString = strings.Trim(cleanedVersionString, "\x00")
	curString := fmt.Sprintf("Saved with HW Version: %s\n", cleanedVersionString)
	buffer.WriteString(curString)

	// format and write tempo information to buffer
	buffer.WriteString(fmt.Sprintf("Tempo: %v\n", curTrack.tempo))

	// write all instruments to buffer
	for _, instrument := range curTrack.instruments {
		// format and write intrument id and humanreadable name to buffer
		buffer.WriteString(fmt.Sprintf("(%v) %s\t", instrument.instrumentID, instrument.instrumentName))

		for i, step := range instrument.steps {
			// format and write intrument step information to buffer
			if i%4 == 0 {
				// write pipe separator every 4 steps
				buffer.WriteString("|")
			}

			if step {
				buffer.WriteString("x")
			} else {
				buffer.WriteString("-")
			}
		}
		buffer.WriteString("|\n")
	}

	// convert buffer to string before returning
	return buffer.String()
}
