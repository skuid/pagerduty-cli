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
	"time"

	"fmt"
	"os"

	pagerduty "github.com/PagerDuty/go-pagerduty"
	"github.com/spf13/cobra"
)

var duration string
var serviceId string

// windowCmd represents the window command
var windowCmd = &cobra.Command{
	Use:   "window",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// client := pagerduty.NewClient(apiToken)
	},
}

var listCmd = &cobra.Command{
	Use:   "list [options]",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client := Client()

		windows, err := client.ListMaintenanceWindows(pagerduty.ListMaintenanceWindowsOptions{})

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		for _, window := range windows.MaintenanceWindows {
			fmt.Println(window.Description)
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create [options]",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		client := Client()

		lengthOfWindow, err := time.ParseDuration(duration)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		start := time.Now()
		end := start.Add(lengthOfWindow)

		serviceReference := pagerduty.APIObject{ID: serviceId, Type: "service_reference"}

		services := make([]pagerduty.APIObject, 1)

		services = append(services, serviceReference)

		mwindow := pagerduty.MaintenanceWindow{
			APIObject:   pagerduty.APIObject{Type: "maintenance_window"},
			StartTime:   start.String(),
			EndTime:     end.String(),
			Description: "Maintenance Window created by pagerduty-cli.",
			Services:    services,
		}

		created, err := client.CreateMaintaienanceWindows(mwindow)

		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}

		fmt.Println("%v", created)

		// fmt.Println("%s put into maintenence mode until %s", created.Services[0].Name, end.String())

	},
}

func init() {
	RootCmd.AddCommand(windowCmd)
	windowCmd.AddCommand(createCmd)
	windowCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// windowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// windowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().StringVar(&duration, "duration", "01:00:00", "Length of the maintenence window. Ex: 1h30m.")
	createCmd.Flags().StringVar(&serviceId, "service", "", "ID of the service to create a maintenence window for.")

}
