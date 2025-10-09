package main

import (
	"fmt"
	"log"

	"github.com/yuvbro/go-greetings"
)
import "rsc.io/quote"

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello, World from Brosh!")
	fmt.Println(quote.Go())

	//names := []string{"Yaron", "Shiran", "Lera", "Noa", "Asf"}

	msg, err := greetings.Hello("Nevo")
	//messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err) // Fatal is equivalent to Print followed by a call to os.Exit(1).
	}

	fmt.Println(msg)
	//fmt.Println(messages)
}
