package util

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

// We'll grab a login context with this if the user desires login
// (they have no choice for some fucntionality, which means handling
// errors)
func Login(login string, pass string) (context.Context, context.CancelFunc) {
	//Create chrome driver contexts
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),)

	//ctx, cancel = context.WithTimeout(ctx, 25*time.Second)

	err := chromedp.Run(ctx,
			chromedp.Navigate("https://twitter.com/i/flow/login"),
			//chromedp.WaitVisible(`input[autocomplete="username"]`),
			chromedp.SetValue(`input[autocomplete="username"]`,login),
			chromedp.SendKeys(`input[autocomplete="username"]`,kb.Enter),
			chromedp.Sleep(2*time.Second),
			//chromedp.WaitVisible(`#input[autocomplete="current-password"]`),
			chromedp.SetValue(`input[autocomplete="current-password"]`,pass),
			chromedp.SendKeys(`input[autocomplete="current-password"]`,kb.Enter),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Make sure to defer() this in whatever calls the util.
	// If there is a login error, we'll just get a not logged-in
	// context
	return ctx, cancel
}