package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gookit/color"
)

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func validateInputs() (string, string, error) {
	// Validate arguments
	if len(os.Args) < 2 {
		return "", "", errors.New("a filepath argument is required")
	}

	fileLocation := flag.String("source", "foo", "source url of text")
	searchString := flag.String("query", "foo", "query string")

	flag.Parse() // actually performs the parse.  pointers etc.

	return *searchString, *fileLocation, nil // de-reference from the pointer
}

func iterateAndSearch(fullText_tokens []string, searchString string) {

	for index, line := range fullText_tokens {
		if strings.Contains(line, searchString) {
			color.Green.Printf(" %-7v ", fmt.Sprint(index))
			color.Println(strings.Replace(line, searchString, "<red>"+searchString+"</>", -1))
		}
	}
}

func main() {
	color.Green.Printf("*** %s ***\n", "Challenge3 Search CLI")

	searchString, fileLocation, err := validateInputs()

	if err != nil {
		exitGracefully(err)
	}

	color.Print("QUERY:<red>" + searchString + "</>")
	fmt.Printf("  SRC:%v\n", fileLocation)

	response, err := http.Get(fileLocation)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	fullText, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fullText_tokens := strings.Split(string(fullText), "\n")

	iterateAndSearch(fullText_tokens, searchString)

	// Copy data from the response to standard output
	// n, err1 := io.Copy(os.Stdout, response.Body) //use package "io" and "os"
	// if err != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	// fmt.Println("Number of bytes copied to STDOUT:", n)
}
