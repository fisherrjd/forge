package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/fisherrjd/forge/internal/search"
)

const search_text = "Searching for dependency..."

// searchCmd
var searchCmd = &cobra.Command{
  Use:   "search [query]",
  Short: "Search for java dependencies",
  Args:  cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    query := args[0]
    fmt.Printf("Searching for: %s\n", query)
    
    // Search by artifact name
    results, err := search.SearchByArtifact(query, 5)
    if err != nil {
      fmt.Printf("Error searching: %v\n", err)
      return
    }
    
    fmt.Printf("\nFound %d results:\n\n", results.Response.NumFound)
    
    for i, doc := range results.Response.Docs {
      fmt.Printf("%d. %s:%s\n", i+1, doc.GroupID, doc.ArtifactID)
      fmt.Printf("   Latest: %s\n", doc.LatestVersion)
      fmt.Printf("   Versions: %d\n\n", doc.VersionCount)
    }
  },
}