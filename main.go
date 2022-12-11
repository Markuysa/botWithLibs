package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(mustToken(), opts...)
	if err != nil {
		panic(err)
	}
	log.Print("running the service")
	b.Start(ctx)
}

func mustToken() string {

	token := flag.String("telegram-token", "", "token for access to tgbot")

	flag.Parse()

	if *token == "" {
		log.Fatal()
	}
	return *token
}
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
