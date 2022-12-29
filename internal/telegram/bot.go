package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"time"
)

type handlers interface {
	sendMessage(chat int64, msgID int, text string)
	sendPhoto(chat int64, msgID int, text string)
}

type methods interface {
	Start() error
	initUpdates() tgbotapi.UpdatesChannel
	handleUpdates(updates tgbotapi.UpdatesChannel)
}

type Functions interface {
	handlers
	methods
}

type TelegramBot struct {
	bot *tgbotapi.BotAPI
	Functions
}

func NewTelegramBot(token string) *TelegramBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	return &TelegramBot{bot: bot}
}

func (tb *TelegramBot) initUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return tb.bot.GetUpdatesChan(u)
}

func (tb *TelegramBot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			// frequently used params
			username := update.Message.From.UserName
			chat := update.Message.Chat.ID
			words := strings.Split(update.Message.Text, " ")
			msg := update.Message.Text
			msgID := update.Message.MessageID

			log.Printf("[%s] %s", username, msg)

			switch words[0] {
			case "/ping":
				tb.sendMessage(chat, msgID, "pong")
			case "/send":
				if len(words) > 1 {
					tb.sendMessage(chat, msgID, "Ожидайте, ваш запрос принят")
					go tb.sendPhoto(chat, msgID, msg)
				} else {
					tb.sendMessage(chat, msgID, errInvalidPrompt)
				}
			default:
				tb.sendMessage(chat, msgID, errUnknownCommand)
			}
		}
	}
}

func (tb *TelegramBot) Start() error {
	log.Printf("%s starts working\nStart time: %v", tb.bot.Self.UserName, time.Now().Format(time.RFC822))

	updates := tb.initUpdates()
	tb.handleUpdates(updates)

	return nil
}
