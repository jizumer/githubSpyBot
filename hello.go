package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(fmt.Errorf("found an error creating the bot: %w", err))
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		log.Println("Received /hello command")
		return c.Send("Hello!")
	})

	b.Start()
}
