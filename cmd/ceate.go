package cmd

import (
	"fmt"

	"github.com/sabari7320/cli-url-shortner/internal/db"
	"github.com/sabari7320/cli-url-shortner/internal/shortener"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [url]",
	Short: "Create a short URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db.Init()
		short, err := shortener.Create(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Short URL: %s\n", short)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
