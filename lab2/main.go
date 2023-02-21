package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	var url string

	switch args[0] {
	case "-u":
		if len(args) < 2 {
			printHelp()
			return
		}
		url = args[1]
	case "-s":
		if len(args) < 2 {
			printHelp()
			return
		}
		url = "www.google.com/search?q=" + strings.Join(args[1:], "+")
	default:
		printHelp()
	}

	fmt.Println(get(addHttpsPrefix(url)))
}

func addHttpsPrefix(url string) string {
	if !strings.HasPrefix(url, "https://") {
		return "https://" + url
	}
	return url
}

func printHelp() {
	fmt.Println("go2web -u <URL> 	# make an HTTP request to the specified URL and print the response")
	fmt.Println("go2web -s <search-term> # make an HTTP request to search the term using your favorite search engine and print top 10 results")
	fmt.Println("go2web -h 		# show this help")
}
