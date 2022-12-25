package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"os"
)

func convertPhoto(filepath string) tgbotapi.FileBytes {
	photo, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo,
	}
}

func deletePhoto(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		return err
	}

	return nil
}

func sendMessage(bot *tgbotapi.BotAPI, chat int64, msgID int, text string) {
	res := tgbotapi.NewMessage(chat, text)
	res.ReplyToMessageID = msgID

	_, err := bot.Send(res)
	if err != nil {
		log.Fatal(err)
	}
}

func sendPhoto(bot *tgbotapi.BotAPI, chat int64, msgID int) {
	photoFileBytes := convertPhoto("internal/upload/обэмэ.jpg")

	uploadedPhoto := tgbotapi.NewPhoto(chat, photoFileBytes)
	uploadedPhoto.ReplyToMessageID = msgID

	_, err := bot.Send(uploadedPhoto)
	if err != nil {
		log.Fatal(err)
	}

	err = deletePhoto("internal/upload/обэмэ.jpg")
	if err != nil {
		log.Fatal(err)
	}
}
