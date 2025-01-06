// Ex004 takes a string of comma-separated numbers and returns a slice of int

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := Ex004("123456789")
	fmt.Print(res)
}

func Ex004(input string) []int {
	numbers := strings.Split(input, ",")

	length := len(numbers)

	var num = make([]int, length)

	for index, v := range numbers {
		s := strings.Trim(v, " ")
		num[index], _ = strconv.Atoi(s)
	}
	return num
}
