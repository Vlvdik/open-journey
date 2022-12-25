package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
)

func sendMessage(bot *tgbotapi.BotAPI, chat int64, msgID int, text string) {
	res := tgbotapi.NewMessage(chat, text)
	res.ReplyToMessageID = msgID

	_, err := bot.Send(res)
	if err != nil {
		log.Fatal(err)
	}
}

func sendPhoto(bot *tgbotapi.BotAPI, chat int64, msgID int) {
	photo, err := ioutil.ReadFile("internal/upload/picture.jpg")
	if err != nil {
		log.Fatal(err)
	}

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo,
	}

	uploadedPhoto := tgbotapi.NewPhoto(chat, photoFileBytes)
	uploadedPhoto.ReplyToMessageID = msgID

	_, err = bot.Send(uploadedPhoto)
	if err != nil {
		log.Fatal(err)
	}
}
