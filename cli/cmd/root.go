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

// Package cmd implements CLI commands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/aws-hpc/pkg"
)

var (
	// Verbose enables verbose output
	Verbose bool
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "aws-hpc",
	Short: "AWS HPC Platform - Manage HPC applications on AWS",
	Long: `AWS HPC Platform is a flexible framework for running scientific computing
applications on AWS with architecture-specific optimizations.

Supports:
  - Multiple CPU architectures (AMD, Intel, ARM Graviton)
  - Container-based deployments
  - AWS Batch job management
  - S3 integration for data
  - Cost optimization

For more information, visit: https://github.com/your-org/aws-hpc`,
	Version: pkg.Version,
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Add subcommands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(jobCmd)
	rootCmd.AddCommand(costCmd)
	rootCmd.AddCommand(baseCmd)
}

// versionCmd shows version information
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		info := pkg.GetVersionInfo()
		fmt.Printf("aws-hpc version %s\n", info.Version)
		fmt.Printf("API version: %s\n", info.APIVersion)
		if info.GitCommit != "" {
			fmt.Printf("Git commit: %s\n", info.GitCommit)
		}
		if info.BuildDate != "" {
			fmt.Printf("Build date: %s\n", info.BuildDate)
		}
		if info.GoVersion != "" {
			fmt.Printf("Go version: %s\n", info.GoVersion)
		}
	},
}
