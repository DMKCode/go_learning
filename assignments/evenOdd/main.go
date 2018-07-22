package main

import (
	"fmt"
)

func main() {
	ints := []int{}

	for i := 0; i < 11; i++ {
		ints = append(ints, i)
	}

	for _, value := range ints {
		if value%2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}
