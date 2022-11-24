package util

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

// We'll grab a login context with this if the user desires login
// (they have no choice for some fucntionality, which means handling
// errors)
func login(login string, pass string) (context.Context, context.CancelFunc) {
	//Create chrome driver contexts
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),)

	err := chromedp.Run(ctx,
			chromedp.Navigate("https://twitter.com/i/flow/login"),
			chromedp.WaitVisible(`input[autocomplete="username"]`),
			
	)


	// Make sure to defer() this in whatever calls the util.
	// If there is a login error, we'll just get a not logged-in
	// context
	return ctx, cancel
}