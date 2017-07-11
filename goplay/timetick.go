package main

import "time"

func main() {
	for range time.Tick(time.Second) {
		println("tick")
	}
}
