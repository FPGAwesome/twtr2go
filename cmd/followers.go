/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"twtr2go/util"

	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

// followersCmd represents the followers command
var followersCmd = &cobra.Command{
	Use:   "followers",
	Short: "Scrape a given users followers",
	Long: `Run this command to scrape the followers of a specific twitter user.
	Format: twtr2go followers <target username>
	
	Optional flags: `,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 || !strings.Contains( cmd.Flag("login").Value.String(), ":" ){
			fmt.Println("Error: You must run this on a username and specify login information.\n",
			"Format: twtr2go posts --login <user>:<pass> <target username>")

			return
		}


		loginDetails := strings.Split(cmd.Flag("login").Value.String(),":")

		_, cancel := util.Login(loginDetails[0],loginDetails[1])
		defer cancel()

		var html string

		chromedp.OuterHTML(`body`, &html)

		fmt.Println(html)

	},
}

func init() {
	rootCmd.AddCommand(followersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// followersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// followersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	followersCmd.Flags().StringP("login", "l", "user:pass", "Specify a user:pass pair for authentication.")
}
