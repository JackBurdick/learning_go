package drum

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
func DecodeFile(path string) (*Pattern, error) {
	// fileContents, err := ioutil.ReadFile(path)
	file, err := os.Open(path)
	checkError(err)
	defer file.Close()

	r := io.Reader(file)

	// decode
	p := &Pattern{}
	*p, err = parseSpliceToPattern(r)
	checkError(err)

	return p, nil
}

// parse the given `.splice` files and store
// relevant information in the struct
// 1. read in file
// 2. get file length
// 3. parse and store relevant parts, subtract size from file length
func parseSpliceToPattern(r io.Reader) (Pattern, error) {

	newTrack := &Pattern{}

	// Header: SPLICE
	var spliceHeader [6]byte // 6
	err := binary.Read(r, binary.BigEndian, &spliceHeader)
	checkError(err)

	// Header: track size is big endian
	var trackSize int64 // 8
	err = binary.Read(r, binary.BigEndian, &trackSize)
	checkError(err)

	// Header: version
	var versionString [32]byte // 32
	err = binary.Read(r, binary.BigEndian, &versionString)
	checkError(err)
	newTrack.versionString = versionString

	// Header: tempo
	// NOTE: tempo is little Endian
	var tempo float32 // 4
	err = binary.Read(r, binary.LittleEndian, &tempo)
	checkError(err)
	newTrack.tempo = tempo

	// Read in body containing instrument information
	lr := io.LimitReader(r, trackSize-36)
	for {
		done, err := readInstrumentsFromTrack(lr, newTrack)
		if done {
			break
		} else if err != nil {
			return *newTrack, err
		}

	}

	return *newTrack, nil
}

func readInstrumentsFromTrack(lr io.Reader, newTrack *Pattern) (bool, error) {
	curInstrument := Instrument{}

	err := binary.Read(lr, binary.BigEndian, &curInstrument.instrumentID)
	if err == io.EOF {
		// we've read all the information
		return true, nil
	} else if err != nil {
		return false, errors.New("unable to decode instrumentID: " + err.Error())
	}

	// Length of instrument name
	var nameLen int32
	err = binary.Read(lr, binary.BigEndian, &nameLen)
	if err != nil {
		return false, errors.New("unable to decode instrument nameLen: " + err.Error())
	}
	//checkError(err)
	if nameLen > 12 {
		// TODO: this is a cheap fix to a larger problem
		fmt.Println("hi")
	}

	// read human readable name of instrument
	nameBuf := make([]byte, nameLen)
	err = binary.Read(lr, binary.LittleEndian, &nameBuf)
	checkError(err)
	curInstrument.instrumentName = nameBuf

	// steps were stored on HW as bytes
	// but we can store them as bools instead
	var stepBuf [numSteps]byte
	err = binary.Read(lr, binary.LittleEndian, &stepBuf)
	checkError(err)

	for i := range stepBuf {
		if stepBuf[i] == 0x0001 {
			curInstrument.steps[i] = true
		} else {
			curInstrument.steps[i] = false
		}
	}
	// add instrument to instruments on track
	newTrack.instruments = append(newTrack.instruments, curInstrument)

	return false, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error: ", err)
	}
}
