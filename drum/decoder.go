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

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := io.Reader(file)

	// decode
	p := &Pattern{}
	*p, err = parseSpliceToPattern(r)
	if err != nil {
		fmt.Println("Error in parseSpliceToPattern: ", err)
	}

	return p, nil
}

// parseSpliceToPattern decodes the given `.splice` files and stores
// the relevant information in the Pattern struct.
func parseSpliceToPattern(r io.Reader) (Pattern, error) {

	newTrack := &Pattern{}

	// Decode SPLICE from header to ensure correct file type
	var spliceHeader [len(expectedHeader)]byte // 6
	err := binary.Read(r, binary.BigEndian, &spliceHeader)
	if err != nil {
		return *newTrack, errors.New("unable to decode splice header: " + err.Error())
	}

	// ensure the decoded header matches the expected file type
	// if the expected header is not present, abort decoding current file
	if string(spliceHeader[:len(expectedHeader)]) != expectedHeader {
		return *newTrack, errors.New("decoded SPLICE header does not match expected value: " + err.Error())
	}

	// Decode trackSize from pattern Header
	var trackSize int64 // 8 bytes
	err = binary.Read(r, binary.BigEndian, &trackSize)
	if err != nil {
		return *newTrack, errors.New("unable to decode trackSize: " + err.Error())
	}

	// Decode versionString and store in struct
	var versionString [32]byte
	err = binary.Read(r, binary.BigEndian, &versionString)
	if err != nil {
		return *newTrack, errors.New("unable to decode versionString: " + err.Error())
	}
	newTrack.versionString = versionString

	// Decode tempo and store in struct
	// NOTE: tempo is little Endian
	var tempo float32 // 4
	err = binary.Read(r, binary.LittleEndian, &tempo)
	if err != nil {
		return *newTrack, errors.New("unable to decode tempo: " + err.Error())
	}
	newTrack.tempo = tempo

	// LimitedReader needs to take `-36` = versionString(32) + tempo(4)
	// into account since we've already read in versionString and tempo
	lr := io.LimitReader(r, trackSize-36)
	for {
		// decode instrument information from the remaining information
		done, err := readInstrumentsFromTrack(lr, newTrack)
		if done {
			break
		} else if err != nil {
			return *newTrack, err
		}
	}

	return *newTrack, nil
}

// readInstrumentsFromTrack decodes the instrument information contained within
// the body of the Pattern and appends the newInstrument to the Pattern.instruments a
// true (bool) is returned when all the instruments have been read.
func readInstrumentsFromTrack(lr io.Reader, newTrack *Pattern) (bool, error) {
	newInstrument := Instrument{}

	// Decode instrumentID (unique id) of instrument and store in Instrument struct
	var instrumentID uint8
	err := binary.Read(lr, binary.BigEndian, &instrumentID)
	if err == io.EOF {
		// all the information has been read
		return true, nil
	} else if err != nil {
		return false, errors.New("unable to decode instrumentID: " + err.Error())
	}
	newInstrument.instrumentID = instrumentID

	// Decode namLen (byte size of the human readable name) of instrument name
	var nameLen int32
	err = binary.Read(lr, binary.BigEndian, &nameLen)
	if err != nil {
		return false, errors.New("unable to decode instrument nameLen: " + err.Error())
	}

	// Decode human readable name of instrument and store in Instrument struct
	nameBuf := make([]byte, nameLen)
	err = binary.Read(lr, binary.LittleEndian, &nameBuf)
	if err != nil {
		return false, errors.New("unable to decode instrument nameBuf: " + err.Error())
	}
	newInstrument.instrumentName = nameBuf

	// steps were stored on HW as bytes but,
	// can be stored as bools instead since only concern is a binary state
	var stepArr [numSteps]byte
	err = binary.Read(lr, binary.LittleEndian, &stepArr)
	if err != nil {
		return false, errors.New("unable to decode instrument steps: " + err.Error())
	}

	// convert from bytes to bool and store in Instrument struct
	for i := range stepArr {
		if stepArr[i] == 0x01 {
			newInstrument.steps[i] = true
		} else if stepArr[i] == 0x00 {
			newInstrument.steps[i] = false
		} else {
			return false, errors.New("unexpected values in instrument steps: " + err.Error())
		}
	}

	// all target information was decoded as expected
	// add the newly decoded instrument information to our instruments slice
	newTrack.instruments = append(newTrack.instruments, newInstrument)
	return false, nil
}
