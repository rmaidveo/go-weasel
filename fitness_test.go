package weasel

import "testing"

func TestFitness_withDifference(test *testing.T) {
	result := Fitness(
		"the quick brown fox jumps over the lazy dog",
		"the quick brown cat jumps over the lazy pig",
	)

	if result != 5 {
		test.Fail()
	}
}

func TestFitness_withNoDifference(test *testing.T) {
	result := Fitness(
		"the quick brown fox jumps over the lazy dog",
		"the quick brown fox jumps over the lazy dog",
	)

	if result != 0 {
		test.Fail()
	}
}
