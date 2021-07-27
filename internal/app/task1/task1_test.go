package task1

import (
	"testing"
)

type maps struct {
	testName string
	numbersOne, numbersTwo, expected map[string]int
}
func TestCompare(t *testing.T) {
	t.Log("Start")
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
	t.Log("range test cases:", tests, "\n")
	for index, value := range tests {
		count := index + 1

		t.Logf("test case number %v: %s", count, value.testName)
		t.Log("maps for comparison: first map:", value.numbersOne, ", second map:", value.numbersTwo)

		result := Compare(value.numbersOne, value.numbersTwo)

		t.Log("range result for key-value comparison with expected. Result:", result, ", Expected:", value.expected)
		for keyResult, valueResult := range result {
			if valueResult != value.expected[keyResult]  {
				t.Error("BAD TEST (different values)", value.testName)
			}
		}
		t.Logf("check maps by length: lenght result - %v, lenght expected - %v", len(result), len(value.expected))
		if len(result) != len(value.expected) {
			t.Log("BAD TEST (different maps len)", value.testName)
		}
		t.Log("\n")
	}
	t.Log("Finish")
}