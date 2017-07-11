package main

import "fmt"

func main() {
	A := []byte("123456")
	fmt.Printf("A: %s\n", A)

	reverse(A)               // passes a copy of header(pointer), not bytes
	fmt.Printf("A: %s\n", A) // A is reversed in place

	B := reverse(A)          // B gets a copy of header of A.
	A[0] = byte('-')         // This modifies B too.
	fmt.Printf("A: %s\n", A) // A is reversed too.
	fmt.Printf("B: %s\n", B) // header(B)==header(A)

	B = reverse(A, "copy")   // A was deep copied
	A[0] = byte('1')         // B isn't mutated
	fmt.Printf("A: %s\n", A) // A stayes the same
	fmt.Printf("B: %s\n", B) // B is reversed
}

func reverse(r []byte, cpy ...string) []byte {
	//r := bytes.Runes(b)
	if cpy != nil && cpy[0] == "copy" {
		//r = []byte(string(r)) // also copies r
		s := make([]byte, len(r))
		copy(s, r)
		r = s
	}

	i, j := 0, len(r)-1
	for i < len(r)/2 {
		r[i], r[j] = r[j], r[i]
		i, j = i+1, j-1
	}
	return r
}
