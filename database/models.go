package database

import (
    "time"
)

type User struct {
    Id int64 `db:"id"`
}

type Record struct {
    Id      int64     `db:"id"`
    UserId  int64     `db:"user_id"`
    Name    string    `db:"name"`
    Comment string    `db:"comment"`
    Value   int64     `db:"value"`
    Time    time.Time `db:"time"`
}
