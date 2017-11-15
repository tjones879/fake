package handlers

import (
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
