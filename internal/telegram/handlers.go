package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func IsPrompt(text string) bool {
	if len(text) > 1 {
		return true
	} else {
		return false
	}
}

func (tb *TelegramBot) sendMessage(chat int64, msgID int, text string) {
	res := tgbotapi.NewMessage(chat, text)
	res.ReplyToMessageID = msgID

	_, err := tb.bot.Send(res)
	if err != nil {
		log.Fatal(err)
	}
}

func (tb *TelegramBot) sendPhoto(chat int64, msgID int, photoURL string) {
	convertedPhoto := tgbotapi.FileURL(photoURL)

	uploadedPhoto := tgbotapi.NewPhoto(chat, convertedPhoto)
	uploadedPhoto.ReplyToMessageID = msgID

	_, err := tb.bot.Send(uploadedPhoto)
	if err != nil {
		log.Fatal(err)
	}
}

func (tb *TelegramBot) sendImaginePhoto(chat int64, msgID int, text string) {
	if IsPrompt(text) {
		tb.sendMessage(chat, msgID, "Ожидайте, ваш запрос принят")

		photoURL, err := GetPromptURL(text)
		if err != nil {
			tb.sendMessage(chat, msgID, ErrImagineTimeOut)
		} else {
			tb.sendPhoto(chat, msgID, photoURL)
		}

	} else {
		tb.sendMessage(chat, msgID, ErrInvalidPrompt)
	}
}
