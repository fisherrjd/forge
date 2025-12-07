package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

const version = "0.1.0"

// versionCmd represents the version subcommand.
var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version of the CLI",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Version:", version)
  },
}