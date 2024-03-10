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
var (
    InfoLog  *log.Logger
    WarnLog  *log.Logger
    ErrorLog *log.Logger
    FatalLog *log.Logger
)

func init() {
    flags := log.LstdFlags | log.Lshortfile
    InfoLog = log.New(os.Stdout, "[INFO] ", flags)
    WarnLog = log.New(os.Stdout, "[WARN] ", flags)
    ErrorLog = log.New(os.Stderr, "[ERROR] ", flags)
    FatalLog = log.New(os.Stderr, "[FATAL] ", flags)
}

// todo: create config loader?
func main() {
    var err error
    bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
    if err != nil {
        FatalLog.Fatalln(err)
    }
    InfoLog.Printf("Logged in as %v\n", bot.Self.UserName)

    updateConfig := tgbotapi.NewUpdate(0)
    timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
    updateConfig.Timeout = timeout

    updateChan := bot.GetUpdatesChan(updateConfig)
    InfoLog.Println("Started polling")
    for update := range updateChan {
        go resolveUpdate(update)
    }
}
