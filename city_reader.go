package parallel_mapper

import (
	"bufio"
	"log"
	"os"
	"sync"
)

type CityReader struct {
	InputFileName string
	output        chan *City
	wg            *sync.WaitGroup
}

func NewCityReader(inputFileName string, output chan *City, wg *sync.WaitGroup) *CityReader {
	return &CityReader{
		InputFileName: inputFileName,
		output:        output,
		wg:            wg,
	}
}

func (this *CityReader) Read() {
	defer this.wg.Done()

	inputFile := this.openInputFile()
	defer inputFile.Close()

	this.readCities(inputFile)
	close(this.output)
}

func (this *CityReader) openInputFile() *os.File {
	f, err := os.Open(this.InputFileName)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func (this *CityReader) readCities(inputFile *os.File) {
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		input := scanner.Text()
		if city := NewCity(input); city != nil {
			this.output <- city
		}
	}
}
