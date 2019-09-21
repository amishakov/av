package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	Laptops string = "noutbuki"
	Phones  string = "telefony"
	PCs     string = "nastolnye_kompyutery"
)

func ParseAd(cat string, que string) {
	// Request the HTML page.
	res, err := http.Get(NewQueue(cat, que))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the ads
	doc.Find(".layout-internal .l-content .catalog-content .catalog-main .item .item_table-header").Each(func(i int, s *goquery.Selection) {
		// For each ad found, get the title, link and the price
		link, _ := s.Find("a").Attr("href")
		name := s.Find("h3").Text()
		price := s.Find(".price").Text()
		fmt.Printf("%s %s %s\n", name, link, price)
	})
}

func NewQueue(cat string, que string) string { // Assuming we want to search in Moscow city.
	return "https://www.avito.ru/moskva/" + url.QueryEscape(cat) + "?cd=1&q=" + strings.Replace(url.QueryEscape(que), " ", "+", -1)
}

func main() {
	ParseAd(Phones, "iPhone 4 16gb")
}
