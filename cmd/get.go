package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/sabari7320/cli-url-shortner/internal/db"
	"github.com/sabari7320/cli-url-shortner/internal/server"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [code]",
	Short: "Open short URL in browser",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db.Init()
		go server.Start()                  // start server in background
		time.Sleep(500 * time.Millisecond) // tiny wait

		shortURL := "http://localhost:8080/" + args[0]

		var err error
		switch runtime.GOOS {
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", shortURL).Start()
		case "darwin":
			err = exec.Command("open", shortURL).Start()
		default:
			err = exec.Command("xdg-open", shortURL).Start()

		}
		if err != nil {
			fmt.Println("Error opening browser:", err)
		} else {
			fmt.Println("Opening:", shortURL)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
