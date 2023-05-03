package main

import (
	"encoding/json"
	"fmt"
	"lab5/config"
	"lab5/models"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var cfg = config.New()

func main() {
	http.HandleFunc("/", updateHandler)

	fmt.Printf("Listenning on port %s\n", cfg.Port)
	fmt.Printf("Access URL: %s%s/setWebhook?url=<NGROK_URL>\n", cfg.BotUrl, cfg.BotToken)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	message := &models.ReceiveMessage{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		fmt.Println(err)
	}

	chatID := message.Message.Chat.ID
	userID := message.Message.From.ID
	msgText := message.Message.Text

	args := strings.Split(msgText, " ")
	command := args[0]

	switch command {
	case "/start":
		sendMessage(chatID, "Hello, I'm a news bot.")
	case "/latest_news":
		showLatestNews(chatID, args[1:])
	case "/save_news":
		saveNews(chatID, userID, args[1:])
	case "/saved_news":
		showSavedNews(chatID, userID)
	default:
		sendMessage(chatID, msgText)
	}
}

func showLatestNews(chatID int, args []string) {
	topic := strings.Join(args, " ")
	if topic == "" {
		topic = "Moldova"
	}

	newUrl := fmt.Sprintf("%s/everything?apiKey=%s&q=%s&pageSize=%d", cfg.NewsUrl, cfg.NewsApiKey, url.QueryEscape(topic), 5)
	fmt.Println(newUrl)

	resp, err := http.Get(newUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	news := &models.NewsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&news)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, article := range news.Articles {
		msg := fmt.Sprintf("%s\n\n%s", article.Title, article.URL)
		sendMessage(chatID, msg)
	}
}

func saveNews(chatID int, userID int, args []string) {
	sendMessage(chatID, "Not implemented yet")
}

func showSavedNews(chatID int, userID int) {
	sendMessage(chatID, "Not implemented yet")
}

func sendMessage(chatID int, text string) {

	msgUrl := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", cfg.BotUrl, cfg.BotToken, chatID, url.QueryEscape(text))

	_, err := http.Get(msgUrl)
	if err != nil {
		fmt.Println(err)
	}
}
