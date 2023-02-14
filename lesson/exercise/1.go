package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats")
		os.Exit(1)
	}
	arguments := os.Args

	var sum float64 = 0
	k := 1

	for i := 1; i < len(arguments); i++ {
		p, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Println("This is not number")
		} else {
			sum += p
			k++
		}
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", sum/float64(k))
}
