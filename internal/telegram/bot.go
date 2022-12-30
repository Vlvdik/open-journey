package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"time"
)

const help = "Привет!\nЧтобы получить картинку по запросу используй следующий синтаксис:\n/imagine <ваш запрос>\nВ целях соблюдения закона и гуманности, нейросеть не будет отображать результаты по запросам эротического и насильственного характера"

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
			case "/help":
				tb.sendMessage(chat, msgID, help)
			case "/imagine":
				go tb.sendImaginePhoto(chat, msgID, msg)
			default:
				tb.sendMessage(chat, msgID, ErrUnknownCommand)
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
