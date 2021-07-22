package task1

//prints non-unique key-value pairs
func Compare(mOne, mTwo map[string]int) map[string]int {
	newMap := make(map[string]int)
	for keyOne, valueOne := range mOne {
		for keyTwo, valueTwo := range mTwo {
			if keyOne == keyTwo && valueOne == valueTwo {
				newMap[keyTwo] = valueTwo
			}
		}
	}
	return newMap
}