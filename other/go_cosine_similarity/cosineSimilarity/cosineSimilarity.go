package cosineSimilarity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/jdkato/prose/tokenize"
)

// createSetFromJson
func createSetFromJSON(jsonPath string) map[string]bool {
	file, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var arr []string
	_ = json.Unmarshal([]byte(file), &arr)

	stopwordSet := make(map[string]bool)
	for _, sword := range arr {
		stopwordSet[sword] = true
	}

	return stopwordSet
}

// cosineSimilarity accepts a path to a directory containing .txt files and
// returns the cosine similarity between each document.
func cosineSimilarity(dPath string) (int, error) {

	fileMap := make(map[string][]string)

	// Create map of filename to tokenized content.
	dFiles, _ := ioutil.ReadDir(dPath)
	for _, f := range dFiles {
		b, err := ioutil.ReadFile(dPath + f.Name())
		if err != nil {
			fmt.Print(err)
		}

		// Convert bytes to string, then use 3rd party to tokenize.
		fileMap[f.Name()] = tokenize.TextToWords(string(b))
	}

	// Convert every word to lowercase in the document.
	for fName, doc := range fileMap {
		for i, word := range doc {
			doc[i] = strings.ToLower(word)
		}
		fileMap[fName] = doc
	}

	// Filter the stop words out of the content.
	swSet := createSetFromJSON("../fixtures/en_stopwords.json")
	for fName, doc := range fileMap {
		filteredDoc := doc[:0]
		for _, word := range doc {
			_, ok := swSet[word]
			if !ok {
				filteredDoc = append(filteredDoc, word)
			}
		}
		fileMap[fName] = filteredDoc
	}

	// Filter the punctuation from the content.
	puncSet := createSetFromJSON("../fixtures/en_punctuation.json")
	for fName, doc := range fileMap {
		filteredDoc := doc[:0]
		for _, word := range doc {
			_, ok := puncSet[word]
			if !ok {
				filteredDoc = append(filteredDoc, word)
			}
		}
		fileMap[fName] = filteredDoc
	}

	// Calculate term frequency and word to doc count.
	fileToTF := make(map[string]map[string]int)
	allWordSet := make(map[string]bool)
	for fileName, doc := range fileMap {
		wordSet := make(map[string]int)
		for _, word := range doc {
			wordSet[word]++
			allWordSet[word] = true
		}
		fileToTF[fileName] = wordSet
	}

	// Calculate number of docs word appears on. (word to doc count)
	// This is also used as a total word set as it includes every word used
	// across all documents.
	wordToDC := make(map[string]int)
	for _, docSet := range fileToTF {
		for word := range docSet {
			wordToDC[word]++
		}
	}

	// Calculate normalized term frequency.
	// Calculate number of documents a word occurs in
	fileToNTF := make(map[string]map[string]float64)
	wordToDF := make(map[string]int)
	for fileName, wordSet := range fileToTF {
		wordToNTF := make(map[string]float64)
		numWords := float64(len(wordSet))
		for word, tf := range wordSet {
			wordToNTF[word] = float64(tf) / numWords
			wordToDF[word]++
		}
		fileToNTF[fileName] = wordToNTF
	}

	// -------------- calculate tf-idf
	// IDF(word) = 1 + loge(Total Number Of Documents / Number Of Documents w/ word  in it)
	// normalized TF * inverse document frequency
	numDocs := len(fileToNTF)
	//wordToIDF := make(map[string]float64)
	fileNameTFIDF := make(map[string]map[string]float64)
	for fileName, wordToNTF := range fileToNTF {
		wordToTFIDF := make(map[string]float64)
		for word, ntf := range wordToNTF {
			docCount := wordToDC[word]
			inner := float64(numDocs) / float64(docCount)
			idf := 1 + math.Log(inner)
			tfidf := ntf * idf
			wordToTFIDF[word] = tfidf
		}
		fileNameTFIDF[fileName] = wordToTFIDF
	}

	//fmt.Println(fileNameTFIDF)

	// -------------- calculate cosine similarity

	// create map of word to tf-idf in each document.
	fileToTFIDFSet := make(map[string]map[string]float64)
	for fName := range fileMap {
		finalWordToTFIDF := make(map[string]float64)
		for word := range allWordSet {
			val, ok := fileNameTFIDF[fName][word]
			if !ok {
				val = 0.0
			}
			finalWordToTFIDF[word] = val
		}
		fileToTFIDFSet[fName] = finalWordToTFIDF
	}

	// Calculate tfidf vector for each document.
	fNameToTFIDFVector := make(map[string]map[string]float64)
	for fName := range fileMap {
		docTFIDFVector := make(map[string]float64)
		for word := range allWordSet {
			val, ok := fileToTFIDFSet[fName][word]
			if !ok {
				val = 0.0
			}
			docTFIDFVector[word] = val
		}
		fNameToTFIDFVector[fName] = docTFIDFVector
	}

	// ------------ Calculate cosine similarity between each document.
	// ----------- numerator
	// sum of the product of each corresponding tfidf value
	fNameToCosineNumMap := make(map[string]map[string]float64)
	for fNameA, docTFIDFVectorA := range fNameToTFIDFVector {
		tempMap := make(map[string]float64)
		for fNameB, docTFIDFVectorB := range fNameToTFIDFVector {
			sumProdAB := 0.0
			for word, valA := range docTFIDFVectorA {
				valB := docTFIDFVectorB[word]
				prodAB := valA * valB
				sumProdAB += prodAB
			}
			tempMap[fNameB] = sumProdAB
		}
		fNameToCosineNumMap[fNameA] = tempMap
	}

	//fmt.Println(fNameToCosineNumMap)

	// ---------- denom
	// square root of each value
	fNameToCosDenPre := make(map[string]float64)
	for fName, docTFIDFVector := range fNameToTFIDFVector {
		var numPre float64
		for _, val := range docTFIDFVector {
			v := math.Pow(val, 2)
			numPre += v
		}
		numPre = math.Sqrt(numPre)
		fNameToCosDenPre[fName] = numPre
	}

	// TODO: this needs to be optimized so that we don't calculate values twice
	// a cross functionality
	fNameToCosDen := make(map[string]map[string]float64)
	for fNameA, valA := range fNameToCosDenPre {
		tempMap := make(map[string]float64)
		for fNameB, valB := range fNameToCosDenPre {
			tempMap[fNameB] = valA * valB
		}
		fNameToCosDen[fNameA] = tempMap
	}

	fNameToCosSim := make(map[string]map[string]float64)
	for fNameA, numMapA := range fNameToCosineNumMap {
		tempCosMap := make(map[string]float64)
		for fNameB, num := range numMapA {
			denom := fNameToCosDen[fNameA][fNameB]
			val := num / denom
			tempCosMap[fNameB] = val
		}
		fNameToCosSim[fNameA] = tempCosMap
	}

	for file, vMap := range fNameToCosSim {
		fmt.Printf("%v| ", file)
		for fB, v := range vMap {
			fmt.Printf(" %v:%.2f |", fB, v)
		}
		fmt.Printf("\n")
	}

	//fmt.Println(fNameToCosSim)

	return 1, nil
}
