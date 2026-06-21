// Exercise1_04 prints the count and text of lines that appear more than
// once in the input, reading from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, fileCounts := range counts {
		var sources []string
		var total int

		for file, n := range fileCounts {
			sources = append(sources, file)
			total += n
		}

		if total > 1 {
			fmt.Printf("%d\t%s\t%s\n", total, line, strings.Join(sources, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		_, ok := counts[input.Text()]
		if !ok {
			counts[input.Text()] = make(map[string]int)
		}

		counts[input.Text()][f.Name()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dup2: error reading %s: %v\n", f.Name(), err)
	}
}