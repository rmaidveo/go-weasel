package weasel

import (
	"math/rand"
	"testing"
)

func TestInitiliaze(test *testing.T) {
	rand.Seed(1)

	result := Initilize(5)

	if result != "OPCLE" {
		test.Fail()
	}
}
