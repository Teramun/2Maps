package task1

import (
	"testing"
)

type maps struct {
	testName string
	numbersOne, numbersTwo, expected map[string]int
}
func TestCompare(t *testing.T) {
	tests := []maps{
	{
		testName: "multiple key-value matches",
		numbersOne: map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5},
		numbersTwo: map[string]int{"first": 1, "second": 2, "three": 3, "four": 4, "five": 0},
		expected: map[string]int{"three": 3, "four": 4},
	},
	{
		testName: "both maps are empty",
		numbersOne: map[string]int{},
		numbersTwo: map[string]int{},
		expected: map[string]int{},
	},
	{
		testName: "no matching key-value",
		numbersOne: map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5},
		numbersTwo: map[string]int{"first": 6, "second": 7, "three": 8, "four": 9, "five": 0},
		expected: map[string]int{},
	},
	}
	for _, value := range tests {
		result := Compare(value.numbersOne, value.numbersTwo)
		for keyResult, valueResult := range result {
			if valueResult != value.expected[keyResult] {
				t.Log("BAD TEST", value.testName)
			}
		}
	}
}



