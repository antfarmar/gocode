package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(r byte) byte {
	switch {
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}

func (rr *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	for i := range b[:n] {
		b[i] = rot13(b[i])
	}
	return n, io.EOF
}

func main() {
	sr := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{sr}
	n, err := io.Copy(os.Stdout, &r)
	fmt.Printf("\nBytes copied: %v\nError: %v\n", n, err)
}
