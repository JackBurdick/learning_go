package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// config
// number of 'steps' played by an instrument each song
const NUMSTEPS = 16

// each track can have multiple instruments
type instrument struct {
	instrumentName []byte
	instrumentID   uint8
	steps          []byte
}

// one track per `.splice` file
type track struct {
	fileLen       int
	spliceHeader  [6]byte  // 6
	trackSize     int64    // 8
	versionString [32]byte // 32
	tempo         float32  // 4
	instruments   []instrument
}

// RESOURCES:
// read file examples: https://gobyexample.com/reading-files
// range: https://github.com/golang/go/wiki/Range
// slice in struct: https://stackoverflow.com/questions/18042439/go-append-to-slice-in-struct
// const: https://blog.golang.org/constants
// binary reader: https://golang.org/pkg/encoding/binary/#Read
// hex dump for visualization/debug: https://golang.org/pkg/encoding/hex/#Dump
// create strings: https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go

func checkError(err error) {
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func createPrintString(curTrack track) string {
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

func parseTrackToStruct(fileContents []byte) track {
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

	newTrack := track{}

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
		curInstrument := instrument{}
		// ID
		err = binary.Read(buf, binary.BigEndian, &id)
		checkError(err)
		fileLen -= binary.Size(id)
		curInstrument.instrumentID = id

		// Length of instrument name
		err = binary.Read(buf, binary.BigEndian, &nameLength)
		checkError(err)
		fileLen -= binary.Size(nameLength)

		// name of instrument
		nameBuf := make([]byte, nameLength)
		err = binary.Read(buf, binary.LittleEndian, &nameBuf)
		checkError(err)
		fileLen -= binary.Size(nameBuf)
		curInstrument.instrumentName = nameBuf

		// steps
		stepBuf := make([]byte, NUMSTEPS)
		err = binary.Read(buf, binary.LittleEndian, &stepBuf)
		checkError(err)
		fileLen -= binary.Size(stepBuf)
		curInstrument.steps = stepBuf
		// add instrument to instruments on track
		newTrack.instruments = append(newTrack.instruments, curInstrument)
	}
	return newTrack
}

func main() {
	var tracks []track
	// config: root input dir
	inDataDirectory := "fixtures"

	// get list of file names at target directory
	files, err := ioutil.ReadDir(inDataDirectory)
	checkError(err)

	// clean list names
	// - remove .DS_Store
	var fileList []string
	for _, file := range files {
		if file.Name() != ".DS_Store" {
			fileList = append(fileList, file.Name())
		}
	}

	// loop directory and store parsed contents
	for _, fileName := range fileList {
		// open, read in file
		fullPath := filepath.Join(inDataDirectory, fileName)
		fileContents, err := ioutil.ReadFile(fullPath)
		checkError(err)

		// parse
		newTrack := parseTrackToStruct(fileContents)

		// store track information
		tracks = append(tracks, newTrack)
	}

	// print track information per specification
	for _, track := range tracks {
		trackOutputFormatted := createPrintString(track)
		fmt.Println(trackOutputFormatted)
	}

}
