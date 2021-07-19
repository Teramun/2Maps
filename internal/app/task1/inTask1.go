package compare

import "fmt"

//prints non-unique key-value pairs
func Compare(mOne, mTwo map[string]int) {
	for keyOne, valueOne := range mOne {
		for keyTwo, valueTwo := range mTwo {
			if keyOne == keyTwo && valueOne == valueTwo {
				fmt.Println(keyOne, valueOne)
			}
		}
	}
}
