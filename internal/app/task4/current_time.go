package task4

import (
	"fmt"
	"time"
)

//displays the current time with an interval of two seconds
func CurrentTime() {
	pause := time.Tick(2 * time.Second)

	for timeNow := range pause {
		fmt.Println(timeNow.Format("15:04:05"))
	}
}
