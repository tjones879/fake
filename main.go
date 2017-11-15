package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tjones879/fake/handlers"
)

func main() {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte(handlers.RandToken(64)))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 3600 * 24 * 7,
	})

	router.Use(sessions.Sessions("fake", store))
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.Static("/js", "./static/js")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.IndexHandler)
	router.GET("/about", handlers.AboutHandler)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/auth", handlers.AuthHandler)
	router.GET("/page", handlers.PageHandler)
	router.GET("/me", handlers.AccountHandler)

	storage := router.Group("/store")
	storage.GET("/", handlers.RootAnnotate)
	storage.GET("/annotations", handlers.IndexAnnotate)
	storage.POST("/annotations", handlers.CreateAnnotate)
	storage.GET("/annotations/:id", handlers.ReadAnnotate)
	storage.POST("/annotations/:id", handlers.UpdateAnnotate)
	storage.DELETE("/annotations/:id", handlers.DeleteAnnotate)

	router.Run("127.0.0.1:8080")

}
