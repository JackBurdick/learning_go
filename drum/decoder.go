package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

// config
// number of 'steps' played by an instrument each song
const numSteps = 16

// each track can have multiple `Instrument`s
type Instrument struct {
	instrumentName []byte
	instrumentID   uint8
	steps          []byte
}

// one `Pattern` per `.splice` file
type Pattern struct {
	fileLen       int
	spliceHeader  [6]byte  // 6
	trackSize     int64    // 8
	versionString [32]byte // 32
	tempo         float32  // 4
	instruments   []Instrument
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func parseTrackToStruct(fileContents []byte) Pattern {
	// parse the given `.splice` files and store
	// relevant information in the struct
	// 1. read in file
	// 2. get file length
	// 3. parse and store relevant parts, subtract size from file length
	//NOTE: Use (for debuging): fmt.Printf("%s\n", hex.Dump(fileContents))

	// track temp vars
	var fileLen int
	var spliceHeader [6]byte   // 6
	var trackSize int64        // 8
	var versionString [32]byte // 32
	var tempo float32          // 4

	// instrument temp vars
	var id uint8
	var nameLength int32

	newTrack := &Pattern{}

	buf := bytes.NewReader(fileContents)
	fileLen = len(fileContents)
	newTrack.trackSize = int64(fileLen)

	// Header: SPLICE
	err := binary.Read(buf, binary.BigEndian, &spliceHeader)
	checkError(err)
	fileLen -= binary.Size(spliceHeader)
	newTrack.spliceHeader = spliceHeader

	// Header: track size is big endian
	err = binary.Read(buf, binary.BigEndian, &trackSize)
	checkError(err)
	fileLen -= binary.Size(trackSize)
	newTrack.trackSize = trackSize

	// Header: version
	err = binary.Read(buf, binary.BigEndian, &versionString)
	checkError(err)
	fileLen -= binary.Size(versionString)
	//versionStringTrimmed := strings.TrimRight(versionString, '\x00')
	newTrack.versionString = versionString

	// Header: tempo
	// NOTE: tempo is little Endian?
	err = binary.Read(buf, binary.LittleEndian, &tempo)
	checkError(err)
	fileLen -= binary.Size(tempo)
	newTrack.tempo = tempo

	// Read in body. id+name + 16 steps
	// TODO: Issue is with pattern 5...
	for fileLen > 0 {
		curInstrument := Instrument{}
		// ID
		err = binary.Read(buf, binary.BigEndian, &id)
		checkError(err)
		fileLen -= binary.Size(id)
		curInstrument.instrumentID = id

		// Length of instrument name
		err = binary.Read(buf, binary.BigEndian, &nameLength)
		checkError(err)
		if nameLength > 10 {
			// TODO: this is a cheap fix to a larger problem
			break
		}
		fileLen -= binary.Size(nameLength)

		// name of instrument
		nameBuf := make([]byte, nameLength)
		err = binary.Read(buf, binary.LittleEndian, &nameBuf)
		checkError(err)
		fileLen -= binary.Size(nameBuf)
		curInstrument.instrumentName = nameBuf

		// steps
		stepBuf := make([]byte, numSteps)
		err = binary.Read(buf, binary.LittleEndian, &stepBuf)
		checkError(err)
		fileLen -= binary.Size(stepBuf)
		curInstrument.steps = stepBuf
		// add instrument to instruments on track
		newTrack.instruments = append(newTrack.instruments, curInstrument)
	}
	return *newTrack
}

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {
	fileContents, err := ioutil.ReadFile(path)
	checkError(err)

	// decode
	p := &Pattern{}
	*p = parseTrackToStruct(fileContents)

	//trackOutputFormatted := createPrintString(*p)

	return p, nil
}
