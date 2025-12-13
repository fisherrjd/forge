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

// Objective: using the command forge build -> generate a pom.xml (used in maven) from the forge.toml file
// Secondary Objective: If the pom.xml already exists we need to parse and add any new dependencies / changes seen in the forge.toml
