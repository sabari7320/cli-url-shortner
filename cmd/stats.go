// cmd/stats.go
package cmd

import (
	"fmt"

	"github.com/sabari7320/cli-url-shortner/internal/db"
	"github.com/sabari7320/cli-url-shortner/internal/models"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats [short-code]",
	Short: "Show statistics for a short URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db.Init()

		var url models.URL
		result := db.DB.Where("short_code = ?", args[0]).First(&url)

		if result.Error != nil {
			fmt.Println("Short URL not found:", args[0])
			return
		}

		fmt.Printf("Short URL: https://s.ly/%s\n", url.ShortCode)
		fmt.Printf("Long URL : %s\n", url.LongUrl)
		fmt.Printf("Clicks   : %d\n", url.Clicks)

		if url.FirstClick != nil {
			fmt.Printf("First    : %s\n", url.FirstClick.Format("2006-01-02 15:04"))
		}
		if url.LastClick != nil {
			fmt.Printf("Last     : %s\n", url.LastClick.Format("2006-01-02 15:04"))
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
