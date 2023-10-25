// call code from another module

package main

import (
	// std libs imports
	"fmt"
	"log"
	// importing local greetings module
	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	// use Hello() from greetings module, get message and print
	message1 := greetings.Hello("prasanna")
	message2, err := greetings.HelloE("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("greeting with name: ", message1, "greeting without name: ", message2)
}