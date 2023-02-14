package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("a and b are equal")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError returned nil")
	} else {
		fmt.Println(err)
	}

	err = returnError(1, 1)
	if err == nil {
		fmt.Println("returnError returned nil")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "a and b are equal" {
		fmt.Println("!!")
	}
}
