package main

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "main/database"
)

func AddUser(user *tgbotapi.User) bool {
    if user == nil {
        return false
    }
    added, err := DB.AddUser(database.User{Id: user.ID})
    if err != nil {
        ErrorLog.Printf("Failed to add user: %v\n", err)
        return added
    }
    if added {
        InfoLog.Printf("Added user %v (%v)\n", user.UserName, user.ID)
    }
    return added
}
