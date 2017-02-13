package models

import (
	"fmt"
	"time"
)

type User struct {
	Id int
	Username string
}

func GetIdFromName(username string) (int, error) {
	var id int
	err := DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	return id, err
}

func InsertUser(username string) int {
	stmt, err := DB.Prepare("INSERT users SET username = ?")
    if err != nil {
		fmt.Println(err)
	}
    _, err = stmt.Exec(username)
    id, _ := GetIdFromName(username)
    return id
}

func UpdateLastSeen(username string) error {
	stmt, err := DB.Prepare("UPDATE users SET last_seen = ? WHERE username = ?")
	if err != nil {
		fmt.Println(err)
	}
    _, err = stmt.Exec(time.Now(), username)
    return err
}

func GetLastSeen(username string) (time.Time, error) {
	var lastSeen time.Time
	err := DB.QueryRow("SELECT last_seen FROM users WHERE username = ?", username).Scan(&lastSeen)
	if err != nil {
		fmt.Println(err)
	}
    return lastSeen, err
}

