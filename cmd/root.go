// Copyright © 2017 edenecke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var cfgAPIKey string
var PurgeURL string
var PurgeService string
var PurgeSurrKey string
var PurgeFile string
var PurgeSleep int
var PurgeSoft bool
var VERSION string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-purge-fastly",
	Short: "",
	Long: ``,
}

func Execute(version string) {
	VERSION = version

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.go-purge-fastly.yaml)")
	RootCmd.PersistentFlags().StringVar(&cfgAPIKey, "apikey", os.Getenv("FASTLY_API_TOKEN"), "Fastly API key, if not set uses env FASTLY_API_KEY)")
	RootCmd.PersistentFlags().StringVar(&PurgeFile, "file", "purge.txt", "Input file with url list to purge from fastly.")
	RootCmd.PersistentFlags().IntVar(&PurgeSleep, "sleep", 500, "Amount of time to wait between purge requests (ms).")
	RootCmd.PersistentFlags().BoolVarP(&PurgeSoft, "soft", "s", true, "Sends a soft purge request")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".go-purge-fastly")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
