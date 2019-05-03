package filehandler

import (
	"log"
	"os"
)

func MakeFileToByte(textfile string) (data []byte, count int) {
	file, err := os.Open(textfile)
	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sizeOfSlice := finfo.Size()

	data = make([]byte, sizeOfSlice)
	count, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	return data, count
}
