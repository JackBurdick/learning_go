package drum

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strings"
)

// config
// number of 'steps' played by an instrument each song
const numSteps = 16

// Instrument is a high level representation of a
// single instrument in the pattern
type Instrument struct {
	instrumentName []byte
	instrumentID   uint8
	steps          []byte
}

// Pattern is a high level representation
// of a track pattern contained within a .splice file
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
	// TODO: there should be safety checks each step of the way
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
func DecodeFile(path string) (*Pattern, error) {
	fileContents, err := ioutil.ReadFile(path)
	checkError(err)

	// decode
	p := &Pattern{}
	*p = parseTrackToStruct(fileContents)

	return p, nil
}
