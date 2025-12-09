package add

import (
	"fmt"

	"github.com/fisherrjd/forge/internal/initialize"
)

// Run adds a dependency to the forge.toml in the current directory
func Run(scope, group, artifact, version string) error {
	tomlPath := "forge.toml"

	// Load existing config
	config, err := initialize.LoadConfig(tomlPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Add the dependency based on scope
	switch scope {
	case "main":
		AddMain(&config.Dependencies, group, artifact, version)
	case "test":
		AddTest(&config.Dependencies, group, artifact, version)
	case "runtime":
		AddRuntime(&config.Dependencies, group, artifact, version)
	case "provided":
		AddProvided(&config.Dependencies, group, artifact, version)
	default:
		return fmt.Errorf("unknown scope: %s (valid: main, test, runtime, provided)", scope)
	}

	// Save the updated config
	if err := initialize.SaveConfig(tomlPath, config); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("Added %s:%s:%s to %s scope\n", group, artifact, version, scope)
	return nil
}

// AddMain adds a dependency to the main scope
func AddMain(deps *initialize.Dependencies, group, artifact, version string) {
	if deps.Main == nil {
		deps.Main = make(map[string]string)
	}
	coord := fmt.Sprintf("%s:%s", group, artifact)
	deps.Main[coord] = version
}

// AddTest adds a dependency to the test scope
func AddTest(deps *initialize.Dependencies, group, artifact, version string) {
	if deps.Test == nil {
		deps.Test = make(map[string]string)
	}
	coord := fmt.Sprintf("%s:%s", group, artifact)
	deps.Test[coord] = version
}

// AddRuntime adds a dependency to the runtime scope
func AddRuntime(deps *initialize.Dependencies, group, artifact, version string) {
	if deps.Runtime == nil {
		deps.Runtime = make(map[string]string)
	}
	coord := fmt.Sprintf("%s:%s", group, artifact)
	deps.Runtime[coord] = version
}

// AddProvided adds a dependency to the provided scope
func AddProvided(deps *initialize.Dependencies, group, artifact, version string) {
	if deps.Provided == nil {
		deps.Provided = make(map[string]string)
	}
	coord := fmt.Sprintf("%s:%s", group, artifact)
	deps.Provided[coord] = version
}
