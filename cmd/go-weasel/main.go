package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rmaidveo/go-weasel"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const sample = "METHINKS IT IS LIKE A WEASEL"
	var generationCounter int
	current := weasel.Initilize(len(sample))
	for current != sample {
		const outputRate = 10
		if generationCounter%outputRate == 0 {
			fmt.Println(generationCounter, current)
		}

		variants := weasel.Populate(current, 0.05, 100)
		current = weasel.Search(variants, sample)

		generationCounter++
	}

	fmt.Println(generationCounter, current)
}
