package parallel_mapper

import (
	"os"
	"testing"
)

func TestCityRankingPipeline_Clean(t *testing.T) {
	t.Log("Given previous output files exist.")
	{
		if err := os.WriteFile("outputSmall.txt", []byte{1}, 0600); err != nil {
			t.Error("Couldn't create outputSmall.txt for tests")
		}
		if err := os.WriteFile("outputMedium.txt", []byte{1}, 0600); err != nil {
			t.Error("Couldn't create outputMedium.txt for tests")
		}
		if err := os.WriteFile("outputBig.txt", []byte{1}, 0600); err != nil {
			t.Error("Couldn't create outputBig.txt for tests")
		}

		t.Log("When Clean step is called.")
		{
			pipeline := NewCityRankingPipeline("input_file.txt", 1)
			pipeline.Clean()

			t.Log("Then all previous out files should be gone.")
			{
				if _, err := os.ReadFile("outputSmall.txt"); err == nil {
					t.Error("but outputSmall.txt should've been removed")
				}
				if _, err := os.ReadFile("outputMedium.txt"); err == nil {
					t.Error("but outputMedium.txt should've been removed")
				}
				if _, err := os.ReadFile("outputBig.txt"); err == nil {
					t.Error("but outputBig.txt should've been removed")
				}
			}
		}
	}
}

func TestCityRankingPipeline_Run(t *testing.T) {
	var tests = []struct {
		pipeline *CityRankingPipeline
		workers  int
	}{
		{
			pipeline: NewCityRankingPipeline("input_file.txt", 1),
			workers:  1,
		},
		{
			pipeline: NewCityRankingPipeline("input_file.txt", 4),
			workers:  4,
		},
	}
	for _, test := range tests {
		t.Logf("Given a new pipeline with %d ranking worker(s) and a valid input file.", test.workers)
		{
			t.Log("When the pipeline is run.")
			{
				test.pipeline.Clean()
				test.pipeline.Run()

				t.Log("Then the cities should be ranked and saved into the output files.")
				{
					content, err := os.ReadFile("outputSmall.txt")
					if err != nil {
						t.Error("but outputSmall.txt should've been created")
					}
					if string(content) != smallContent {
						t.Errorf("but outputSmall.txt contained %s instead", string(content))
					}

					content, err = os.ReadFile("outputMedium.txt")
					if err != nil {
						t.Error("but outputMedium.txt should've been created")
					}
					if string(content) != mediumContent {
						t.Errorf("but outputMedium.txt contained %s instead", string(content))
					}

					content, err = os.ReadFile("outputBig.txt")
					if err != nil {
						t.Error("but outputBig.txt should've been created")
					}
					if string(content) != bigContent {
						t.Errorf("but outputBig.txt contained %s instead", string(content))
					}
				}
			}
		}
	}
}

const smallContent = `=== Small ===
Teplice (49,705)
Chomutov (48,349)
Karlovy Vary (48,319)
Děčín (47,951)
Jablonec nad Nisou (45,317)
Mladá Boleslav (44,506)
Prostějov (43,381)
Přerov (42,451)
Třinec (34,778)
`

const mediumContent = `=== Medium ===
České Budějovice (94,229)
Hradec Králové (92,683)
Ústí nad Labem (91,982)
Pardubice (91,755)
Zlín (74,478)
Havířov (70,165)
Kladno (68,896)
Most (65,341)
Opava (55,996)
Frýdek-Místek (55,006)
Jihlava (51,125)
Karviná (50,902)
`

const bigContent = `=== Big ===
Prague (1,335,084)
Brno (382,405)
Ostrava (284,982)
Plzeň (175,219)
Liberec (104,261)
Olomouc (100,514)
`
