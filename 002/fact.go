package main

/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/
import (
	"fmt"
	"log"
)

func main() {
	var input int

	fmt.Print("Please enter a number : ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("Please enter a number")
	}
	result := Factorial(input)

	fmt.Printf("Factorial of %d = %d", input, result)
}

func Factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * Factorial(number-1)
}
