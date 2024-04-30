package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gocolly/colly"
	"log"
)

type IkmanAd struct {
	title, price, link, image string
}

func getAds() {
	var ikmanAds []IkmanAd

	c := colly.NewCollector()

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
		var price = e.ChildText(".price--3SnqI")
		var ad = IkmanAd{
			title: title,
			link:  link,
			image: image,
			price: price,
		}
		ikmanAds = append(ikmanAds, ad)
	})
	print(ikmanAds)

	err := c.Visit("https://ikman.lk/en/ads/colombo/house-rentals?money.price.maximum=30000")
	if err != nil {
		return
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Hello World",
		StatusCode: 200,
	}, nil
}

func main() {
	//lambda.Start(handler)
	getAds()
}
