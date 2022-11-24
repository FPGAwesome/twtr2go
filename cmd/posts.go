/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	cdptypes "github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

type TweetData struct {
	Date string
	TweetContent string
}

// postsCmd represents the posts command
var postsCmd = &cobra.Command{
	Use:   "posts",
	Short: "The default twitter scraper commands",
	Long: `Run this command to scrape the posts of a specific twitter user.
		   Format: twtr2go posts <target username>
		   
		   Optional flags: `,
	Run: func(cmd *cobra.Command, args []string) {
		
		if len(args) < 1 {
			fmt.Println("Error: You must run this on a username.\n",
			"Format: twtr2go posts <target username>")

			return
		}


		//Create chrome driver contexts
		ctx, cancel := chromedp.NewContext(
				context.Background(),
				chromedp.WithLogf(log.Printf),)
    	defer cancel()

		// Get a nice, big window going to get more tweets
		// in our view
		chromedp.WindowSize(1920,1080)
		chromedp.Run(ctx, chromedp.ResetViewport(),)

		search_url := "https://mobile.twitter.com/search?q=(from%3A"+args[0]+")&src=typed_query"
		emulation.SetUserAgentOverride("Hello Tweeter v0.01")

		// Setup our file creation stuff
		file, err := os.Create("export.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := chromedp.Run(ctx,
			chromedp.Navigate(search_url),
			chromedp.WaitVisible("div[data-testid=\"tweetText\"]"),
		); err != nil{
			log.Println(err)
		}
		
		var tweetCounter = 1
		var nodes []*cdptypes.Node
		var divNodes []*cdptypes.Node

		// tweet storage map
		var tweetMap map[string] int

		tweetMap = make(map[string]int)

		for {
			tweetCounter=1
			nodes=nil
			// Begin the scrapening
			if err := chromedp.Run(ctx,
				chromedp.WaitVisible(`div[data-testid="tweetText"]`),
				chromedp.Sleep(2*time.Second), // lazy wait for dom loading
				chromedp.Nodes(`article[data-testid="tweet"`, &divNodes),
				chromedp.Nodes(`div[data-testid="tweetText"] > span`, &nodes, chromedp.ByQueryAll, chromedp.NodeVisible),
				chromedp.ActionFunc(func(ctx context.Context) error {
					_, exp, err := runtime.Evaluate(`window.scrollTo(0,document.body.scrollHeight);`).Do(ctx)
					if err != nil {
						return err
					}
					if exp != nil {
						return exp
					}
					return nil
				}),
			); err != nil {

				log.Fatal(err)
			}

			for _, n := range nodes {
				if len(n.Children) < 1 {
					continue
				}
				//fmt.Println(tweetCounter, n.Children[0].NodeValue)
				tweetMap[n.Children[0].NodeValue]=1
				tweetCounter++
			}

			// for _,n := range divNodes {
			// 	// ungodly way to grab divs, rendering without the css
			// 	// is a nightmare
				
			// 	var html string
			// 	chromedp.Run(ctx,
			// 		chromedp.ActionFunc(func(ctx context.Context) error {
			// 			html, err = dom.GetOuterHTML().WithNodeID(n.NodeID).Do(ctx)
			// 			return err
			// 		}),
			// 	)
			// 	fmt.Println(html)
			// }

			var scrollStatus bool

			if err := chromedp.Run(ctx,
				chromedp.Sleep(2*time.Second), // lazy wait for dom loading
				chromedp.Evaluate(
					`
					function checkScroll(){ 
							return ((window.innerHeight + window.pageYOffset) >= document.body.scrollHeight ) 
						}

						checkScroll();`,
					&scrollStatus,
				)); err != nil {
					log.Fatal(err)
				}
			
				if scrollStatus {
					break
				}
		}

		for key,_ := range tweetMap {
			fmt.Println(key)
		}


	},
}

func init() {
	rootCmd.AddCommand(postsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
