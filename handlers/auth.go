package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tjones879/fake/database"
	"github.com/tjones879/fake/structs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	conf = &oauth2.Config{
		ClientID:     "200798239307-b5c40lataekifnhc470cfrt088e3bd3t.apps.googleusercontent.com",
		ClientSecret: "zfvqDPEoEFndtzsL0sugt-kw",
		RedirectURL:  "http://127.0.0.1:8080/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
)

// RandToken generates a random token of @len bytes
func RandToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

// LoginHandler handles /login
func LoginHandler(c *gin.Context) {
	state := RandToken(32)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	//c.Writer.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + getLoginURL(state) + "'><button>Login with Google!</button> </a> </body></html>"))
	c.Redirect(http.StatusSeeOther, getLoginURL(state))
}

func handleErr(c *gin.Context, err error) {
	log.Println(err)
	c.AbortWithError(http.StatusBadRequest, err)
	return
}

// AuthHandler handles the callback from oauth2
func AuthHandler(c *gin.Context) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	queryState := c.Request.URL.Query().Get("state")
	if retrievedState != queryState {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
		return
	}

	token, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		handleErr(c, err)
		return
	}

	client := conf.Client(oauth2.NoContext, token)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		handleErr(c, err)
		return
	}
	defer email.Body.Close()

	data, err := ioutil.ReadAll(email.Body)
	if err != nil {
		handleErr(c, err)
	}

	user := structs.User{}
	if err = json.Unmarshal(data, &user); err != nil {
		handleErr(c, err)
		return
	}

	storedUser, err := database.GetUserByID(user.ID)
	if err != nil {
		err = database.InsertUser(user)
	} else {
		user = storedUser
	}

	session.Set("user-id", user.ID)
	err = session.Save()
	if err != nil {
		handleErr(c, err)
	}

	c.Redirect(http.StatusFound, "/")
}
