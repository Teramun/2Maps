package main

import (
	"2maps/internal/app/task3"
	"fmt"
	"log"
)

func main() {
	user := task3.User{}
	user.SetName("Anya")
	err := user.SetPhone(79803086117)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(user.Name())
	fmt.Println(user.Phone())
}