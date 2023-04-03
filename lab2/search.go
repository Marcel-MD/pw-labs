package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type SearchResult struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type SearchResults struct {
	Items []SearchResult `json:"items"`
}

func search(args []string) {

	query := strings.Join(args, "+")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("GOOGLE_API_KEY")
	id := os.Getenv("SEARCH_ENGINE_ID")

	apiUrl := "https://www.googleapis.com/customsearch/v1?" +
		"key=" + key +
		"&cx=" + id +
		"&q=" + query +
		"&num=10" +
		"&gl=md"

	body, _ := get(apiUrl)
	body = body[strings.Index(body, "{") : strings.LastIndex(body, "}")+1]

	results := SearchResults{}
	err = json.Unmarshal([]byte(body), &results)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	for i := 0; i < len(results.Items); i++ {
		fmt.Print(i + 1)
		fmt.Println(". " + results.Items[i].Title)
		fmt.Println("    Link: " + results.Items[i].Link)
		fmt.Println()
	}

	n := 0

	for n <= 0 || n > 10 {
		fmt.Print("Type the number of the link you want to access: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Try again.")
			continue
		}
		n, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Try again.")
			continue
		}
		if n <= 0 || n > 10 {
			fmt.Println("Try again.")
		}
	}

	body, bodyType := get(results.Items[n-1].Link)
	processBody(body, bodyType)
}
