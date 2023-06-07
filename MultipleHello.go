package main

import (
    "fmt"
    "log"
	"example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(1)

    names := []string{"Shreshtha", "Alia", "Katrina"}
	messages, err := greetings.Hellos(names)
    
    if err != nil {
        log.Fatal(err)
    }

	
    fmt.Println(messages)
}