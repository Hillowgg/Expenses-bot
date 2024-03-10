package main

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "main/database"
)

func AddUser(user *tgbotapi.User) {
    if user == nil {
        return
    }
    err := DB.AddUser(database.User{Id: user.ID})
    if err != nil {
        ErrorLog.Printf("Failed to add user %v\n", err)
        return
    }
    InfoLog.Printf("Added user %v (%v)\n", user.UserName, user.ID)
}
