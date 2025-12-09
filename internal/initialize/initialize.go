package initialize

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

// ForgeConfig represents the entire forge.toml structure
type ForgeConfig struct {
	Project      Project           `toml:"project"`
	Build        Build             `toml:"build"`
	Repositories map[string]string `toml:"repositories"`
	Dependencies Dependencies      `toml:"dependencies"`
}

// Project defines the project metadata
type Project struct {
	Group    string `toml:"group"`
	Artifact string `toml:"artifact"`
	Version  string `toml:"version"`
}

// Build defines build configuration
type Build struct {
	JavaVersion string `toml:"java-version"`
}

// Dependencies defines all dependency scopes
type Dependencies struct {
	Main     map[string]string `toml:"main"`     // "group:artifact" -> "version"
	Test     map[string]string `toml:"test"`     // "group:artifact" -> "version"
	Runtime  map[string]string `toml:"runtime"`  // "group:artifact" -> "version"
	Provided map[string]string `toml:"provided"` // "group:artifact" -> "version"
}

// Coordinate represents a parsed dependency coordinate
type Coordinate struct {
	Group    string
	Artifact string
	Version  string
}

// ParseCoordinate parses a "group:artifact:version" string
func ParseCoordinate(coord string) (*Coordinate, error) {
	parts := strings.Split(coord, ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid coordinate format: expected 'group:artifact:version', got '%s'", coord)
	}
	return &Coordinate{
		Group:    parts[0],
		Artifact: parts[1],
		Version:  parts[2],
	}, nil
}

// LoadConfig reads and parses a forge.toml file
func LoadConfig(path string) (*ForgeConfig, error) {
	var config ForgeConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}
	return &config, nil
}

// SaveConfig writes a ForgeConfig to a TOML file
func SaveConfig(path string, config *ForgeConfig) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	encoder := toml.NewEncoder(f)
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	return nil
}

// Run initializes a new Forge project at the specified path
func Run(path string, useDefaults bool) error {
	config := getDefaultConfig()

	if !useDefaults {
		var err error
		config, err = promptForConfig()
		if err != nil {
			return fmt.Errorf("failed to gather config: %w", err)
		}
	}

	tomlPath := resolveTomlPath(path)

	// Ensure directory exists
	dir := filepath.Dir(tomlPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := SaveConfig(tomlPath, &config); err != nil {
		return err
	}

	fmt.Printf("Created forge.toml at %s\n", tomlPath)
	return nil
}

// resolveTomlPath converts a directory path to a full forge.toml file path
func resolveTomlPath(path string) string {
	fileInfo, err := os.Stat(path)
	if err == nil && fileInfo.IsDir() {
		return filepath.Join(path, "forge.toml")
	}

	if err != nil && os.IsNotExist(err) && filepath.Ext(path) == "" {
		return filepath.Join(path, "forge.toml")
	}

	return path
}

// getDefaultConfig returns the default configuration
func getDefaultConfig() ForgeConfig {
	return ForgeConfig{
		Project: Project{
			Group:    "com.example",
			Artifact: "my-app",
			Version:  "0.1.0",
		},
		Build: Build{
			JavaVersion: "17",
		},
		Repositories: make(map[string]string),
		Dependencies: Dependencies{
			Main:     make(map[string]string),
			Test:     make(map[string]string),
			Runtime:  make(map[string]string),
			Provided: make(map[string]string),
		},
	}
}

// promptForConfig interactively prompts the user for configuration values
func promptForConfig() (ForgeConfig, error) {
	reader := bufio.NewReader(os.Stdin)
	config := ForgeConfig{}

	var err error
	config.Project.Group, err = promptWithDefault(reader, "Project group", "com.example")
	if err != nil {
		return config, err
	}

	config.Project.Artifact, err = promptWithDefault(reader, "Project artifact", "my-app")
	if err != nil {
		return config, err
	}

	config.Project.Version, err = promptWithDefault(reader, "Project version", "0.1.0")
	if err != nil {
		return config, err
	}

	config.Build.JavaVersion, err = promptWithDefault(reader, "Java version", "17")
	if err != nil {
		return config, err
	}

	// Initialize empty maps
	config.Repositories = make(map[string]string)
	config.Dependencies = Dependencies{
		Main:     make(map[string]string),
		Test:     make(map[string]string),
		Runtime:  make(map[string]string),
		Provided: make(map[string]string),
	}

	return config, nil
}

// promptWithDefault prompts the user for input with a default value
func promptWithDefault(reader *bufio.Reader, prompt, defaultValue string) (string, error) {
	fmt.Printf("%s [%s]: ", prompt, defaultValue)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue, nil
	}

	return input, nil
}
