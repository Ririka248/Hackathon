package models

import (
	"backend/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameHeight  = "heights"
	tableNameWeight  = "weights"
	tableNameSession = "sessions"
	tableNameFood    = "foods"
	tableNameSports  = "sports"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		password STRING,
		maxscore INTEGER,
		item1 INTEGER,
		item2 INTEGER,
		item3 INTEGER,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

	cmdH := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		heightvalue INTEGER,
		user_id INTEGER,
		date DATETIME,
		created_at DATETIME)`, tableNameHeight)

	Db.Exec(cmdH)

	cmdW := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		weightvalue FLOAT,
		user_id INTEGER,
		date DATETIME,
		created_at DATETIME)`, tableNameWeight)

	Db.Exec(cmdW)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)

	cmdF := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		category INTEGER,	
		id INTEGER,
		name STRING,
		protein STRING,
		fat STRING,
		carbohydrate STRING)`, tableNameFood)

	Db.Exec(cmdF)

	cmdM := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,	
		mets INTEGER,
		sports_name STRING)`, tableNameSports)

	Db.Exec(cmdM)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
