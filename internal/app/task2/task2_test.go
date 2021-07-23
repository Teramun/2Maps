package task2

import "testing"

type Maps struct {
	mapOne, mapTwo, expectedOne, expectedTwo map[string]int
	Count int
}

func TestNCompare(t *testing.T) {
	tests := []Maps{
		{
			mapOne: map[string]int{"one": -1, "two": 2, "three": 3, "four": 4, "five": -5},
			mapTwo: map[string]int{},
			expectedOne: map[string]int{},
			expectedTwo: map[string]int{"one": -1, "two": 2, "three": 3, "four": 4, "five": -5},
		},
		{
			mapOne: map[string]int{},
			mapTwo: map[string]int{},
			expectedOne: map[string]int{},
			expectedTwo: map[string]int{},
		},
		{
			mapOne: map[string]int{"one": 1, "two": 1, "three": 1, "four": 1, "five": 1},
			mapTwo: map[string]int{"one": 1, "two": 1, "three": 1, "four": 1, "five": 1},
			expectedOne: map[string]int{},
			expectedTwo: map[string]int{},
		},
		{
			mapOne: map[string]int{"one": 1, "two": -1, "three": 1, "four": 1, "five": -1},
			mapTwo: map[string]int{"first": 0, "second": -1, "three": 1, "four": 1, "five": 155},
			expectedOne: map[string]int{"first": 0, "second": -1, "five": 155},
			expectedTwo: map[string]int{"one": 1, "two": -1, "five": -1},
		},
	}
	for index, value := range tests {
		count := index + 1
		resultOne, resultTwo := NCompare(value.mapOne, value.mapTwo)
		for k, v := range resultOne {
			if v != value.expectedOne[k] {
				t.Errorf("Test number %v failed.", count)
			}
		}
		for k, v := range resultTwo {
			if v != value.expectedTwo[k] {
				t.Errorf("Test number %v failed.", count)
			}
		}
	}
}

