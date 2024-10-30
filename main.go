package main

import (
	
	"context"
	"time"
	"log"
	"os"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"github.com/dgraph-io/badger"
)

var (
	telegramBotToken	= os.Getenv("TELEGRAM_BOT_TOKEN")
	pollInterval		= 10 * time.Minute
)

func main() {
    botOptions := []bot.Option{
	bot.WithMessageTextHandler("/help", bot.MatchTypeExact, handlerHelp),
	bot.WithMessageTextHandler("/start", bot.MatchTypeExact, handlerHelp),
    }

    TelegramBot := initializeBot(botOptions)
    telegramContext, cancelContext := context.WithCancel(context.Background())
    if telegramContext == nil {
	log.Fatal(`Fatal error starting bot:
            There is no context for the bot; exiting now.`)
    }
    TelegramBot.Start(telegramContext)
    defer cancelContext()
}


func initializeBot(options []bot.Option) *bot.Bot {
    if err := godotenv.Load(); err != nil {
        log.Fatal("-- Couldn't load .env file; error: ", err.Error())
    }

    telegramApiToken, found := os.LookupEnv("TELEGRAM_BOT_TOKEN") 
    if !found {
        log.Fatal("-- Couldn't load API_TOKEN variable from .env; exiting now.")
    }

    b, err := bot.New(telegramApiToken, options...)
    if err != nil {
        log.Fatal("-- Couldn't construct Telegram bot object; error: ", err.Error())
    }

    return b
}

func handlerHelp(ctx context.Context, telegramBot *bot.Bot, update *models.Update) {
    if update.Message.Text == "/start" {
	log.Printf(
	    "-- Chat ID: %s; /start command received",
	    strconv.FormatInt(update.Message.Chat.ID, 10),
	)
    }

    if _, err := telegramBot.SendMessage(ctx, &bot.SendMessageParams{
	ChatID:	update.Message.Chat.ID,
	Text:	"This bot will notify you whenever you're asked to review " +
		"a pull request.\n" +
		"It was made by @effygp from Stinkworks; feel free to " +
		"review our code or contact us here:\n" +
		"https://github.com/stinkworks/purrequest",
    }); err != nil {
	log.Print("Couldn't reply to /help or /start command; error: ", err.Error())
    }
}

// func checkForReviewRequests() {
// }
