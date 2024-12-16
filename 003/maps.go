/*With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value

Suppose the following input is supplied to the program: 8
Then, the output should be:
map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]*/

package main

import (
	"fmt"
	"log"
)

func main() {
	var input int
	fmt.Println("Enter a number : ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal("Enter a number : ")
	}

	result := NumbertoMap(input)

	fmt.Print(result)
}

func NumbertoMap(n int) map[int]int {
	m := make(map[int]int)

	for i := 1; i <= n; i++ {
		m[i] = i * i
	}
	return m
}
