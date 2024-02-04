package weasel

import "testing"

func TestSearch(test *testing.T) {
	result := Search([]string{"cat", "dog", "pig"}, "dig")

	if result != "dog" {
		test.Fail()
	}
}
