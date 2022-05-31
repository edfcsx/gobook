/*
	Fetch exibe o contéudo encontrado em cada URL especificada
	e armazena no buffer os.Stdout
	exemplo de url: http://gopl.io

	esse código contém os exercicios:
		1.7
		1.8
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		requstURL := mountHttpURL(url)
		resp, err := http.Get(requstURL)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("test %d\t%s\n", resp.StatusCode, requstURL)

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy body to buffer %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func mountHttpURL(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}

	return "http://" + url
}
