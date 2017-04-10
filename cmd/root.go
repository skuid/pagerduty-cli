// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	pagerduty "github.com/PagerDuty/go-pagerduty"
	"github.com/spf13/cobra"
)

func Client() (client *pagerduty.Client) {
	return pagerduty.NewClient(apiToken)
}

var apiToken string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pagerduty-cli",
	Short: "A CLI interface for PagerDuty.",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&apiToken, "api-token", "", "PagerDuty API Token. Also supports environment variable \"PGDUTY_TOKEN\".")

	if apiToken == "" && os.Getenv("PGDUTY_TOKEN") == "" {
		fmt.Println("Please provide a PagerDuty API Token.")
		os.Exit(1)
	} else {
		apiToken = os.Getenv("PGDUTY_TOKEN")
	}

}
