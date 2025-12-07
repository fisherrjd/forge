package initialize

import (
    "fmt"    
)


func Run(path string, useDefaults bool) error{
	return generateToml(path, useDefaults)
}

func generateToml(tomlPath string, useDefaults bool) error {
    // Implementation here
    return nil
}