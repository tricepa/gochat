package models

import (
	"time"
	"fmt"
)

type Message struct {
  Username string
  Content string
  ContentType string
}

func GetAllMessagesAscending() []string {
	var (
	  messageContent string
	  allContent []string
	)

	rows, _ := DB.Query("SELECT content FROM messages ORDER BY timestamp ASC")
	defer rows.Close()
	for rows.Next() {
	  err := rows.Scan(&messageContent)
	  if err != nil {
	    fmt.Println(err)
	  }
	  allContent = append(allContent, messageContent)
	}

	return allContent
}

func GetNewMessagesAscending(username string) []string {
	var (
	  messageContent string
	  newContent []string
	)

	//get messages timestamped after end of user's last session
	lastSeen, _ := GetLastSeen(username)
	rows, _ := DB.Query("SELECT content FROM messages WHERE timestamp > ? ORDER BY timestamp ASC", lastSeen)
	defer rows.Close()

	for rows.Next() {
	  err := rows.Scan(&messageContent)
	  if err != nil {
	    fmt.Println(err)
	  }
	  newContent = append(newContent, messageContent)
	}

	return newContent
}

func InsertMessage(username string, contentType string, content string) {
	insertStatement, _ := DB.Prepare("INSERT messages SET username = ?, message_type = ?, content = ?, timestamp = ?")
    _, err := insertStatement.Exec(username, contentType, content, time.Now())
    if err != nil {
      fmt.Println(err)
    }
}