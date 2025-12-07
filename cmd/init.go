package cmd

import (
    "fmt"
    "github.com/fisherrjd/forge/internal/initialize"
    "github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
    Use:   "init [path]",
    Short: "Initialize a new Forge project",
    Args:  cobra.MaximumNArgs(1), // 0 or 1 args (path is optional)
    RunE: func(cmd *cobra.Command, args []string) error {
        path := "."
        if len(args) > 0 {
            path = args[0]
        }
        
        useDefaults, _ := cmd.Flags().GetBool("defaults")
        
        fmt.Printf("Initializing Forge\n")
        return initialize.Run(path, useDefaults)
    },
}

func init() {
    initCmd.Flags().Bool("defaults", false, "Use default values without prompting")
}