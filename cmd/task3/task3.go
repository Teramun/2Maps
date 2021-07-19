//внутри get-метода имени пользователя менять все гласные на большие гласные
package main

import (
	"2maps/internal/app/task3"
	"fmt"
	"log"
)

func main() {
	user := task3.User{}
	user.SetName("elena")
	err := user.SetPhone(79803086117)
	if err != nil {
		log.Fatal()
	}
	name := user.Name()
	for i := 0; i < len(name); i++ {
		fmt.Print(string(name[i]))
	}
	fmt.Print("\n")
	fmt.Println(user.Phone())
}