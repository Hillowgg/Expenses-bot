package main

import (
    "log"
    "os"
    "strconv"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func resolveUpdate(update tgbotapi.Update) {
    if update.Message == nil {
        return
    }
    switch update.Message.Command() {
    case "start":
        startCommand(update)
    }
}

var bot *tgbotapi.BotAPI

// todo: create config loader?
func main() {
    var err error
    bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
    if err != nil {
        log.Fatalf("[FATAL] %v\n", err)
    }
    log.Printf("[INFO] Logged in as %v\n", bot.Self.UserName)

    updateConfig := tgbotapi.NewUpdate(0)
    timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
    updateConfig.Timeout = timeout

    updateChan := bot.GetUpdatesChan(updateConfig)
    log.Println("[INFO] Started polling")
    for update := range updateChan {
        go resolveUpdate(update)
    }
}
