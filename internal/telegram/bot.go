package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func listen(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			// frequently used params
			username := update.Message.From.UserName
			chat := update.Message.Chat.ID
			msg := update.Message.Text
			msgID := update.Message.MessageID

			log.Printf("[%s] %s", username, msg)

			switch msg {
			case "/ping":
				sendMessage(bot, chat, msgID, "pong")
			case "/send":
				sendPhoto(bot, chat, msgID)
			default:
				sendMessage(bot, chat, msgID, "Unknown command")
			}
		}
	}
}

func Init(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	listen(bot, updates)
}
