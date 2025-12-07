package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

const log_text = "Create forge.lock"

// addCmd
var lockCmd = &cobra.Command{
  Use:   "lock",
  Short: "Create a forge.lock for managing dependencies",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(log_text)
  },
}