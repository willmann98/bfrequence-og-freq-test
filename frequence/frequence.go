package main

import (
	"fmt"
	"sort"

	filehandler "github.com/Brofo/03/frequence/filehandler"
)

const fil = "C:/Users/Sindre/go/src/github.com/Brofo/is105-ica03/files/pg100.txt"

var lines int

func main() {

	LineCount(fil)
	fmt.Println("Linjer i tekstfilen: ", lines)
	MostFrequentRuneCounter(fil)
}

func MostFrequentRuneCounter(file string) {
	data, count := filehandler.MakeFileToByte(file)
	runeCount := make(map[byte]int)

	for i := 0; i < len(data[:count]); i++ {
		runeCounter := runeCount[data[i]]
		runeCount[data[i]] = runeCounter + 1

	}

	type kv struct {
		Key   byte
		Value int
	}

	var runeTeller []kv
	for k, v := range runeCount {
		runeTeller = append(runeTeller, kv{k, v})
	}

	sort.Slice(runeTeller, func(i, j int) bool {
		return runeTeller[i].Value > runeTeller[j].Value
	})

	for i := 0; i <= 4; i++ {
		fmt.Printf("%q, %d\n", runeTeller[i].Key, runeTeller[i].Value)
	}
}

func LineCount(file string) {
	data, count := filehandler.MakeFileToByte(file)

	for i := 0; i < len(data[:count]); i++ {
		byteValue := data[i]
		if byteValue == 13 {
			lines++
		}
	}
}
