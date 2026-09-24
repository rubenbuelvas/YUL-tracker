package clients

import (
	"context"
	"fmt"
	"log"
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

func (ys *YulScrapper) GetNextArrival() string {
	url := viper.GetString("YUL_ARRIVALS_URL")
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		//chromedp.Flag("disable-gpu", true),
		//chromedp.Flag("no-sandbox", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlSnapshot string
	var flightTexts []string
	jsGetShadowHTML := `(() => {
    try {
			const xpath = "//c-osf-flight-listing-line";
			const snapshot = document.evaluate(xpath, document, null, XPathResult.ORDERED_NODE_SNAPSHOT_TYPE, null);
			const results = [];
			for (let i = 0; i < snapshot.snapshotLength; i++) {
				const text = snapshot.snapshotItem(i).innerText.trim();
				if (text) {
					results.push(text);
				}
			}
			return results;
		} catch (err) {
			return []
		}
	})()`

	// jsGetShadowHTML := `(() => {
	// try {
	// 		const host = document.querySelector('c-osf-flights-listings');
	// 		if (host && host.shadowRoot) {
	// 			return host.shadowRoot.innerHTML || "";
	// 		}
	// 		if (host) {
	// 			return host.innerHTML || "";
	// 		}
	// 		return "ELEMENT_NOT_FOUND";
	// 	} catch (err) {
	// 		return "JS_ERROR: " + err.message;
	// 	}
	// })()`

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		// 2. Wait for the accept cookies button to become visible
		chromedp.WaitVisible("#didomi-notice-agree-button", chromedp.ByID),

		// 3. Click the accept cookies button
		chromedp.Click("#didomi-notice-agree-button", chromedp.ByID),
		chromedp.Sleep(1*time.Second),

		// Capture full static HTML snapshot of the component
		chromedp.Evaluate(jsGetShadowHTML, &flightTexts),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Parse static snapshot locally using Goquery (completely immune to browser DOM updates)
	// _, err := goquery.NewDocumentFromReader(strings.NewReader(htmlSnapshot))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(htmlSnapshot)
	// doc.Find("div.today").Each(func(i int, s *goquery.Selection) {
	// 	text := strings.TrimSpace(s.Text())
	// 	if text != "" {
	// 		flightTexts = append(flightTexts, text)
	// 	}
	// })

	fmt.Printf("Extracted %d flights from static HTML snapshot!\n", len(flightTexts))
	return strings.Join(flightTexts, "\n")
}

func (ys *YulScrapper) GetNextDeparture() string {
	return ""
}

func (ys *YulScrapper) GetFlightById() {
	// Implement this method to get flight data by ID
}
