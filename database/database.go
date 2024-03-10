package database

import (
    "database/sql"
    "time"
)

type MyDB struct {
    db sql.DB
}

func (d *MyDB) GetUser(userId int64) (User, error) {
    row := d.db.QueryRow("SELECT * FROM users WHERE id=?", userId)
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

func (d *MyDB) GetRecord(recordId int64) (Record, error) {
    row := d.db.QueryRow("SELECT * FROM records WHERE id=?", recordId)

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
    rows, err := d.db.Query("SELECT * FROM records WHERE user_id=?", userId)
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
