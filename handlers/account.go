package handlers

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	db "github.com/tjones879/fake/database"
	"github.com/tjones879/fake/structs"
	"net/http"
)

// AccountHandler handles /me
func AccountHandler(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("user-id")
	var files []structs.FileReference

	if id == nil {
		//c.AbortWithStatus(http.StatusForbidden)
		files = []structs.FileReference{
			structs.FileReference{
				ID:   "123456789",
				Name: "HELLO",
			},
			structs.FileReference{
				ID:   "987654321",
				Name: "GOODBYE",
			},
		}
	} else {
		// Get all pages annotated by the user
		user, _ := db.GetUserByID(id.(string))
		fmt.Println("Files", user.Files)
		files = user.Files
		// Allow deletion of old files
		// Allow signing out
		// Allow deleting account
	}
	c.HTML(http.StatusOK, "me.tmpl", gin.H{
		"files": files,
	})
}
