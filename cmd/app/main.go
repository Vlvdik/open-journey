package main

import (
	"log"
	"shtrafobot/internal/config"
	bot "shtrafobot/internal/telegram"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot.Init(cfg.TelegramToken)
}
