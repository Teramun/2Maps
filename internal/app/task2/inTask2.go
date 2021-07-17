package task2

import "fmt"

func NCompare(mapOne, mapTwo map[string]int) {
	for keyTwo, valueTwo := range mapTwo {
		if valueTwo != mapOne[keyTwo] {
			fmt.Println(keyTwo, valueTwo)
		}
	}
	for keyOne, valueOne := range mapOne {
		if valueOne != mapTwo[keyOne] {
			fmt.Println(keyOne, valueOne)
		}
	}
}