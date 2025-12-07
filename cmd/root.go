package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)


var (
  greeting string
  name     string
)

var rootCmd = &cobra.Command{
  Use:   "forge",
  Short: "A CLI tool for managing Java Maven projects with ease",
  Long:  "Forge is a modern command-line tool that simplifies Java Maven project management.",
  Run: func(cmd *cobra.Command, args []string) {
    hello := fmt.Sprintf("%s, %s!", greeting, name)
    fmt.Println(hello)
  },
}

func Execute() {
  rootCmd.AddCommand(versionCmd)
  rootCmd.AddCommand(addCmd)
  rootCmd.AddCommand(searchCmd)
  rootCmd.AddCommand(buildCmd)
  rootCmd.PersistentFlags().StringVar(&name, "name", "World", "Your name")
  rootCmd.PersistentFlags().StringVar(&greeting, "greeting", "Hello", "Greeting message")

  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}