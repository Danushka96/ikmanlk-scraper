package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"time"
)

type IkmanAd struct {
	ID        string    `bson:"_id, omitempty"`
	Title     string    `bson:"title"`
	Price     string    `bson:"price"`
	Link      string    `bson:"link"`
	Image     string    `bson:"image"`
	CreatedAt time.Time `bson:"created_at"`
}

func getAds() {
	var ikmanAds []IkmanAd

	c := colly.NewCollector()
	c.SetRequestTimeout(10 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	c.OnHTML("li.gtm-normal-ad", func(e *colly.HTMLElement) {
		var link = e.ChildAttr("a", "href")
		var title = e.ChildAttr("a", "title")
		var image = e.ChildAttr("img", "src")
		image = strings.Replace(image, "142", "800", 1)
		image = strings.Replace(image, "107", "500", 1)
		var price = e.ChildText(".price--3SnqI")
		var ad = IkmanAd{
			ID:        link,
			Title:     title,
			Link:      "https://ikman.lk" + link,
			Image:     image,
			Price:     price,
			CreatedAt: time.Now(),
		}
		if !ExistAd(ad) {
			SendMessage(ad)
			_, err := SaveAd(ad)
			if err != nil {
				return
			}
		}
	})
	print(ikmanAds)

	err := c.Visit("https://ikman.lk/en/ads/colombo/house-rentals?money.price.maximum=30000")
	if err != nil {
		return
	}
}

func handler(_ events.CloudWatchEvent) error {
	getAds()
	return nil
}

func main() {
	lambda.Start(handler)
	//getAds()
}
