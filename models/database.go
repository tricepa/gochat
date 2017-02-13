package models

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dsn string) {
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println(err)
    }
    //check connection on startup
    if err = DB.Ping(); err != nil {
        fmt.Println(err)
    }
}
