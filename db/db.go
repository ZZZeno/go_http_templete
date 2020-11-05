package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-gin-boilerplate/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var dbRaw *sql.DB
var db *gorm.DB

func Init() {
	c := config.GetConfig()
	var errRaw error
	var err error
	dbRaw, errRaw = sql.Open("mysql", c.GetString("db.uri"))
	if errRaw != nil {
		panic(errRaw)
	}
	db, err = gorm.Open(mysql.New(mysql.Config{Conn: dbRaw}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	dbRaw.SetConnMaxLifetime(time.Minute * 3)
	dbRaw.SetMaxOpenConns(1024)
	dbRaw.SetMaxIdleConns(1024)
}

func GetDBRaw() *sql.DB{
	return dbRaw
}

func GetDB() *gorm.DB{
	return db
}

func CloseDb(){
	dbRaw.Close()
}
