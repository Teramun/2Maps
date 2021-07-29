package main

import (
	"2maps/internal/app/task4"
	"time"
)

func main() {
	go task4.Matrix()
	go task4.CurrentTime()
	go task4.ChangeNumber()
	time.Sleep(20 * time.Second)
}

