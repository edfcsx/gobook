/*
	dup4 é igual a dup2 porém exibe o nome dos arquivos onde o texto
	foi encontrado.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, file := range files {
		f, err := os.Open(file)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
			continue
		}

		countLines(f, counts, file)
		f.Close()
	}

	print(counts)
}

func countLines(f *os.File, counts map[string]int, filename string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		for _, word := range strings.Split(input.Text(), " ") {
			counts[filename+"-"+word]++
		}
	}
}

func print(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			wordBuffer := strings.Split(line, "-")
			filename := wordBuffer[0]
			word := wordBuffer[1]

			fmt.Printf("%d\t%s\t%s\n", n, filename, word)
		}
	}
}
