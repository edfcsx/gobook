/*
	Fetch salva o body de cada url especificada em um arquivo
	e informa o status de cada request.
	exemplo de url: http://gopl.io
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		requestURL := mountHttpURL(url)
		resp, err := http.Get(requestURL)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(mountOutputURL(requestURL), readResponseBody(resp), 0644)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error write file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("test %d\t%s\n", resp.StatusCode, requestURL)
	}
}

func mountHttpURL(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}

	return "http://" + url
}

func readResponseBody(response *http.Response) []byte {
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch erro read body content: %v", err)
		os.Exit(1)
	}

	response.Body.Close()
	return body
}

func getExecutableDir() string {
	pwd, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch error get current dir: %v\n", err)
		os.Exit(1)
	}

	return pwd
}

func mountOutputURL(requestURL string) string {
	pwd := getExecutableDir()
	var sb strings.Builder
	sb.WriteString(pwd)
	sb.WriteString("/")

	var protocol string

	if strings.HasPrefix(requestURL, "https://") {
		protocol = "https://"
	} else if strings.HasPrefix(requestURL, "http://") {
		protocol = "http://"
	}

	sb.WriteString(strings.Split(requestURL, protocol)[1])
	sb.WriteString(".html")

	return sb.String()
}
