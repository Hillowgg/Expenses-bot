package main

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommand(update tgbotapi.Update) {
    user := update.SentFrom()
    AddUser(user)
    chat := update.Message.Chat.ID

    msg := tgbotapi.NewMessage(chat, "Welcome to expenses bot!\nWrite /help for help")
    _, err := bot.Send(msg)
    if err != nil {
        ErrorLog.Printf("Failed to send start answer %v\n", err)
        return
    }

    InfoLog.Printf("Sent Start message to %v (%v)\n", user.UserName, user.ID)
}
