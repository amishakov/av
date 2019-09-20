package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

<<<<<<< HEAD
const (
	Laptops string = "noutbuki"
	Phones  string = "telefony"
	PCs     string = "nastolnye_kompyutery"
)

func ParseAd(cat string, que string) {
=======
func ParseAds() {
>>>>>>> 011cee934d36f3aab3284067932889e0d43568c3
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

<<<<<<< HEAD
	// Find the ads.
	doc.Find(".layout-internal .l-content .catalog-content .catalog-main .item .item_table-header").Each(func(i int, s *goquery.Selection) {
		// For each ad found, get the link, price and title.
=======
	// Find the ads
	doc.Find(".layout-internal .l-content .catalog-content .catalog-main .item .item_table-header").Each(func(i int, s *goquery.Selection) {
		// For each ad found, get the title, link and the price
>>>>>>> 011cee934d36f3aab3284067932889e0d43568c3
		link, _ := s.Find("a").Attr("href")
		name := s.Find("h3").Text()
		price := s.Find(".price").Text()
		fmt.Printf("%s %s %s\n", name, link, price)
	})
}

func NewQueue(cat string, que string) string { // Assuming we want to search in Moscow city.
	return "https://www.avito.ru/moskva/" + cat + "?cd=1&q=" + strings.Replace(que, " ", "+", -1)
}

func main() {
<<<<<<< HEAD
	ParseAd(Phones, "iPhone 4 16gb")
=======
	ParseAds()
>>>>>>> 011cee934d36f3aab3284067932889e0d43568c3
}
