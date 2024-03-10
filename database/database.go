package database

import (
    "database/sql"
    "time"
)

type MyDB struct {
    Db *sql.DB
}

func NewDB(driver string, dsn string) (*MyDB, error) {
    pool, err := sql.Open(driver, dsn)
    if err != nil {
        // FatalLog.Fatalf("Failed to connect to db: %v\n", err)
        return nil, err
    }
    return &MyDB{pool}, nil
}

func (d *MyDB) GetUser(userId int64) (User, error) {
    row := d.Db.QueryRow("SELECT * FROM users WHERE id=?", userId)
    if row.Err() != nil {
        return User{}, row.Err()
    }
    var id int64
    err := row.Scan(&id)
    if err != nil {
        return User{}, nil
    }
    return User{id}, nil
}

func (d *MyDB) AddUser(user User) error {
    row := d.Db.QueryRow("INSERT INTO users VALUES (?)", user.Id)
    if row.Err() != nil {
        return row.Err()
    }
    return nil
}

func (d *MyDB) GetRecord(recordId int64) (Record, error) {
    row := d.Db.QueryRow("SELECT * FROM records WHERE id=?", recordId)

    if row.Err() != nil {
        return Record{}, row.Err()
    }
    var rec Record
    var t int64
    err := row.Scan(&rec.Id, &rec.UserId, &rec.Name, &rec.Comment, &rec.Value, &t)
    rec.Time = time.Unix(t, 0)
    if err != nil {
        return Record{}, err
    }
    return rec, nil
}

func (d *MyDB) GetRecordsByUserId(userId int64) ([]Record, error) {
    rows, err := d.Db.Query("SELECT * FROM records WHERE user_id=?", userId)
    if err != nil {
        return nil, err
    }
    var rec Record
    var t int64
    recs := make([]Record, 0)
    for rows.Next() {
        err = rows.Scan(&rec.Id, &rec.UserId, &rec.Name, &rec.Comment, &rec.Value, &t)
        recs = append(recs, rec)
        if err != nil {
            return nil, err
        }
    }
    return recs, nil
}

func (d *MyDB) AddRecord(rec Record) error {
    row := d.Db.QueryRow(
        "INSERT INTO records (user_id, name, comment, value, time) VALUES (?, ?, ?, ?, ?)",
        rec.UserId,
        rec.Name,
        rec.Comment,
        rec.Value,
        rec.Time.Unix(),
    )
    if row.Err() != nil {
        return row.Err()
    }
    return nil
}
