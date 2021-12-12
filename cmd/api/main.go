package main

import (
	"fmt"
	"todolist/server"
)

func main() {
	if err := server.RunApp(); err != nil {
		fmt.Println("Error!")
	}
}
