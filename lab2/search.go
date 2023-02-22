package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	g "github.com/serpapi/google-search-results-golang"
)

func search(args []string) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	key := os.Getenv("KEY")

	parameter := map[string]string{
		"q":        strings.Join(args, " "),
		"location": "Moldova",
		"num":      "10",
	}

	query := g.NewGoogleSearch(parameter, key)
	rsp, err := query.GetJSON()
	if err != nil {
		fmt.Println(err)
	}

	results := rsp["organic_results"].([]interface{})
	for _, result := range results {
		fmt.Println(result.(map[string]interface{})["title"])
		fmt.Println(result.(map[string]interface{})["link"])
		fmt.Println()
	}
}
