package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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

func scrapePage(url string) string {
	var buffer bytes.Buffer
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

	depth := 0
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return buffer.String()
		case html.TextToken:
			if depth > 0 {
				text := z.Raw()
				buffer.WriteString(string(text[:]))
			}
		case html.StartTagToken:
			raw := z.Raw()
			tn, attr := z.TagName()
			if tn[0] == 'p' {
				depth++
				buffer.WriteString(string(raw[:]))
			} else if tn[0] == 'a' && attr && depth > 0 {
				for {
					key, val, more := z.TagAttr()
					if string(key[:]) == "href" {
						buffer.WriteString("<a " + string(key[:]) + "=" + fixURL(url, string(val[:])) + "</>")
					}
					if !more {
						break
					}
				}

			}
		case html.EndTagToken:
			raw := z.Raw()
			tn, _ := z.TagName()
			if depth > 0 {
				buffer.WriteString(string(raw[:]))
			}
			if tn[0] == 'p' {
				depth--
				if depth <= 0 {
					buffer.WriteString("")
				}
			}
		}
	}
}

// PageHandler handles /page?u= requests.
func PageHandler(c *gin.Context) {
	page := c.DefaultQuery("u", "https://en.wikipedia.org/wiki/Dune_(novel)")
	t, c := scrapePage(page)
	fmt.Println(c)
	title := template.HTML(t)
	contents := template.HTML(c)

	c.HTML(200, "article.tmpl", gin.H{
		"title":    title,
		"contents": contents,
	})
}
