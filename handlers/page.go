package handlers

import (
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	db "github.com/tjones879/fake/database"
	"github.com/tjones879/fake/structs"
	"github.com/tjones879/fake/util"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	p = createPolicy()
)

func createPolicy() *bluemonday.Policy {
	p := bluemonday.UGCPolicy()

	return p
}

func fixURL(base, href string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseURL.ResolveReference(uri)
	return uri.String()
}

func scrapePage(url string) []byte {
	var buffer []byte
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	b := resp.Body
	buffer, err = ioutil.ReadAll(b)
	b.Close()
	if err != nil {
		log.Fatal(err)
	}

	html := p.SanitizeBytes(buffer)
	return html
}

// PageHandler handles /page?u= requests.
func PageHandler(c *gin.Context) {
	page := c.DefaultQuery("u", "https://en.wikipedia.org/wiki/Dune_(novel)")
	var scraped string
	if !db.IsPageSaved(page) {
		scraped = string(scrapePage(page))
		file := util.CreateStorage(xxhash.Sum64String(scraped), page, scraped)
		util.SaveFile(&file)
		db.SavePage(page, file)
		db.SaveFile(file)
		if uid := getUserID(c); uid != "" {
			db.AddFileToUser(uid, structs.FileReference{
				ID:   file.UID,
				Name: file.Name,
			})
		}
	} else {
		sp := db.GetSavedPage(page)
		uid := sp.Versions[0]
		file, _ := db.GetFileByID(uid)
		util.LoadFile(&file)
		scraped = file.Contents
	}
	contents := template.HTML(scraped)

	c.HTML(200, "article.tmpl", gin.H{
		"contents": contents,
	})
}

// SavedHandler TODO
func SavedHandler(c *gin.Context) {
	page := c.DefaultQuery("id", "123456789")
	fmt.Println("SavedHandler", page)
	f, _ := db.GetFileByID(page)
	util.LoadFile(&f)
	contents := template.HTML(f.Contents)

	c.HTML(200, "article.tmpl", gin.H{
		"contents": contents,
	})
}
