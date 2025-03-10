package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var arrays = []int{100, 101, 102, 103, 200, 201, 202, 203, 204, 401, 402, 403, 404, 500, 501, 502, 503, 504, 505, 506, 507, 508}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Токен бота не найден в .env файле")
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			text := update.Message.Text
			chatId := update.Message.Chat.ID

			switch text {
			case "/start":
				msg := "Получить новый мем: /newMeme"
				bot.Send(tgbotapi.NewMessage(chatId, msg))

			case "/newMeme":
				code := randomNumber()
				photoURL := fmt.Sprintf("https://http.cat/status/%d", code)

				if isValidImage(photoURL) {
					photoMsg := tgbotapi.NewPhoto(chatId, tgbotapi.FileURL(photoURL))
					photoMsg.Caption = fmt.Sprintf("HTTP Cat %d", code)
					bot.Send(photoMsg)
				} else {
					bot.Send(tgbotapi.NewMessage(chatId, "Ошибка загрузки картинки"))
				}
			default:
				msg := "Неверная команда! пожалуйста используйте /start или /newMeme"
				bot.Send(tgbotapi.NewMessage(chatId, msg))
			}
		}
	}
}

func randomNumber() int {
	return arrays[rand.Intn(len(arrays))]
}

func isValidImage(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
