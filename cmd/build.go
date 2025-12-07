package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

const build_text = "Building pom.xml..."

// addCmd
var buildCmd = &cobra.Command{
  Use:   "build",
  Short: "Build pom.xml",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(build_text)
  },
}