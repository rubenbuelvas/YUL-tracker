package clients

import (
	"context"
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

	var flightTexts string
	//preferNotToAnswerButton := "//button[contains(normalize-space(.), 'I prefer not to answer')]"

	xpathSelector := `//c-osf-flights-listings//div[contains(@class, 'today')]`
	removeElementJS := `() => {
		const el = document.querySelector('div[class*="QSIWebResponsiveDialog"]');
		if (el) {
			el.remove();
			return true;
		}
		return false;
	}`

	err := chromedp.Run(ctx,
		// 1. Navigate to page
		chromedp.Navigate(url),

		// 2. Wait for the accept cookies button to become visible
		chromedp.WaitVisible("#didomi-notice-agree-button", chromedp.ByID),

		// 3. Click the accept cookies button
		chromedp.Click("#didomi-notice-agree-button", chromedp.ByID),
		chromedp.Sleep(1*time.Second),

		// 2. Wait for the prefer not to answer button to become visible
		chromedp.Evaluate(removeElementJS, nil),

		//chromedp.Sleep(5*time.Second), // Optional: brief sleep to ensure button is interactable

		// 3. Click the prefer not to answer button
		//chromedp.Click(preferNotToAnswerButton, chromedp.ByQuery),

		// 4. Wait for the main content container to load after cookies are accepted
		chromedp.WaitVisible(xpathSelector, chromedp.BySearch),

		// 5. Brief sleep to allow LWC components to finish rendering flight rows
		//chromedp.Sleep(200*time.Second),

		// 6. Extract flight text data
		chromedp.Text(xpathSelector, &flightTexts, chromedp.BySearch),
		//chromedp.Sleep(1000*time.Second), // Optional: brief sleep to ensure data is captured
	)
	if err != nil {
		return err.Error()
	}
	return flightTexts
}

func (ys *YulScrapper) GetNextDeparture() string {
	return ""
}

func (ys *YulScrapper) GetFlightById() {
	// Implement this method to get flight data by ID
}
