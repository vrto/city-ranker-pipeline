package main

import (
	. "parallel-mapper"
)

func main() {
	pipeline := NewCityRankingPipeline("../../input_file.txt", 8)
	pipeline.Clean()
	pipeline.Run()
}
