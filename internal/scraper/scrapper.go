package scraper

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gocolly/colly"
)

func Scrape() {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(t)

	pages := []string{}

	c.OnHTML("#search div:has(span) div[role=\"list\"]", func(e *colly.HTMLElement) {
		html, err := e.DOM.Html()
		if err != nil {
			fmt.Println(err)
			return
		}
		pages = append(pages, html)
		// fmt.Println(html)
	})

	c.Visit("file://" + dir + "/files/van-gogh-paintings.html")
	c.Wait()
	// for i, p := range pages {
	// 	fmt.Printf("%d : %s\n", i, p)
	// }
	fmt.Println(len(pages))
	// fmt.Println(pages[1])
}
