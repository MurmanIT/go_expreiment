package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "/tmp/lGO.log"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: %v", err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "INFO: ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("This is a test log entry")
	iLog.Println("This is another test log entry")
}
