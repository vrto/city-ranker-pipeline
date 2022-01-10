package parallel_mapper

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strconv"
)

type City struct {
	name       string
	population int
}

func NewCity(input string) *City {
	var (
		start, end int
	)
	for idx, char := range input {
		if char == '(' {
			start = idx
		}
		if char == ')' {
			end = idx
		}
	}
	if end == 0 {
		return nil
	}
	cityPart := input[:start-1]

	number := ""
	for _, char := range input[start+1 : end] {
		if char != ',' {
			number += string(char)
		}
	}

	population, err := strconv.Atoi(number)
	if err != nil {
		return nil
	}

	return &City{
		name:       cityPart,
		population: population,
	}
}

func (this *City) String() string {
	populationPart := message.NewPrinter(language.English).Sprintf("%d", this.population)
	return fmt.Sprintf("%s (%s)", this.name, populationPart)
}
