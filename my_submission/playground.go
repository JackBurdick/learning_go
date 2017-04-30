package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
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

	var spliceHeader [6]byte
	var trackSize int64
	var versionString [32]byte
	var tempo float32
	// inspect data contents
	for _, fileName := range fileList {
		// open file
		fullPath := filepath.Join(inDataDirectory, fileName)
		fileContents, err := ioutil.ReadFile(fullPath)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("%s\n", hex.Dump(fileContents))
		buf := bytes.NewReader(fileContents)

		// Header: SPLICE
		err = binary.Read(buf, binary.BigEndian, &spliceHeader)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("%s\n", spliceHeader)

		// Header: track size is big endian
		err = binary.Read(buf, binary.BigEndian, &trackSize)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("%v\n", trackSize)

		// Header: version
		err = binary.Read(buf, binary.BigEndian, &versionString)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("%s\n", versionString)

		// Header: tempo
		// NOTE: tempo is little Endian?
		err = binary.Read(buf, binary.LittleEndian, &tempo)
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("%v\n", tempo)

		// read file examples: https://gobyexample.com/reading-files

		// Read in body. id+name + 16 steps

	}

	// printouts contain (header + data)
	// * Version
	// * Tempo
	// * tracks
	// 		* id, name, 16 steps
	//
	//

}
