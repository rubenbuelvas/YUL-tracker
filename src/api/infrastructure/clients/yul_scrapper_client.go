package clients

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/spf13/viper"
)

type YulScrapper struct {
}

func NewYulScrapper() *YulScrapper {
	return &YulScrapper{}
}

func (ys *YulScrapper) GetArrivals() ([]string, error) {
	url := viper.GetString("YUL_ARRIVALS_URL")
	return ys.scrape(url)
}

func (ys *YulScrapper) GetDepartures() ([]string, error) {
	url := viper.GetString("YUL_DEPARTURES_URL")
	return ys.scrape(url)
}

func (ys *YulScrapper) scrape(url string) ([]string, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		// 	chromedp.Flag("disable-gpu", true),
		// 	chromedp.Flag("no-sandbox", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var flights []string
	getFlightDataJS := `(() => {
    try {
			const xpath = "//c-osf-flight-listing-line";
			const snapshot = document.evaluate(xpath, document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null);
			const results = [];
			for (let i = 0; i < snapshot.snapshotLength; i++) {
				const text = snapshot.snapshotItem(i).innerText;
				if (text) {
					results.push(text);
				}
			}
			return results;
		} catch (err) {
			return []
		}
	})()`

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.WaitVisible("#didomi-notice-agree-button", chromedp.ByID),

		// Deal with cookies popup
		chromedp.Click("#didomi-notice-agree-button", chromedp.ByID),
		chromedp.Sleep(1*time.Second),

		// Get data
		chromedp.Evaluate(getFlightDataJS, &flights),
	)
	fmt.Println(strings.Split(flights[0], "\n"))
	return flights, err
}
