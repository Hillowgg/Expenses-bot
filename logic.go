package main

import (
    "errors"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "main/database"
)

func AddUser(user *tgbotapi.User) error {
    if user == nil {
        return errors.New("nil user")
    }
    s, err := DB.AddUser(database.User{Id: user.ID})
    if err != nil {
        ErrorLog.Printf("Failed to add user: %v\n", err)
        return err
    }
    if s {
        InfoLog.Printf("Added user %v (%v)\n", user.UserName, user.ID)
    }
    return nil
}
