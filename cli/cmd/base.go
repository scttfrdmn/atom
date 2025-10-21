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

	"github.com/spf13/cobra"
)

var baseCmd = &cobra.Command{
	Use:   "base",
	Short: "Manage base images",
	Long:  "Build and manage HPC base container images with optimized libraries",
}

var baseListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available base images",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available base images:")
		fmt.Println("\nAMD:")
		fmt.Println("  hpc-base-amd-zen4:20251018    - AMD EPYC Genoa (Zen 4)")
		fmt.Println("  hpc-base-amd-zen3:20251018    - AMD EPYC Milan (Zen 3)")
		fmt.Println("  hpc-base-amd-zen2:20251018    - AMD EPYC Rome (Zen 2)")
		fmt.Println("\nIntel:")
		fmt.Println("  hpc-base-intel-spr:20251018   - Intel Sapphire Rapids")
		fmt.Println("  hpc-base-intel-icl:20251018   - Intel Ice Lake")
		fmt.Println("  hpc-base-intel-clk:20251018   - Intel Cascade Lake")
		fmt.Println("\nARM:")
		fmt.Println("  hpc-base-arm-graviton4:20251018 - AWS Graviton 4 (Neoverse V2)")
		fmt.Println("  hpc-base-arm-graviton3:20251018 - AWS Graviton 3 (Neoverse V1)")
		fmt.Println("  hpc-base-arm-graviton2:20251018 - AWS Graviton 2 (Neoverse N1)")

		// TODO: Query actual images from registry
		fmt.Println("\n[NOT IMPLEMENTED] Registry query functionality coming soon")
	},
}

var baseBuildCmd = &cobra.Command{
	Use:   "build [family/generation]",
	Short: "Build a base image",
	Long: `Build a base container image with optimized compilers and libraries.

Examples:
  # Build specific architecture
  aws-hpc base build amd/zen4

  # Build all images in a family
  aws-hpc base build amd/all

  # Build all base images
  aws-hpc base build all`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		push, _ := cmd.Flags().GetBool("push")

		fmt.Printf("Building base image: %s\n", target)
		if push {
			fmt.Println("Will push to registry after build")
		}

		// TODO: Implement base image building
		fmt.Println("\n[NOT IMPLEMENTED] Base image build functionality coming soon")
	},
}

var baseInfoCmd = &cobra.Command{
	Use:   "info [image]",
	Short: "Show base image information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		image := args[0]

		fmt.Printf("Base Image: %s\n", image)
		fmt.Println("\nIncluded libraries:")
		fmt.Println("  - GCC 11.5.0")
		fmt.Println("  - Spack v0.23.1")
		fmt.Println("  - AMD AOCL 4.2 (BLIS + libFLAME)")
		fmt.Println("  - OpenMPI 4.1.6")
		fmt.Println("  - HDF5 1.14.3")
		fmt.Println("  - NetCDF-C 4.9.2")
		fmt.Println("  - NetCDF-Fortran 4.6.1")

		fmt.Println("\nCompiler flags:")
		fmt.Println("  CFLAGS:   -march=znver4 -mavx512f -O3")
		fmt.Println("  FCFLAGS:  -march=znver4 -mavx512f -O3 -fopenmp")

		fmt.Println("\nTarget instances:")
		fmt.Println("  - c7a.xlarge, c7a.2xlarge, c7a.4xlarge, c7a.8xlarge")

		// TODO: Query actual image metadata
		fmt.Println("\n[NOT IMPLEMENTED] Image metadata query coming soon")
	},
}

func init() {
	// base build flags
	baseBuildCmd.Flags().Bool("push", false, "Push to registry after build")

	// Add subcommands
	baseCmd.AddCommand(baseListCmd)
	baseCmd.AddCommand(baseBuildCmd)
	baseCmd.AddCommand(baseInfoCmd)
}
