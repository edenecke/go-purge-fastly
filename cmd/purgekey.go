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

// purgekeyCmd represents the purgekey command
var purgekeyCmd = &cobra.Command{
	Use:   "purgekey",
	Short: "Purge based on key",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := fastly.NewClient(cfgAPIKey)
		if err != nil {
			log.Fatal(err)
		}

		purge, err := client.PurgeKey( &fastly.PurgeKeyInput{
			Service: PurgeService,
			Key: PurgeSurrKey,
			Soft: PurgeSoft,
		} )

		fmt.Printf("Purge Key (%s) with Service ID: (%s)", PurgeSurrKey, PurgeService)
		if err != nil {
			log.Fatal(err)
			fmt.Printf("Purge-Error (%s)\n", err)
		} else {
			fmt.Printf("Purge-Response (%s)\n", purge)
		}

		time.Sleep(time.Duration(PurgeSleep)*time.Millisecond)

		_, err = os.Stat(PurgeFile)
		if os.IsNotExist(err) {
			fmt.Printf("+ file not found, skipping freshness check (%s)\n", PurgeFile)
	  } else {
				file, err := os.Open(PurgeFile)
				fmt.Printf("+ reading file (%s)\n", PurgeFile)
				if err != nil {
						log.Fatal(err)
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)

				for scanner.Scan() {
						PurgeURL = scanner.Text()

						RequestURL, err := http.Get(PurgeURL)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Printf("%s\tAge: (%s) X-Cache: (%s)\tLast-Modified (%s)\tURL (%s)\n", time.Now(), RequestURL.Header["Age"], RequestURL.Header["X-Cache"], RequestURL.Header["Last-Modified"], PurgeURL)
				}
		}
	},
}

func init() {
	RootCmd.AddCommand(purgekeyCmd)
	purgekeyCmd.Flags().StringVar(&PurgeSurrKey, "surrkey", "", "Key to purge")
}
