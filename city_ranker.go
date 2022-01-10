package parallel_mapper

import (
	"math/rand"
	"sync"
	"time"
)

const smallCities = "Small"
const midCities = "Medium"
const bigCities = "Big"

type Writers map[string]*CityWriter

type CityRanker struct {
	id      int
	input   chan *City
	wg      *sync.WaitGroup
	writers Writers
}

func NewCityRanker(id int, input chan *City, wg *sync.WaitGroup, writers Writers) *CityRanker {
	return &CityRanker{id: id, input: input, wg: wg, writers: writers}
}

func (this *CityRanker) Rank() {
	defer this.wg.Done()

	for {
		city, ok := <-this.input
		if !ok {
			// work finished, nothing else to rank
			return
		}
		writer := this.rankCity(city)
		if writer != nil {
			writer.Write(city)
		}
	}
}

func (this *CityRanker) rankCity(city *City) *CityWriter {
	population := city.population

	// imagine this is some expensive HTTP Call
	// that's why it makes sense to parallelize rankers
	timeToWait := rand.Intn(50-20) + 20
	time.Sleep(time.Duration(timeToWait) * time.Millisecond)

	switch {
	case population < 50000:
		return this.writers[smallCities]
	case population < 100000:
		return this.writers[midCities]
	default:
		return this.writers[bigCities]
	}
}
