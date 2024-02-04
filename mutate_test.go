package weasel

import (
	"math/rand"
	"testing"
)

func TestMutate(test *testing.T) {
	rand.Seed(1)

	result := Mutate("the quick brown fox jumps over the lazy dog", 0.2)

	if result != "the qu Pk brown fox jumps oveF tGD Nazy dog" {
		test.Fail()
	}
}
