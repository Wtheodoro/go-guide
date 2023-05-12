package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
	response, err := http.Get("https://www.walisontsx.com/")
	
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// Só Deus sabe pq posso usar lw que é do typo logWriter como primeiro parametro de io.Copy que é do tipo io.Writter
	// Descobri! io.Writer é uma interface
	io.Copy(lw, response.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))

	return len(bs), nil
}