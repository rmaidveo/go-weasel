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
	current := weasel.Initilize(len(sample))
	for current != sample {
		fmt.Println(current)

		variants := weasel.Populate(current, 0.05, 100)
		current = weasel.Search(variants, sample)
	}

	fmt.Println(current)
}
