package handlers

import (
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	db "github.com/tjones879/fake/database"
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
		file.SaveFile()
		db.SavePage(page, file)
	} else {
		sp := db.GetSavedPage(page)
		file := sp.Versions[0]
		fmt.Println(file)
		file.LoadFile()
		scraped = file.Contents
	}
	contents := template.HTML(scraped)

	c.HTML(200, "article.tmpl", gin.H{
		"contents": contents,
	})
}
