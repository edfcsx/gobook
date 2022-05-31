/*
	Dup1 exibe o texto de toda linha que aparece mais de
	uma vez na entrada-padrão, precedida por sua
	contagem.

	para para o scanner é necessário
	apertar CTRL+D para indicar o fim da leitura.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
