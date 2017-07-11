package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	//fileIn    = os.Stdin
	fileIn, _ = os.Open(`.\data\icecream.txt`)
	fileOut   = os.Stdout
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
)

// MAIN ==================================================
func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	data, _ := rw.ReadString('\n')
	T, _ := strconv.Atoi(strings.TrimSpace(data))

	for t := 0; t < T; t++ {
	}
	rw.WriteString("out\n")
}
