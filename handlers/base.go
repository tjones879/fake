package handlers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexHandler handles the / route.
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

// AboutHandler handles /about route.
func AboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", gin.H{})
}

// LogoutHandler clears a user's session and redirects to index
func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	c.Redirect(http.StatusFound, "/")
}
