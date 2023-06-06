package main

import "fmt"

func main() {
    // Define the variables
    var num1, num2 int

    // Prompt for the first number
    fmt.Print("Enter the first number: ")
    fmt.Scanln(&num1)

    // Prompt for the second number
    fmt.Print("Enter the second number: ")
    fmt.Scanln(&num2)

    // Calculate the sum
    sum := num1 + num2

    // Print the result
    fmt.Printf("Sum: %d\n", sum)
}