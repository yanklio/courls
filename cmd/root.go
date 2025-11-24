package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	scrapper "github.com/yanklio/courls/pkg"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "courls",
	Short: "Counter of urls in web domain",
	Long:  "Scraper counter of all urls in this web domain",
	Run: func(cmd *cobra.Command, args []string) {

		url := getUrl(args)

		filePath, _ := cmd.Flags().GetString("filepath")
		file := getFile(filePath)

		defer file.Close()

		limit, _ := cmd.Flags().GetInt("limit")

		c := scrapper.GetScrapper(url, file, limit)
		c.Visit(url)

		fmt.Fprintln(file, "-----------------------")
		fmt.Fprintln(file, scrapper.GetCount())
	},
}

func getUrl(args []string) string {

	if len(args) != 1 {
		log.Fatalln("courls must accept only one parameter that a link to site")
	}

	url := args[0]

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		log.Fatalln("courls must have before link a http or https")
	}

	return url
}

func getFile(filePath string) *os.File {

	if filePath == "" {
		file, err := os.Create("res.txt")

		if err != nil {
			log.Fatalln("Can't create file in current folder")
		}

		return file
	}

	file, err := os.Create(filePath)

	if err != nil {
		log.Fatalln("Can't create file in given directory")
	}

	return file
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("filepath", "f", "", "specify filepath to resFile")
	rootCmd.Flags().IntP("limit", "l", 1000000, "specify limit of links")
}
