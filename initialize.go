package weasel

import "math/rand"

func Initilize(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += string(makeRandomCharacter())
	}

	return result
}

func makeRandomCharacter() byte {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	randomIndex := rand.Intn(len(alphabet))
	return alphabet[randomIndex]
}
