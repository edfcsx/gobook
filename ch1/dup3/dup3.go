/*
	Dup3 exibe a contagem e o texto das linhas que aparecem mais de uma
	vez na entrada. Ele lê apartir de um arquivo, diferente do dup2
	é introduzida a função ioutil.ReadFile que lê todo o arquivo de uma
	única vez armazenando seu conteudo em um grande bloco de memoria
	diferente do dup2 que le os dados via streaming o dup3 dependente
	do tamanho do arquivo pode ter seu desempenho afetado.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
