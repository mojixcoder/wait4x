package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"time"
	errs "github.com/atkrad/wait4x/internal/errors"
)

var (
	Interval time.Duration
	Timeout time.Duration
)

func NewWait4X() *cobra.Command {
	wait4x := &cobra.Command{
		Use:   "wait4x",
		Short: "wait4x allows waiting for a port or a service to enter into specify state",
		Long: `wait4x allows waiting for a port to enter into specify state or waiting for a service e.g. redis, mysql, postgres, ... to enter inter ready state`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Prevent showing usage when subcommand return error.
			cmd.SilenceUsage = true
		},
	}

	wait4x.PersistentFlags().DurationVarP(&Interval, "interval", "i", 1 * time.Second, "Interval time between each loop.")
	wait4x.PersistentFlags().DurationVarP(&Timeout, "timeout", "t", 10 * time.Second, "Timeout is the maximum amount of time that Wait4X will wait for a checking operation.")

	return wait4x
}

func Execute() {
	wait4x := NewWait4X()
	wait4x.AddCommand(NewTcpCommand())
	wait4x.AddCommand(NewHttpCommannd())
	wait4x.AddCommand(NewMysqlCommand())
	wait4x.AddCommand(NewRedisCommand())
	wait4x.AddCommand(NewVersionCommand())

	if err := wait4x.Execute(); err != nil {
		var commandError *errs.CommandError
		if errors.As(err, &commandError) {
			os.Exit(commandError.ExitCode)
		}

		os.Exit(errs.EXIT_ERROR)
	}
}