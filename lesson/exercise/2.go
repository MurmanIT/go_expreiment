package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum float64 = 0
	k := 1

	for scanner.Scan() {
		argument := scanner.Text()
		n, err := strconv.ParseFloat(argument, 64)

		if err == nil {
			sum += n
			k++
			fmt.Println("Sum:", sum)
			fmt.Println("Average:", sum/float64(k))
		}

		if argument == "END" {
			fmt.Println("END")
			os.Exit(1)
		}
	}
}
