package service

import (
	"fmt"
	"log"
	"os"
	"questbotAI/internal/config"
	"strings"

	"github.com/NicoNex/echotron/v3"
	"github.com/joho/godotenv"
)

type bot struct {
	chatID int64
	echotron.API
}

func (b *bot) Update(update *echotron.Update) {
	if update.Message == nil {
		return
	}

	db := ConnectDB()

	switch update.Message.Text {
	case "/start":

		_, err := b.SendMessage(config.Welcome, b.chatID, nil)

		if err != nil {
			log.Println("Ошибка отправки сообщения:", err)
		}
	default:

		if _, err := GetMainPromt(db, int(b.chatID)); err != nil {
			userPromt := update.Message.Text
			lastedPromts, _ := GetMove(db, int(b.chatID))
			mainPront := MainPromtCombinator(userPromt, lastedPromts)
			InsertUser(db, int(b.chatID), mainPront)
			fmt.Printf("\n Ошибка (%s) в Get mainPromt создаю нового user\n", err)
		}

		InsertMove(db, int(b.chatID), update.Message.Text)
		mainPromt, _ := GetMainPromt(db, int(b.chatID))
		lastedPromts, err := GetMove(db, int(b.chatID))

		promt := mainPromt + strings.Join(lastedPromts, " ")

		fmt.Printf("\n GetMove ОШИБКА!(%s)\n", err)

		fmt.Printf("\n GetMainPromt(%s)\n", promt)
		message := RequesGroq(promt, update.Message.Text)

		b.SendMessage(message, b.chatID, nil)
		InsertMove(db, int(b.chatID), message)
	}
}

func StartBot() {
	err := godotenv.Load("../../internal/config/.env")
	if err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("TELETOKEN")
	if token == "" {
		log.Fatal("Токен бота не указан")
	}

	dsp := echotron.NewDispatcher(token, func(chatID int64) echotron.Bot {
		return &bot{
			chatID: chatID,
			API:    echotron.NewAPI(token),
		}
	})

	log.Println("Бот запущен...")
	log.Fatal(dsp.Poll())
}

func MainPromtCombinator(userPront string, lastedPromts []string) string {
	return fmt.Sprintf(config.MainPromt, userPront, strings.Join(lastedPromts, " "))
}
