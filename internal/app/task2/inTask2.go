package task2

//prints unique key-value pairs
func NCompare(mapOne, mapTwo map[string]int) (map[string]int, map[string]int) {
	newMap := make(map[string]int)
	newMapTwo := make(map[string]int)
	for keyTwo, valueTwo := range mapTwo {
		if valueTwo != mapOne[keyTwo] {
			newMap[keyTwo] = valueTwo
		}
	}
	for keyOne, valueOne := range mapOne {
		if valueOne != mapTwo[keyOne] {
			newMapTwo[keyOne] = valueOne
		}
	}
	return newMap, newMapTwo
}