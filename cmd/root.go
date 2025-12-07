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
  Use:   "simple-cli",
  Short: "A simple CLI that greets you",
  Long:  "A simple CLI tool built with Cobra in Go that prints a greeting message. You can customize the greeting and name using flags.",
  Run: func(cmd *cobra.Command, args []string) {
    hello := fmt.Sprintf("%s, %s!", greeting, name)
    fmt.Println(hello)
  },
}

func Execute() {
  rootCmd.AddCommand(versionCmd)
  rootCmd.PersistentFlags().StringVar(&name, "name", "World", "Your name")
  rootCmd.PersistentFlags().StringVar(&greeting, "greeting", "Hello", "Greeting message")

  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}