package parallel_mapper

import (
	"sync"
	"testing"
)

func TestCityRanker_Rank_NoInputInChannel(t *testing.T) {
	t.Log("Given CityRanker with mocked writers.")
	{
		smallAppender, midAppender, bigAppender, writers := createTestFixtures()
		wg := sync.WaitGroup{}
		testChannel := make(chan *City)
		ranker := NewCityRanker(1, testChannel, &wg, writers)
		t.Log("When input channel has nothing to rank.")
		{
			wg.Add(1)
			go ranker.Rank()
			close(testChannel)
			wg.Wait()
		}
		t.Log("Then writers should've captured only headers.")
		{
			if smallAppender.captured != "=== small ===\n" {
				t.Errorf("But small appender has captured: %s", smallAppender.captured)
			}
			if midAppender.captured != "=== medium ===\n" {
				t.Errorf("But mid appender has captured: %s", midAppender.captured)
			}
			if bigAppender.captured != "=== big ===\n" {
				t.Errorf("But big appender has captured: %s", bigAppender.captured)
			}
		}
	}
}

func TestCityRanker_Rank_SomeCities(t *testing.T) {
	t.Log("Given CityRanker with mocked writers.")
	{
		smallAppender, midAppender, bigAppender, writers := createTestFixtures()
		wg := sync.WaitGroup{}
		testChannel := make(chan *City)
		ranker := NewCityRanker(1, testChannel, &wg, writers)
		t.Log("When input channel has some cities to rank.")
		{
			wg.Add(1)
			go ranker.Rank()
			testChannel <- NewCity("Brno (400,000)")
			testChannel <- NewCity("Zlin (75,000)")
			testChannel <- NewCity("Znojmo (34,000)")
			close(testChannel)
			wg.Wait()

			for _, writer := range writers {
				writer.Flush()
			}
		}
		t.Log("Then writers should've captured one city each.")
		{
			if smallAppender.captured != "=== small ===\nZnojmo (34,000)\n" {
				t.Errorf("But small appender has captured: %s", smallAppender.captured)
			}
			if midAppender.captured != "=== medium ===\nZlin (75,000)\n" {
				t.Errorf("But mid appender has captured: %s", midAppender.captured)
			}
			if bigAppender.captured != "=== big ===\nBrno (400,000)\n" {
				t.Errorf("But big appender has captured: %s", bigAppender.captured)
			}
		}
	}
}

func createTestFixtures() (*mockAppender, *mockAppender, *mockAppender, map[string]*CityWriter) {
	smallAppender := mockAppender{}
	midAppender := mockAppender{}
	bigAppender := mockAppender{}
	writers := map[string]*CityWriter{
		smallCities: NewCityWriter("small", &smallAppender),
		midCities:   NewCityWriter("medium", &midAppender),
		bigCities:   NewCityWriter("big", &bigAppender),
	}
	return &smallAppender, &midAppender, &bigAppender, writers
}
