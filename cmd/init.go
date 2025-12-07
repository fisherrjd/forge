package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

const init_text = "Init Forge usage"

// addCmd
var initCmd = &cobra.Command{
  Use:   "init",
  Short: "Begin the usage of forge",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(init_text)
  },
}