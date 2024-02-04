package weasel

import "math/rand"

func Mutate(text string, rate float64) string {
	var result string
	for _, character := range text {
		if rand.Float64() < rate {
			result += string(makeRandomCharacter())
		} else {
			result += string(character)
		}
	}

	return result
}
