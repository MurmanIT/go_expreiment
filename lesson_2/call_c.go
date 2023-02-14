package main

// #include <stdio.h>
// void callC() {
//     printf("Calling C code!\n");
// }
import "C"
import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Returned to Go code!")
}
