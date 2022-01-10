package parallel_mapper

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type CityRankingPipeline struct {
	InputFileName      string
	NumberOfProcessors int
	waitGroup          sync.WaitGroup
	citiesChannel      chan *City
}

func NewCityRankingPipeline(inputFileName string, numberOfProcessors int) *CityRankingPipeline {
	return &CityRankingPipeline{
		InputFileName:      inputFileName,
		NumberOfProcessors: numberOfProcessors,
		citiesChannel:      make(chan *City),
	}
}

func (*CityRankingPipeline) Clean() {
	_ = os.Remove("outputSmall.txt")
	_ = os.Remove("outputMedium.txt")
	_ = os.Remove("outputBig.txt")
}

func (this *CityRankingPipeline) Run() {
	start := time.Now()

	this.waitGroup.Add(this.NumberOfProcessors + 1) // city reader
	writers := this.prepareWriters()
	this.startReading()
	this.startRanking(writers)
	this.waitGroup.Wait()

	for _, writer := range writers {
		writer.Flush()
	}

	fmt.Printf("Pipeline completed in %v\n", time.Since(start))
}

func (this *CityRankingPipeline) startReading() {
	go NewCityReader(this.InputFileName, this.citiesChannel, &this.waitGroup).Read()
}

func (this *CityRankingPipeline) prepareWriters() Writers {
	appender := &OsFileAppender{}
	writers := map[string]*CityWriter{
		smallCities: NewCityWriter(smallCities, appender),
		midCities:   NewCityWriter(midCities, appender),
		bigCities:   NewCityWriter(bigCities, appender),
	}
	return writers
}

func (this *CityRankingPipeline) startRanking(writers Writers) {
	for i := 1; i <= this.NumberOfProcessors; i++ {
		go this.addRanker(i, writers)
	}
}

func (this *CityRankingPipeline) addRanker(id int, writers Writers) {
	ranker := NewCityRanker(id, this.citiesChannel, &this.waitGroup, writers)
	ranker.Rank()
}
