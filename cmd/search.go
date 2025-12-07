package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

const search_text = "Searching for dependency..."

// searchCmd
var searchCmd = &cobra.Command{
  Use:   "search",
  Short: "Search for java dependencies",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(search_text)
  },
}