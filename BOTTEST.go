package main

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"

	"time"
)

var print = fmt.Println
var timenow = time.Now()

func connect(url string) int {
	res, err1 := http.Get(url)
	if err1 != nil {

		log.Fatal(err1)
	}
	StatusCode := res.StatusCode

	return StatusCode
}

func main() {
	url := "http://10.68.1.222/reports/reports/20181212/report-120002/"
	_ = url
	resurs, _ := http.Get(url)
	_ = resurs

	os.Setenv("HTTP_PROXY", "46.101.74.238:3128")
	bot, err := tgbotapi.NewBotAPI("701590135:AAGucb1_8W4m6wOqsM4kOL6mGJZgPclgeuo")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
