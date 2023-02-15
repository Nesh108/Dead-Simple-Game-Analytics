package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"strings"
)

var (
	Token  string
	ChatId string
)

func getUrl(endpoint string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", Token, endpoint)
}

func SendTelegramMessages(text string) (bool, error) {
	chatIdsStr := os.Getenv("TELEGRAM_CHAT_IDS")
	chatIds := strings.Split(chatIdsStr, ";")

	for i := 0; i < len(chatIds); i++ {
		log.Println("chat id " + chatIds[i])
		SendTelegramMessage(text, chatIds[i])
	}

	return true, nil
}

func SendTelegramMessage(text string, chatId string) (bool, error) {
	if Token == "" {
		Token = os.Getenv("TELEGRAM_TOKEN")
	}

	var err error
	var response *http.Response

	url := getUrl("sendMessage")
	body, _ := json.Marshal(map[string]string{
		"chat_id": chatId,
		"text":    text,
	})

	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// Log
	log.Println("Message  was sent: " + text)
	log.Println("Response JSON: " + string(body))

	// Return
	return true, nil
}