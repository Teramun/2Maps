package compare

import "fmt"

func Compare(mOne, mTwo map[string]int) {
	for keyOne, valueOne := range mOne {
		for keyTwo, valueTwo := range mTwo {
			if keyOne == keyTwo && valueOne == valueTwo {
				fmt.Println(keyOne, valueOne)
			}
		}
	}
}
