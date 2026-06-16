// Exercise1_03 prints the execution time difference between the for loop and strings.Join versions.
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	forVersion()
	joinVersion()
}

func forVersion() {
	elements := make([]string, 20000)
	start := time.Now()
	s, sep := "", ""
	for _, arg := range elements {
		s += sep + arg
		sep = " "
	}
	fmt.Println(time.Since(start).Milliseconds())
}

func joinVersion() {
	elements := make([]string, 20000)
	start := time.Now()
	strings.Join(elements, " ")
	fmt.Println(time.Since(start).Milliseconds())
}
