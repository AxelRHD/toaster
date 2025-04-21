package toaster

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nullism/bqb"
)

type DB struct {
	DB                  *sqlx.DB
	ToastTemplate       string
	HyperscriptTemplate string
}

const (
	sqlite_dt = "2006-01-02 15:04:05"
)

func ConnectDB(driverName string, dataSourceName string) (*DB, error) {
	var db *DB

	sdb, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return db, err
	}

	return &DB{
		DB:                  sdb,
		ToastTemplate:       toastTemplate,
		HyperscriptTemplate: jsTemplate,
	}, nil
}

func (db DB) CreateSchema() error {
	qry := `CREATE TABLE IF NOT EXISTS messages(
		id TEXT primary key,
		title TEXT DEFAULT '' NOT NULL,
		message TEXT DEFAULT '' NOT NULL,
		location TEXT DEFAULT 'top-right' NOT NULL,
		icon INTEGER DEFAULT 1 NOT NULL,
		dismissable INTEGER DEFAULT 0 NOT NULL,
		msg_type TEXT DEFAULT '' NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL
	)`

	_, err := db.DB.Exec(qry)
	return err
}

func (db *DB) Add(toast *Toast) (string, error) {
	id := generateUniqueID()

	cols := bqb.Optional("")
	cols.Comma("id")
	cols.Comma("title")
	cols.Comma("message")
	cols.Comma("location")
	cols.Comma("icon")
	cols.Comma("dismissable")
	cols.Comma("msg_type")
	cols.Comma("created_at")

	vals := bqb.Optional("")
	vals.Comma("?", id)
	vals.Comma("?", toast.Title)
	vals.Comma("?", toast.Message)
	vals.Comma("?", string(toast.Location))
	vals.Comma("?", toast.Icon)
	vals.Comma("?", toast.Dismissable)
	vals.Comma("?", string(toast.Type))

	if toast.CreatedAt.IsZero() {
		vals.Comma("?", time.Now().Format(sqlite_dt))
	}

	q := bqb.New("INSERT INTO messages (?) VALUES(?)", cols, vals)

	qry, err := q.ToRaw()
	if err != nil {
		return "", err
	}
	fmt.Println(qry)

	_, err = db.DB.Exec(qry)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *DB) Get(key string) (*Toast, error) {
	var toast Toast

	q := bqb.New("SELECT id,title,message,location,icon,dismissable,msg_type,created_at FROM messages WHERE id = ?", key)
	qry, err := q.ToRaw()
	if err != nil {
		return nil, err
	}

	err = db.DB.Get(&toast, qry)
	if err != nil {
		return nil, err
	}

	dt, err := time.Parse(sqlite_dt, toast.CreatedAtStr)
	if err == nil {
		toast.CreatedAt = dt
	}

	toast.ToastTemplate = db.ToastTemplate
	toast.JsTemplate = db.HyperscriptTemplate

	return &toast, nil
}

func (db *DB) Delete(key string) error {
	q := bqb.New("DELETE FROM messages WHERE id = ?", key)

	qry, err := q.ToRaw()
	if err != nil {
		return err
	}

	fmt.Println(qry)

	_, err = db.DB.Exec(qry)

	return err
}
