package main

import (
	"log"
	"shtrafobot/internal/configs"
	bot "shtrafobot/internal/telegram"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	tb := bot.NewTelegramBot(cfg.TelegramToken, cfg.XRapidAPIKey)

	err = tb.Start()
	if err != nil {
		log.Fatal(err)
	}
}
