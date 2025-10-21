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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/aws-hpc/pkg/config"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Manage applications",
	Long:  "Build, deploy, and manage HPC applications",
}

var appValidateCmd = &cobra.Command{
	Use:   "validate [path]",
	Short: "Validate application specification",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appPath := args[0]

		fmt.Printf("Validating application at %s...\n", appPath)

		app, err := config.LoadApplication(appPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Validation failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✓ Validation passed")
		fmt.Printf("\nApplication: %s v%s\n", app.DisplayName, app.Version)
		fmt.Printf("Platform version: %s\n", app.PlatformVersion)
		fmt.Printf("Variants: %d\n", len(app.Variants))
		fmt.Printf("Architectures: %d\n", len(app.Compute.Architectures))

		if Verbose {
			fmt.Println("\nSupported architectures:")
			for _, arch := range app.Compute.Architectures {
				fmt.Printf("  - %s (%s %s)\n", arch.Name, arch.Family, arch.Generation)
			}
		}
	},
}

var appListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available applications",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement listing applications
		fmt.Println("Available applications:")
		fmt.Println("  geos-chem - Global 3-D atmospheric chemistry transport model")
		fmt.Println("\nUse 'aws-hpc app info <name>' for details")
	},
}

var appInfoCmd = &cobra.Command{
	Use:   "info [name]",
	Short: "Show application information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]

		// Load application from applications directory
		appPath := fmt.Sprintf("applications/%s", appName)
		app, err := config.LoadApplication(appPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading application: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Application: %s\n", app.DisplayName)
		fmt.Printf("Version: %s\n", app.Version)
		fmt.Printf("Description: %s\n", app.Metadata.Description)
		fmt.Printf("Homepage: %s\n", app.Metadata.Homepage)
		fmt.Printf("License: %s\n", app.Metadata.License)

		fmt.Println("\nVariants:")
		for _, variant := range app.Variants {
			fmt.Printf("  %s - %s\n", variant.Name, variant.Description)
		}

		fmt.Println("\nSupported Architectures:")
		for _, arch := range app.Compute.Architectures {
			fmt.Printf("  %s (%s %s) - %s\n",
				arch.Name, arch.Family, arch.Generation, arch.BaseImage)
		}

		fmt.Println("\nEnvironments:")
		for _, env := range app.Environments {
			fmt.Printf("  %s - %s\n", env.Name, env.Description)
		}
	},
}

var appBuildCmd = &cobra.Command{
	Use:   "build [name]",
	Short: "Build application containers",
	Long: `Build container images for specified application and architecture.

Examples:
  # Build for specific architecture
  aws-hpc app build geos-chem --arch c7a

  # Build for all architectures
  aws-hpc app build geos-chem --all-arch

  # Build and push to registry
  aws-hpc app build geos-chem --arch c7a --push`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		arch, _ := cmd.Flags().GetString("arch")
		allArch, _ := cmd.Flags().GetBool("all-arch")
		push, _ := cmd.Flags().GetBool("push")
		noPush, _ := cmd.Flags().GetBool("no-push")

		if !allArch && arch == "" {
			fmt.Fprintln(os.Stderr, "Error: Either --arch or --all-arch must be specified")
			os.Exit(1)
		}

		fmt.Printf("Building application: %s\n", appName)
		fmt.Printf("Architecture: %s\n", arch)
		if push && !noPush {
			fmt.Println("Will push to registry after build")
		}

		// TODO: Implement container build
		fmt.Println("\n[NOT IMPLEMENTED] Container build functionality coming soon")
	},
}

var appDeployCmd = &cobra.Command{
	Use:   "deploy [name]",
	Short: "Deploy application infrastructure",
	Long: `Deploy AWS infrastructure for the specified application.

This creates:
  - AWS Batch compute environments
  - Job queues
  - IAM roles and policies
  - S3 buckets (if needed)

Examples:
  aws-hpc app deploy geos-chem --env production
  aws-hpc app deploy gaussian --env test --region us-west-2`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		env, _ := cmd.Flags().GetString("env")
		region, _ := cmd.Flags().GetString("region")

		fmt.Printf("Deploying application: %s\n", appName)
		fmt.Printf("Environment: %s\n", env)
		fmt.Printf("Region: %s\n", region)

		// TODO: Implement deployment
		fmt.Println("\n[NOT IMPLEMENTED] Deployment functionality coming soon")
	},
}

func init() {
	// app build flags
	appBuildCmd.Flags().String("arch", "", "Target architecture (c7a, c7i, graviton4, etc.)")
	appBuildCmd.Flags().Bool("all-arch", false, "Build for all supported architectures")
	appBuildCmd.Flags().Bool("push", false, "Push to container registry after build")
	appBuildCmd.Flags().Bool("no-push", false, "Do not push to registry")

	// app deploy flags
	appDeployCmd.Flags().String("env", "production", "Environment name")
	appDeployCmd.Flags().String("region", "us-east-1", "AWS region")

	// Add subcommands
	appCmd.AddCommand(appValidateCmd)
	appCmd.AddCommand(appListCmd)
	appCmd.AddCommand(appInfoCmd)
	appCmd.AddCommand(appBuildCmd)
	appCmd.AddCommand(appDeployCmd)
}
