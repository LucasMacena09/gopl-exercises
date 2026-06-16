// Exercise1_02 prints its command-line arguments along with their indices, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}