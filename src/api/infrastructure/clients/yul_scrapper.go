package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/viper"
)

type YulScrapper struct {
}

func NewYulScrapper() *YulScrapper {
	return &YulScrapper{}
}

func (ys *YulScrapper) GetNextArrival() string {
	url := viper.GetString("YUL_ARRIVALS_URL")
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	results := doc.Find("div").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println("Flight:", text)
		log.Println(s.Text())
	})
	return results.Text()
}

func (ys *YulScrapper) GetNextDeparture() string {
	return ""
}

func (ys *YulScrapper) GetFlightById() {
	// Implement this method to get flight data by ID
}
