package main

import (
	"fmt"

	"github.com/Adilfarooque/Wire/greeting"
)

func main() {
	greetingService := greeting.InitializeGreetingService()
	msg := greetingService.Greet("Gojo Saturu")
	fmt.Println(msg)
}
