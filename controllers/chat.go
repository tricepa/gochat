package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/tricepa/gochat/models"
)

func GetAllMessages(c *gin.Context) {
	allMessages := models.GetAllMessagesAscending()
	c.JSON(200, gin.H{
	  "messages": allMessages,
	})
}

func GetNewMessages(c *gin.Context) {
	username := c.PostForm("username")
	newMessages := models.GetNewMessagesAscending(username)
	c.JSON(200, gin.H{
	  "messages": newMessages,
	})
}

