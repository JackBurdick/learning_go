package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const NUMSTEPS = 16

// RESOURCES:
// read file examples: https://gobyexample.com/reading-files

func checkError(err error) {
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func main() {
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
		//fmt.Printf("%s\n", hex.Dump(fileContents))
		buf := bytes.NewReader(fileContents)
		fileLen = len(fileContents)

		// Header: SPLICE
		err = binary.Read(buf, binary.BigEndian, &spliceHeader)
		checkError(err)
		fileLen -= binary.Size(spliceHeader)

		// Header: track size is big endian
		err = binary.Read(buf, binary.BigEndian, &trackSize)
		checkError(err)
		fileLen -= binary.Size(trackSize)

		// Header: version
		err = binary.Read(buf, binary.BigEndian, &versionString)
		checkError(err)
		fileLen -= binary.Size(versionString)

		// Header: tempo
		// NOTE: tempo is little Endian?
		err = binary.Read(buf, binary.LittleEndian, &tempo)
		checkError(err)
		fileLen -= binary.Size(tempo)

		// Read in body. id+name + 16 steps
		// TODO: Issue is with pattern 5...
		for fileLen > 0 {
			// ID
			err = binary.Read(buf, binary.BigEndian, &id)
			checkError(err)
			fileLen -= binary.Size(id)

			// Length of instrument name
			err = binary.Read(buf, binary.BigEndian, &nameLength)
			checkError(err)
			fileLen -= binary.Size(nameLength)

			// name of instrument
			nameBuf := make([]byte, nameLength)
			err = binary.Read(buf, binary.LittleEndian, &nameBuf)
			checkError(err)
			fileLen -= binary.Size(nameBuf)

			// steps
			stepBuf := make([]byte, NUMSTEPS)
			err = binary.Read(buf, binary.LittleEndian, &stepBuf)
			checkError(err)
			fileLen -= binary.Size(stepBuf)

		}

	}

}
