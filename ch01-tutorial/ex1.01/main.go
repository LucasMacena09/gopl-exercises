// Exercise1_01 prints its command name and command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}