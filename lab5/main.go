package main

import (
	"encoding/json"
	"fmt"
	"lab5/config"
	"lab5/models"
	"log"
	"net/http"
)

var cfg = config.New()

func main() {
	http.HandleFunc("/", updateHandler)

	fmt.Printf("Listenning on port %s\n", cfg.Port)
	fmt.Printf("Access URL: %s%s/setWebhook?url=<NGROK_URL>\n", cfg.Url, cfg.BotToken)

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
	msgText := message.Message.Text

	sendMessage(chatID, msgText)
}

func sendMessage(chatId int, text string) {
	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", cfg.Url, cfg.BotToken, chatId, text)

	_, err := http.Get(respMsg)
	if err != nil {
		fmt.Println(err)
	}
}
