package main

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startCommand(update tgbotapi.Update) {
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello world")
    _, err := bot.Send(msg)
    if err != nil {
        ErrorLog.Printf("Start command: %v\n", err)
        return
    }
    InfoLog.Printf("Start command send to %v\n", update.Message.From.UserName)
}
