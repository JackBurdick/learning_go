package drum

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
