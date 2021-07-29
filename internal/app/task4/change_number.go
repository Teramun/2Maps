package task4

import "fmt"

//performs mathematical operations with the entered number
func ChangeNumber() {
	var number float64
	fmt.Scan(&number)

	for i := 0; i < 3; i++ {
		number = number * 5 / 2
	}

	fmt.Println(number, "\n")
}

