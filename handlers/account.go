package handlers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AccountHandler handles /me
func AccountHandler(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("user-id")

	if id == nil {
		fmt.Println("Sorry, you must sign in.")
	} else {
		fmt.Println("user-id:", session.Get("user-id"))
	}
}
