package parallel_mapper

import (
	"strings"
	"sync"
	"testing"
)

func TestCityReader_Read_EmptyFile(t *testing.T) {
	t.Log("Given CityReader with an output channel.")
	{
		wg := sync.WaitGroup{}
		testChannel := make(chan *City)
		reader := NewCityReader("input_file_empty.txt", testChannel, &wg)
		t.Log("When reading an empty input file.")
		{
			wg.Add(2)
			go reader.Read()
			go func() {
				defer wg.Done()
				val, ok := <-testChannel
				if !ok {
					t.Log("Should close the test channel.")
				} else {
					t.Errorf("Should've closed the test channel, but received %s.", val)
				}
			}()
			wg.Wait()
		}
	}
}

func TestCityReader_Read_InputFile(t *testing.T) {
	t.Log("Given CityReader with an output channel.")
	{
		wg := sync.WaitGroup{}
		testChannel := make(chan *City)
		reader := NewCityReader("input_file.txt", testChannel, &wg)
		t.Log("When reading a populated input file.")
		{
			wg.Add(2)
			go reader.Read()
			go func() {
				defer wg.Done()
				cities := readCitiesFromChannel(testChannel)

				if len(cities) == 27 {
					t.Log("Should've read all valid cities from the input file, and skipped invalid entries.")
				} else {
					t.Errorf("Should've read all 27 cities from the input file, but read %d.", len(cities))
				}
				if containsAllOf(cities, "Prague", "Brno", "Plzeň") {
					t.Log("The read cities contain major important cities (data looks valid).")
				} else {
					t.Errorf("The read cities don't look valid, got this: %s.", cities)
				}
			}()
			wg.Wait()
		}
	}
}

func readCitiesFromChannel(testChannel chan *City) []*City {
	var cities []*City
	for {
		city, ok := <-testChannel
		if !ok {
			break
		}
		cities = append(cities, city)
	}
	return cities
}

func containsAllOf(cities []*City, candidates ...string) bool {
	for _, candidate := range candidates {
		found := false
		for _, city := range cities {
			if strings.Contains(city.name, candidate) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
