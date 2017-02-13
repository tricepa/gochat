package controllers

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/tricepa/go-react-chat/models"
  "github.com/olahol/melody"
  "database/sql"
)

func PostUserLogIn(c *gin.Context) {
	username := c.PostForm("username")
	id, err := models.GetIdFromName(username)
  
	//if username does not exist, save into db
	switch {
    case err == sql.ErrNoRows:
        id = models.InsertUser(username)
    case err != nil:
        fmt.Println(err)
    }

  c.JSON(200, gin.H {
    "user": models.User{ id, username },
  })
}

func PostUserLogOut(c *gin.Context) {
  username := c.PostForm("username")
  _ = models.UpdateLastSeen(username)
}

func GetWebsocket(c *gin.Context) {
  var m = melody.New()
  m = c.MustGet("melody").(*melody.Melody)
  m.HandleRequest(c.Writer, c.Request)
}