package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/rmaidveo/go-weasel"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sample := flag.String("sample", "METHINKS IT IS LIKE A WEASEL", "target string")
	rate := flag.Float64("rate", 0.05, "character mutate rate")
	populationCount := flag.Int("count", 100, "population size")
	flag.Parse()

	var generationCounter int
	current := weasel.Initilize(len(*sample))
	for current != *sample {
		const outputRate = 10
		if generationCounter%outputRate == 0 {
			fmt.Println(generationCounter, current)
		}

		variants := weasel.Populate(current, *rate, *populationCount)
		current = weasel.Search(variants, *sample)

		generationCounter++
	}

	fmt.Println(generationCounter, current)
}
