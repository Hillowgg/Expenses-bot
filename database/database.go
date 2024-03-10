package database

import (
    "context"
    "errors"
    "time"

    "github.com/jackc/pgx/v5"
)

type MyDB struct {
    Db *pgx.Conn
}

func NewDB(dsn string) (*MyDB, error) {
    pool, err := pgx.Connect(context.Background(), dsn)
    if err != nil {
        // FatalLog.Fatalf("Failed to connect to db: %v\n", err)
        return nil, err
    }
    return &MyDB{pool}, nil
}

func (d *MyDB) GetUser(userId int64) (User, error) {
    row := d.Db.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", userId)
    var id int64
    err := row.Scan(&id)
    if err != nil {
        return User{}, nil
    }
    return User{id}, nil
}

func (d *MyDB) AddUser(user User) (bool, error) {
    row := d.Db.QueryRow(context.Background(), "INSERT INTO users VALUES ($1) ON CONFLICT (id) DO NOTHING RETURNING true", user.Id)
    var res bool
    err := row.Scan(&res)
    if !errors.Is(err, pgx.ErrNoRows) {
        return res, err
    }
    return res, nil
}

func (d *MyDB) GetRecord(recordId int64) (Record, error) {
    row := d.Db.QueryRow(context.Background(), "SELECT * FROM records WHERE id=$1", recordId)

    if err := row.Scan(); err != nil {
        return Record{}, err
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
    rows, err := d.Db.Query(context.Background(), "SELECT * FROM records WHERE user_id=$1", userId)
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
        context.Background(),
        "INSERT INTO records (user_id, name, comment, value, time) VALUES ($1, $2, $3, $4, $5)",
        rec.UserId,
        rec.Name,
        rec.Comment,
        rec.Value,
        rec.Time.Unix(),
    )
    if err := row.Scan(); err != nil {
        return err
    }
    return nil
}
