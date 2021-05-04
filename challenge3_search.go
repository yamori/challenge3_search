package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func validateInputs() (string, error) {
	// Validate arguments
	if len(os.Args) < 2 {
		return "", errors.New("a filepath argument is required")
	}

	fileLocation := flag.String("source", "foo", "source url of text")

	flag.Parse()

	return *fileLocation, nil
}

func main() {
	fmt.Println("*** Challenge3 Search CLI ***")

	fileLocation, err := validateInputs()

	if err != nil {
		exitGracefully(err)
	}

	fmt.Println(fileLocation)
	response, err := http.Get(fileLocation)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err1 := io.Copy(os.Stdout, response.Body) //use package "io" and "os"
	if err != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println("Number of bytes copied to STDOUT:", n)
}
