package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"github.com/tricepa/gochat/controllers"
	"github.com/tricepa/gochat/models"
)

func main() {
	r := gin.New()
	m := melody.New()	
	r.Use(setMelodyToContext(m))

	dsn := os.Getenv("CHAT_DB_DSN")
	models.InitDB(dsn)

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./client/index.html")
	})
	r.POST("/login", controllers.PostUserLogIn)
	r.GET("/allMessages", controllers.GetAllMessages)
	r.POST("/newMessages", controllers.GetNewMessages)
	r.POST("/logout", controllers.PostUserLogOut)
	r.GET("/ws", controllers.GetWebsocket)

	m.HandleMessage(func(s *melody.Session, message []byte) {
	    var msg models.Message
	    err := json.Unmarshal(message, &msg)
	    if err != nil {
	      fmt.Println(err)
	    }
	    models.InsertMessage(msg.Username, "text", msg.Content)
	    m.Broadcast([]byte(msg.Content))
  	})

	r.Static("/assets/js", "./build") 
	r.Run(":5000")
}

// middleware to set Melody instance to Gin context for use in session controller 
func setMelodyToContext(m *melody.Melody) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("melody", m)
    }
}
