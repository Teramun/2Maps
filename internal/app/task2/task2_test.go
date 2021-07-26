package task2

import "testing"

type Maps struct {
	testName string
	mapOne, mapTwo, expectedOne, expectedTwo map[string]int
	Count int
}

func TestNCompare(t *testing.T) {
	tests := []Maps{
		{
			testName: "empty second map",
			mapOne: map[string]int{"one": -1, "two": 2, "three": 3, "four": 4, "five": -5},
			mapTwo: map[string]int{},
			expectedOne: map[string]int{},
			expectedTwo: map[string]int{"one": -1, "two": 2, "three": 3, "four": 4, "five": -5},
		},
		{
			testName: "both maps are empty",
			mapOne: map[string]int{},
			mapTwo: map[string]int{},
			expectedOne: map[string]int{},
			expectedTwo: map[string]int{},
		},
		{
			testName: "both maps are the same",
			mapOne: map[string]int{"one": 1, "two": 1, "three": 1, "four": 1, "five": 1},
			mapTwo: map[string]int{"one": 1, "two": 1, "three": 1, "four": 1, "five": 1},
			expectedOne: map[string]int{},
			expectedTwo: map[string]int{},
		},
		{
			testName: "multiple key-value matches",
			mapOne: map[string]int{"one": 1, "two": -1, "three": 1, "four": 1, "five": -1},
			mapTwo: map[string]int{"first": 0, "second": -1, "three": 1, "four": 1, "five": 155},
			expectedOne: map[string]int{"first": 0, "second": -1, "five": 155},
			expectedTwo: map[string]int{"one": 1, "two": -1, "five": -1},
		},
	}
	for _, value := range tests {
		resultOne, resultTwo := NCompare(value.mapOne, value.mapTwo)
		for keyResultOne, valueResultOne := range resultOne {
			if valueResultOne != value.expectedOne[keyResultOne] {
				t.Log("BAD TEST", value.testName)
			}
		}
		for keyResultTwo, valueResultTwo := range resultTwo {
			if valueResultTwo != value.expectedTwo[keyResultTwo] {
				t.Log("BAD TEST", value.testName)
			}
		}
	}
}

