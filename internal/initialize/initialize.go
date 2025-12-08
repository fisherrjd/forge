package initialize

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/pelletier/go-toml/v2"
)

type Config struct {
    Project      ProjectInfo              `toml:"project"`
    Build        BuildConfig              `toml:"build"`
    Dependencies map[string][]Dependency  `toml:"dependencies"`
}

type ProjectInfo struct {
    Group    string `toml:"group"`
    Artifact string `toml:"artifact"`
    Version  string `toml:"version"`
}

type BuildConfig struct {
    JavaVersion string `toml:"java-version"`
    Encoding    string `toml:"encoding"`
}

type Dependency struct {
    Name    string `toml:"name"`
    Version string `toml:"version"`
}

func Run(path string, useDefaults bool) error{
	return generateToml(path, useDefaults)
}

func generateToml(tomlPath string, useDefaults bool) error {
    // Implementation here
    return nil
}