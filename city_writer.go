package parallel_mapper

import (
	"fmt"
	"os"
	"sort"
	"sync"
)

type FileAppender interface {
	Append(location string, content string)
}

type CityWriter struct {
	outputFileName string
	appender       FileAppender
	buffer         []*City
}

func NewCityWriter(name string, appender FileAppender) *CityWriter {
	outputFileName := fmt.Sprintf("output%s.txt", name)

	writer := &CityWriter{
		outputFileName: outputFileName,
		appender:       appender,
	}
	appender.Append(outputFileName, fmt.Sprintf("=== %s ===\n", name))
	return writer
}

var mutex = sync.Mutex{}

func (this *CityWriter) Write(city *City) {
	mutex.Lock()
	defer mutex.Unlock()
	{
		this.buffer = append(this.buffer, city)
	}
}

func (this *CityWriter) Flush() {
	sort.Slice(this.buffer, func(i, j int) bool {
		return this.buffer[i].population > this.buffer[j].population
	})

	for _, city := range this.buffer {
		this.appender.Append(this.outputFileName, city.String()+"\n")
	}
}

type OsFileAppender struct{}

func (this *OsFileAppender) Append(location string, content string) {
	file, _ := os.OpenFile(location, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	_, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
