package task4

import "fmt"

//output a numeric matrix
func Matrix() {
	matrix := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 2, 3, 4, 5},
		[]int{3, 3, 3, 4, 5},
		[]int{4, 4, 4, 4, 5},
		[]int{5, 5, 5, 5, 5},
	}

	for indexOne := 0; indexOne < len(matrix); {
		for indexTwo := 0; indexTwo < len(matrix); indexTwo++ {
			var numbers []int
			number := matrix[indexOne][indexTwo]

			numbers = append(numbers, number)

			for index := 0; index < len(numbers); index++ {
				fmt.Print(numbers[index], " ")
			}
		}
		fmt.Print("\n")
		indexOne++
	}
}