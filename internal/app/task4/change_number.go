package task4

import "fmt"

//performs mathematical operations with the entered number
func ChangeNumber() {
	var number float64
	number = 2

	for i := 0; i < 20; i++ {
		number = number * 5 / 2
		fmt.Println(number)
	}
}

