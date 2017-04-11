package cmd

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"time"
	"net/http"

	"github.com/spf13/cobra"
	fastly "github.com/sethvargo/go-fastly"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge an individual URL.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := fastly.NewClient(cfgAPIKey)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Open(PurgeFile)
		fmt.Printf("+ reading file (%s)\n", PurgeFile)
		if err != nil {
				log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
				PurgeURL = scanner.Text()

				RequestURLB, err := http.Get(PurgeURL)
				if err != nil {
					log.Fatal(err)
				}

				purge, err := client.Purge( &fastly.PurgeInput{
					URL: PurgeURL,
					Soft: PurgeSoft,
				} )

				fmt.Printf("Url: (%s)", PurgeURL)
				if err != nil {
					log.Fatal(err)
					fmt.Printf(" Purge-Error (%s)\n", err)
				} else {
					fmt.Printf(" Purge-Response (%s)\n", purge)
				}

				time.Sleep(time.Duration(PurgeSleep)*time.Millisecond)

				RequestURLA, err := http.Get(PurgeURL)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Before-Purge - Age: (%s) X-Cache: (%s) Last-Modified (%s)\n", RequestURLB.Header["Age"], RequestURLB.Header["X-Cache"], RequestURLB.Header["Last-Modified"], )
				fmt.Printf("After--Purge - Age: (%s) X-Cache: (%s) Last-Modified (%s)\n\n", RequestURLA.Header["Age"], RequestURLA.Header["X-Cache"], RequestURLA.Header["Last-Modified"], )
		}
	},
}

func init() {
	RootCmd.AddCommand(purgeCmd)
}
