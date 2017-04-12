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
	"strings"
	"time"

	"fmt"
	"os"

	"text/tabwriter"

	pagerduty "github.com/PagerDuty/go-pagerduty"
	"github.com/spf13/cobra"
)

var duration string
var serviceID string

// windowCmd represents the window command
var windowCmd = &cobra.Command{
	Use:   "mwindow",
	Short: "Maintenance Windows",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// client := pagerduty.NewClient(apiToken)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all maintenance windows.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		client := Client()

		windows, err := client.ListMaintenanceWindows(pagerduty.ListMaintenanceWindowsOptions{})

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		columns := []string{
			"ID", "Service", "Start", "End",
		}

		fmt.Fprintln(w, strings.Join(columns, "\t"))

		for _, window := range windows.MaintenanceWindows {

			svc := window.Services[0]
			columns := []string{
				window.ID,
				fmt.Sprintf("%s - %s", svc.ID, svc.Summary),
				window.StartTime,
				window.EndTime,
			}

			fmt.Fprintln(w, strings.Join(columns, "\t"))
		}

		w.Flush()
	},
}

var createCmd = &cobra.Command{
	Use:   "create [options]",
	Short: "Create a maintenance window for a given service.",
	Long: `Create a new maintenance window for a given service. Duration can be in the form of 3h30m30s.
	Windows will be created from the current time and end at current time + duration.`,
	Run: func(cmd *cobra.Command, args []string) {

		client := Client()

		lengthOfWindow, err := time.ParseDuration(duration)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		start := time.Now()
		end := start.Add(lengthOfWindow)

		serviceReference := pagerduty.APIObject{ID: serviceID, Type: "service_reference"}

		services := make([]pagerduty.APIObject, 0)

		services = append(services, serviceReference)

		mwindow := pagerduty.MaintenanceWindow{
			APIObject:   pagerduty.APIObject{Type: "maintenance_window"},
			StartTime:   start.String(),
			EndTime:     end.String(),
			Description: "Maintenance Window created by pagerduty-cli.",
			Services:    services,
		}

		created, err := client.CreateMaintenanceWindows(mwindow)

		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}

		fmt.Printf("Maintenance window created for %s until %s (%s).", created.Services[0].Summary, end, duration)

	},
}

func init() {
	RootCmd.AddCommand(windowCmd)
	windowCmd.AddCommand(createCmd)
	windowCmd.AddCommand(listCmd)

	createCmd.Flags().StringVar(&duration, "duration", "01:00:00", "Length of the maintenence window. Ex: 1h30m.")
	createCmd.Flags().StringVar(&serviceID, "service", "", "ID of the service to create a maintenence window for.")

}
