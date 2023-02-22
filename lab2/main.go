package main

import (
	"fmt"
)

func main() {

	// args := os.Args[1:]
	args := []string{"-u", "https://utm.md/"}
	// args := []string{"-s", "hello", "world"}

	if len(args) == 0 {
		help()
		return
	}

	switch args[0] {
	case "-u":
		if len(args) < 2 {
			help()
			return
		}
		url := args[1]
		get(url)
	case "-s":
		if len(args) < 2 {
			help()
			return
		}
		search(args[1:])
	default:
		help()
	}

}

func help() {
	fmt.Println("go2web -u <URL> 	# make an HTTP request to the specified URL and print the response")
	fmt.Println("go2web -s <search-term> # make an HTTP request to search the term using your favorite search engine and print top 10 results")
	fmt.Println("go2web -h 		# show this help")
}
