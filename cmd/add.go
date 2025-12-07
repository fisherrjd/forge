package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

const add_text = "Adding dependency to pom.xml..."

// addCmd
var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add a dependency to pom.xml",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(add_text)
  },
}