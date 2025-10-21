// Copyright 2025 Scott Friedman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Configuration Generator for HPC Applications
// Generates application-specific configuration files from templates
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration structure
type Config map[string]interface{}

// Variables holds template substitution variables
type Variables map[string]string

func main() {
	// Parse command-line arguments
	templatePath := flag.String("template", "", "Path to configuration template file (required)")
	outputPath := flag.String("output", "", "Path to output configuration file (required)")
	validate := flag.Bool("validate", false, "Validate generated configuration")
	flag.Parse()

	// Check required arguments
	if *templatePath == "" || *outputPath == "" {
		fmt.Fprintln(os.Stderr, "Error: --template and --output are required")
		flag.Usage()
		os.Exit(1)
	}

	// Load template
	templateContent, err := os.ReadFile(*templatePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to read template file: %v\n", err)
		os.Exit(1)
	}

	// Get variables from environment and overrides
	variables := getVariables()

	// Apply variable overrides from remaining arguments
	for _, arg := range flag.Args() {
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			variables[strings.ToUpper(parts[0])] = parts[1]
		}
	}

	// Substitute variables in template
	configContent := substituteVariables(string(templateContent), variables)

	// Validate if requested
	if *validate {
		if err := validateConfig([]byte(configContent)); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Configuration validation failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Configuration validation: PASSED")
	}

	// Write output
	if err := os.MkdirAll(filepath.Dir(*outputPath), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*outputPath, []byte(configContent), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to write output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Configuration generated: %s\n", *outputPath)

	// Print summary
	fmt.Println("\nConfiguration variables:")
	for key, value := range variables {
		fmt.Printf("  %s: %s\n", key, value)
	}
}

// getVariables returns default variables from environment
func getVariables() Variables {
	vars := Variables{
		"INPUT_DIR":   getEnv("APP_DATA", "/data/input"),
		"OUTPUT_DIR":  getEnv("APP_OUTPUT", "/data/output"),
		"NUM_THREADS": getEnv("OMP_NUM_THREADS", fmt.Sprintf("%d", runtime.NumCPU())),
		"VERSION":     getEnv("APP_VERSION", "1.0.0"),
	}
	return vars
}

// substituteVariables replaces {{VARIABLE}} placeholders in template
func substituteVariables(template string, variables Variables) string {
	result := template
	for key, value := range variables {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = strings.ReplaceAll(result, placeholder, value)
	}
	return result
}

// validateConfig validates the generated configuration YAML
func validateConfig(content []byte) error {
	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return fmt.Errorf("invalid YAML: %w", err)
	}

	// Application-specific validation
	if _, ok := config["input_dir"]; !ok {
		return fmt.Errorf("input_dir not specified")
	}

	if _, ok := config["output_dir"]; !ok {
		return fmt.Errorf("output_dir not specified")
	}

	return nil
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
