package main

import (
	m "test/internal/app/task1"
)

func main() {
	mOne := make(map[string]int)
	mTwo := make(map[string]int)

	mOne = map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
	}

	mTwo = map[string]int {
		"one": 2,
		"two": 2,
		"three": 2,
		"four": 4,
	}
	m.Compare(mOne, mTwo)
}