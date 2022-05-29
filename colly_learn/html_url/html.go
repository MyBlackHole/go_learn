package main

import (
	"fmt"
	"net/http"
	// "net/url"
	"os"
	"path/filepath"

	"github.com/gocolly/colly/v2"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(t)

	pages := []string{}

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		pages = append(pages, e.Text)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
		// c.Visit("http://www.baidu.com")
		c.Visit(e.Request.AbsoluteURL(dir + link))
		// c.Visit("file://" + dir + "/html" + e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting --- ", r.URL.String())
	})

	fmt.Println("file://" + dir + "/html/index.html")
	c.Visit("file://" + dir + "/html/index.html")
	c.Wait()
	for i, p := range pages {
		fmt.Printf("%d : %s\n", i, p)
	}
}
