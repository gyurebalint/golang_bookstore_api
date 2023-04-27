package fetch

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gyurebalint/golang_bookstore_api/pkg/models"
)

func ScrapeBookFromAmazon(link string) models.Book {
	var book models.Book

	fmt.Println("Started scraping...")
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/112.0")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})

	//TITLE
	c.OnHTML("#productTitle", func(h *colly.HTMLElement) {
		book = models.Book{
			Title: h.Text,
		}
	})

	var authors []string
	//AUTHOR
	c.OnHTML("#bylineInfo", func(h *colly.HTMLElement) {
		h.ForEach(".author.notFaded", func(_ int, element *colly.HTMLElement) {
			authors = append(authors, element.ChildText("a"))
		})
	})

	//DESCRIPTION
	c.OnHTML("#bookDescription_feature_div", func(h *colly.HTMLElement) {
		bla := h.ChildText("div p")

		h.ForEach("li", func(_ int, elem *colly.HTMLElement) {
			if strings.Contains(elem.Text, ":") {
				bla += "\n" + elem.Text
			} else {
				bla += "\n" + "-" + elem.Text
			}
		})

		book.Description = bla
	})

	//COVER IMAGE URL
	c.OnHTML("#imgBlkFront", func(h *colly.HTMLElement) {
		//size: 260x360
		bla := h.Attr("src")
		book.CoverImageUrl = bla
	})

	c.Visit(link)
	book.Authors = strings.Join(authors[:], ",")
	book.Link = link

	fmt.Printf("Title: %s\nDescription: %s \nImageURL: %s", book.Title, book.Description, book.CoverImageUrl)
	fmt.Println("\nFinished scraping...")

	return book
}
