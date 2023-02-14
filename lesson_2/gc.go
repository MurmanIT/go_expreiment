package main

import (
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	println("Alloc:", mem.Alloc)
	println("TotalAlloc:", mem.TotalAlloc)
	println("HeapAlloc:", mem.HeapAlloc)
	println("Sys:", mem.Sys)
	println("NumGC:", mem.NumGC)
	println("NumForcedGC:", mem.NumForcedGC)
}
func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			println("Operation failed")
		}
		time.Sleep(5 * time.Second)
		/*println("=====================================")
		printStats(mem)*/
	}
}
