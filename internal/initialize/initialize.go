package initialize

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config holds the configuration for a Forge project
type Config struct {
	Project ProjectInfo
	Build   BuildConfig
}

// ProjectInfo contains project metadata
type ProjectInfo struct {
	Group    string
	Artifact string
	Version  string
}

// BuildConfig contains build-related settings
type BuildConfig struct {
	JavaVersion string
	Encoding    string
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
	content := buildTomlContent(config)

	if err := writeTomlFile(tomlPath, content); err != nil {
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

// buildTomlContent generates the TOML file content with comments
func buildTomlContent(config Config) string {
	var content strings.Builder

	content.WriteString("[project]\n")
	content.WriteString(fmt.Sprintf("group = %q\n", config.Project.Group))
	content.WriteString(fmt.Sprintf("artifact = %q\n", config.Project.Artifact))
	content.WriteString(fmt.Sprintf("version = %q\n", config.Project.Version))
	content.WriteString("\n")

	content.WriteString("[build]\n")
	content.WriteString(fmt.Sprintf("java-version = %q\n", config.Build.JavaVersion))
	content.WriteString(fmt.Sprintf("encoding = %q\n", config.Build.Encoding))
	content.WriteString("\n")

	content.WriteString("[repositories]\n")
	content.WriteString("# Where to fetch dependencies from\n")
	content.WriteString("# maven-central = \"https://repo.maven.apache.org/maven2/\"\n")
	content.WriteString("\n")

	content.WriteString("[dependencies.main]\n")
	content.WriteString("# Production dependencies go here\n")
	content.WriteString("\n")

	content.WriteString("[dependencies.test]\n")
	content.WriteString("# Test dependencies go here\n")
	content.WriteString("\n")

	content.WriteString("[dependencies.runtime]\n")
	content.WriteString("# Runtime-only dependencies (optional)\n")
	content.WriteString("\n")

	content.WriteString("[dependencies.provided]\n")
	content.WriteString("# Provided scope dependencies (optional)\n")

	return content.String()
}

// writeTomlFile writes the TOML content to the specified path
func writeTomlFile(path, content string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// getDefaultConfig returns the default configuration
func getDefaultConfig() Config {
	return Config{
		Project: ProjectInfo{
			Group:    "com.example",
			Artifact: "my-app",
			Version:  "0.1.0",
		},
		Build: BuildConfig{
			JavaVersion: "17",
			Encoding:    "UTF-8",
		},
	}
}

// promptForConfig interactively prompts the user for configuration values
func promptForConfig() (Config, error) {
	reader := bufio.NewReader(os.Stdin)
	config := Config{}

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

	config.Build.Encoding, err = promptWithDefault(reader, "Encoding", "UTF-8")
	if err != nil {
		return config, err
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
