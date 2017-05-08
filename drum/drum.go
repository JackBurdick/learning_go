package drum

import (
	"bytes"
	"fmt"
	"strings"
)

// config
// number of 'steps' played by an instrument each song
const numSteps = 16

// Instrument is a high level representation of a
// single instrument in the pattern
type Instrument struct {
	// instrumentName is a human readable instrument id
	instrumentName []byte

	// instrumentID is a unique id for instrument
	instrumentID uint8

	// steps is the number of parts the instrument may have in a given pattern
	steps []byte // 16 steps max
}

// Pattern is a high level representation
// of a track pattern contained within a .splice file
type Pattern struct {
	fileLen       int
	spliceHeader  [6]byte
	trackSize     int64
	versionString [32]byte
	tempo         float32
	instruments   []Instrument // multiple instruments could exist
}

func (curTrack *Pattern) String() string {
	// create string of specified track information
	// include: specific track information from struct
	// loop instruments in the track and print their steps

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
