package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
// func DecodeFile(path string) (*Pattern, error) {
// 	p := &Pattern{}
// 	return p, nil
// }
//
// type Pattern struct{
//
// }
//

func main() {
	// get list of file names at target directory
	inDataDirectory := "fixtures"
	files, err := ioutil.ReadDir(inDataDirectory)
	if err != nil {
		fmt.Println("error: ", err)
	}

	// clean list names
	// > remove .DS_Store
	var fileList []string
	for _, file := range files {
		if file.Name() != ".DS_Store" {
			fileList = append(fileList, file.Name())
		}
	}

	var fileLen int
	// 6
	var spliceHeader [6]byte
	// 8
	var trackSize int64
	// 32
	var versionString [32]byte
	// 4
	var tempo float32

	// inspect data contents
	var id uint8
	var nameLength int32
	for _, fileName := range fileList {
		// open file
		fullPath := filepath.Join(inDataDirectory, fileName)
		fileContents, err := ioutil.ReadFile(fullPath)
		if err != nil {
			fmt.Println("ERROR")
		}
		//fmt.Printf("%s\n", hex.Dump(fileContents))
		buf := bytes.NewReader(fileContents)
		fileLen = len(fileContents)
		//fmt.Printf("Full length: %v\n", len(fileContents))

		// Header: SPLICE
		err = binary.Read(buf, binary.BigEndian, &spliceHeader)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("SPLICE: %s\n", spliceHeader)
		//fmt.Printf("JACK - %v\n", binary.Size(spliceHeader))
		fileLen -= binary.Size(spliceHeader)

		// Header: track size is big endian
		err = binary.Read(buf, binary.BigEndian, &trackSize)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("trackSize: %v\n", trackSize)
		fileLen -= binary.Size(trackSize)

		// Header: version
		err = binary.Read(buf, binary.BigEndian, &versionString)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("VERSION: %s\n", versionString)
		fileLen -= binary.Size(versionString)

		// Header: tempo
		// NOTE: tempo is little Endian?
		err = binary.Read(buf, binary.LittleEndian, &tempo)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("Tempo: %v\n", tempo)
		fileLen -= binary.Size(tempo)
		fmt.Printf("Cur Len -------> %v\n", fileLen)
		//fmt.Printf("len - %v", len(buf))

		// read file examples: https://gobyexample.com/reading-files

		// Read in body. id+name + 16 steps
		// while != EOF
		// TODO: There is an error with the end of file logic here..
		for fileLen > 0 {
			// ID
			//curIndex
			err = binary.Read(buf, binary.BigEndian, &id)
			fmt.Printf("id: %v\n", id)
			fileLen -= binary.Size(id)

			// Length of instrument name
			err = binary.Read(buf, binary.BigEndian, &nameLength)
			fmt.Printf("name length: %v\n", nameLength)
			fileLen -= binary.Size(nameLength)

			// name of instrument
			nameBuf := make([]byte, nameLength)
			err = binary.Read(buf, binary.LittleEndian, &nameBuf)
			fmt.Printf("name: %s\n", nameBuf)
			fileLen -= binary.Size(nameBuf)

			// steps
			// 16 is const
			stepBuf := make([]byte, 16)
			err = binary.Read(buf, binary.LittleEndian, &stepBuf)
			for _, num := range stepBuf {
				fmt.Printf("%v", num)
			}
			fmt.Printf("\nsteps: %v\n", stepBuf)
			fileLen -= binary.Size(stepBuf)
			fmt.Printf("JACK ==============: %v\n", fileLen)
		}

	}

	// printouts contain (header + data)
	// * Version
	// * Tempo
	// * tracks
	// 		* id, name, 16 steps
	//
	//

}
