/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		fmt.Println("followers called")
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
}
