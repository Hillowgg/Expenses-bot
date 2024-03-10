package main

import (
    "log"
    "os"
    "strconv"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "github.com/joho/godotenv"

    "main/database"
)

var bot *tgbotapi.BotAPI
var (
    InfoLog  *log.Logger
    WarnLog  *log.Logger
    ErrorLog *log.Logger
    FatalLog *log.Logger
)
var DB *database.MyDB

func init() {
    flags := log.LstdFlags | log.Lshortfile
    InfoLog = log.New(os.Stdout, "[INFO] ", flags)
    WarnLog = log.New(os.Stdout, "[WARN] ", flags)
    ErrorLog = log.New(os.Stderr, "[ERROR] ", flags)
    FatalLog = log.New(os.Stderr, "[FATAL] ", flags)
    err := godotenv.Load(".env")
    if err != nil {
        FatalLog.Fatalf("Loading .env %v\n", err)
    }
    DB, err = database.NewDB("Postgres", os.Getenv("DATABASE_DSN"))
    if err != nil {
        FatalLog.Fatalf("Connect to db %v\n", err)
    }
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
func resolveUpdate(update tgbotapi.Update) {
    if update.Message == nil {
        return
    }
    switch update.Message.Command() {
    case "start":
        StartCommand(update)
    }
}
