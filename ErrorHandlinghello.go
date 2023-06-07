package main

import (
    "fmt"
    "log"
	"example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(1)

    
    message, err := greetings.HelloNew("Glady")
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}