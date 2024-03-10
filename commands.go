package main

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommand(msg *tgbotapi.Message) {
    user := msg.From
    added := AddUser(user)
    chat := msg.Chat.ID
    var toSend tgbotapi.MessageConfig
    if added {
        toSend = tgbotapi.NewMessage(chat, "Welcome to expenses bot!\nWrite /help for help")
    } else {
        toSend = tgbotapi.NewMessage(chat, "Help message")
    }
    _, err := bot.Send(toSend)
    // todo: add retrying
    if err != nil {
        ErrorLog.Printf("Failed to send start answer %v\n", err)
        return
    }

    InfoLog.Printf("Sent Start message to %v (%v)\n", user.UserName, user.ID)
}
