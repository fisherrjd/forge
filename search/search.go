package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SearchResult struct {
	Response struct {
		NumFound int `json:"numFound"`
		Docs     []struct {
			ID            string `json:"id"`
			GroupID       string `json:"g"`
			ArtifactID    string `json:"a"`
			LatestVersion string `json:"latestVersion"`
			VersionCount  int    `json:"versionCount"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"docs"`
	} `json:"response"`
}

func Search(query string, maxResults int) (*SearchResult, error) {
	baseURL := "https://search.maven.org/solrsearch/select"

	params := url.Values{}
	params.Add("q", query)
	params.Add("rows", fmt.Sprintf("%d", maxResults))
	params.Add("wt", "json")

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SearchByArtifact searches Maven Central by artifact name (a:query)
func SearchByArtifact(artifactName string, maxResults int) (*SearchResult, error) {
	// Search with 'a:' prefix to search artifact names only
	query := "a:" + artifactName
	return Search(query, maxResults)
}