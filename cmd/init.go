package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/config"
)

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Create a default configuration file",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "safeanalyze.yaml"
		if len(args) > 0 {
			path = args[0]
		}

		if err := config.WriteExample(path); err != nil {
			return err
		}

		color.Green("Created configuration file: %s\n", path)
		fmt.Println("Edit this file to customize scanners, sanitization rules, and output format.")
		return nil
	},
}
