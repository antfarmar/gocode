package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()

	println("Processors:", maxProcs)
	println("Cores:", numCPU)

	for i := numCPU; i > 1; i-- {
		go say("go", i)
	}

	say("Main", 1) // When main terminates, so do goroutines
}

func say(s string, n int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(n*1000) * time.Millisecond)
		fmt.Printf("[i=%d] Proc #%d: %s\n", i, n, s)
	}
}
