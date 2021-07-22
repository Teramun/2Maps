package task1

import "testing"

type maps struct {
	numbersOne, numbersTwo, expected map[string]int
}
func TestCompare(t *testing.T) {
	tests := []maps{
	{
		numbersOne: map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5},
		numbersTwo: map[string]int{"first": 1, "second": 2, "three": 3, "four": 4, "five": 0},
		expected: map[string]int{"three": 3, "four": 4},
	},
	{
		numbersOne: map[string]int{"brown": 1, "orange": 2, "yellow": 0, "green": 4, "blue": 5},
		numbersTwo: map[string]int{"blue": 1, "red": 2, "yellow": 0, "green": 4, "five": 0},
		expected: map[string]int{"yellow": 0, "green": 4},
	},
	{
		numbersOne: map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5},
		numbersTwo: map[string]int{"first": 6, "second": 7, "three": 8, "four": 9, "five": 0},
		expected: map[string]int{},
	},
	}
	for _, value := range tests {
		result := Compare(value.numbersOne, value.numbersTwo)

		for k, v := range result {
			if v != value.expected[k] {
			t.Errorf("bb")
			}
		}
	}
}



