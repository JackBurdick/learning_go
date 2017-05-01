package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const NUMSTEPS = 16

// each track can have multiple instruments
type instrument struct {
	instrumentName []byte
	instrumentID   uint8
	steps          []byte
}

// one track per slice file
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

// func printTrackFormat(curTrack track) {
// 	fmt.Println("%v\n", track.spliceHeader)
// 	fmt.Println("%v\n", track.versionString)
// }

func checkError(err error) {
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func createPrintString(curTrack track) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Saved with HW Version: %s\n", curTrack.versionString))
	buffer.WriteString(fmt.Sprintf("Tempo: %v\n", curTrack.tempo))
	for _, instrument := range curTrack.instruments {
		buffer.WriteString(fmt.Sprintf("(%v) %s\t", instrument.instrumentID, instrument.instrumentName))
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
	// fmt.Println(buffer.String())
	return buffer.String()
}

// func parseTrackToStruct() track {
//
// }

func main() {
	var tracks []track
	// get list of file names at target directory
	inDataDirectory := "fixtures"
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

	var fileLen int
	var spliceHeader [6]byte   // 6
	var trackSize int64        // 8
	var versionString [32]byte // 32
	var tempo float32          // 4

	// inspect data contents
	var id uint8
	var nameLength int32
	for _, fileName := range fileList {
		// open file
		fullPath := filepath.Join(inDataDirectory, fileName)
		fileContents, err := ioutil.ReadFile(fullPath)
		checkError(err)
		// PARSE
		//parseTrackToStruct(fileContents)
		newTrack := track{}
		//fmt.Printf("%s\n", hex.Dump(fileContents))
		buf := bytes.NewReader(fileContents)
		fileLen = len(fileContents)
		// NOTE: this will need to be looked at
		newTrack.trackSize = int64(fileLen)

		// Header: SPLICE
		err = binary.Read(buf, binary.BigEndian, &spliceHeader)
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
			newTrack.instruments = append(newTrack.instruments, curInstrument)
		}
		tracks = append(tracks, newTrack)

	}

	// print each track information per specification
	for _, track := range tracks {
		trackOutputFormatted := createPrintString(track)
		fmt.Println(trackOutputFormatted)
	}

}
