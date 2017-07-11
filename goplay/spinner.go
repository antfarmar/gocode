// spinner
package main

import (
	"fmt"
	"time"
)

func main() {
	spinner := []string{`-`, `\`, `|`, `/`}
	for t := range time.Tick(time.Second) {
		i := int(t.Second())
		fmt.Printf("\r%v %v", spinner[i%4], t)
	}
}
