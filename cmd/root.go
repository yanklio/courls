package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	scrapper "github.com/yanklio/courls/scraper"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "courls",
	Short: "Scrape and count URLs from a web domain",
	Long:  "A web scraper that visits a domain and counts all unique URLs found, with configurable visit limits and file output",
	Run: func(cmd *cobra.Command, args []string) {
		url := getUrl(args)
		limit, _ := cmd.Flags().GetInt("limit");
		fileName, _ := cmd.Flags().GetString("filepath");

		props := scrapper.NewScraperProps(url, limit, fileName)
		resultCh := scrapper.Scrap(props)

		output(resultCh)
	},
}

func getUrl(args []string) string {
	if len(args) != 1 {
		log.Fatalln("courls must accept only one parameter that a link to site")
	}

	url := args[0]

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		log.Fatalln("courls must have before link a http:// or https://")
	}

	return url
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("filepath", "f", "", "specify filepath to result file")
	rootCmd.Flags().IntP("limit", "l", 1000, "specify limit of links to be visited")
}

func output(results <-chan *scrapper.CompletedUrl) {
	fmt.Println("Count   Code    URL")
	fmt.Println("------  -----   ------------------------")

	for result := range results {
		fmt.Printf("%5d   %3d     %s\n", result.Id, result.StatusCode, result.Url)
	}
}
