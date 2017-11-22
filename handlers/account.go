package handlers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	db "github.com/tjones879/fake/database"
	"net/http"
)

// AccountHandler handles /me
func AccountHandler(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("user-id")

	if id == nil {
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		// Get all pages annotated by the user
		user, _ := db.GetUserByID(id.(string))
		fmt.Println("Files", db.GetUserPages(user))
		// Allow deletion of old files
		// Allow signing out
		// Allow deleting account
		c.JSON(http.StatusOK, user)
	}
}
