// Copyright 2022 Mohammad Abdolirad
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
	"github.com/atkrad/wait4x/internal/pkg/errors"
	"github.com/atkrad/wait4x/pkg/checker/influxdb"
	"github.com/atkrad/wait4x/pkg/waiter"
	"github.com/spf13/cobra"
)

// NewInfluxDBCommand creates the influxdb sub-command
func NewInfluxDBCommand() *cobra.Command {
	influxdbCommand := &cobra.Command{
		Use:   "influxdb SERVER_URL",
		Short: "Check InfluxDB connection",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.NewCommandError("SERVER_URL is required argument for the influxdb command")
			}

			return nil
		},
		Example: `
  # Checking InfluxDB connection
  wait4x influxdb http://localhost:8086
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			interval, _ := cmd.Flags().GetDuration("interval")
			timeout, _ := cmd.Flags().GetDuration("timeout")
			invertCheck, _ := cmd.Flags().GetBool("invert-check")

			ic := influxdb.New(args[0])
			ic.SetLogger(Logger)

			return waiter.Wait(
				ic.Check,
				waiter.WithTimeout(timeout),
				waiter.WithInterval(interval),
				waiter.WithInvertCheck(invertCheck),
			)
		},
	}

	return influxdbCommand
}