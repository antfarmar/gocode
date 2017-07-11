package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//rdr := os.Stdin
	rdr, err := os.Open("data.txt")
	defer rdr.Close()
	var n int
	var s string
	num, err := readData(rdr, "%v%v", &n, &s)
	fmt.Println(n, s, num, err)
}

func readData(r io.Reader, format string, a ...interface{}) (n int, err error) {
	return fmt.Fscanf(r, format, a...)
}
