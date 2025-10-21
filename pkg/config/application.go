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

// Package config provides configuration management for applications
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Application represents a complete application specification
type Application struct {
	Name            string              `yaml:"name"`
	DisplayName     string              `yaml:"display_name"`
	Version         string              `yaml:"version"`
	PlatformVersion string              `yaml:"platform_version"`
	Metadata        ApplicationMetadata `yaml:"metadata"`
	Variants        []Variant           `yaml:"variants"`
	Compute         ComputeSpec         `yaml:"compute"`
	Containers      ContainerSpec       `yaml:"containers"`
	Storage         StorageSpec         `yaml:"storage"`
	Environments    []Environment       `yaml:"environments"`
	Cost            CostSpec            `yaml:"cost,omitempty"`
	Licensing       LicensingSpec       `yaml:"licensing,omitempty"`
	GPU             GPUSpec             `yaml:"gpu,omitempty"`
	Networking      NetworkingSpec      `yaml:"networking,omitempty"`
}

// ApplicationMetadata contains application metadata
type ApplicationMetadata struct {
	Description  string       `yaml:"description"`
	Homepage     string       `yaml:"homepage"`
	Documentation string      `yaml:"documentation"`
	Repository   string       `yaml:"repository"`
	License      string       `yaml:"license"`
	Maintainers  []Maintainer `yaml:"maintainers"`
	Tags         []string     `yaml:"tags"`
}

// Maintainer represents an application maintainer
type Maintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

// Variant represents an application variant (e.g., Classic vs GCHP)
type Variant struct {
	Name            string `yaml:"name"`
	DisplayName     string `yaml:"display_name"`
	Description     string `yaml:"description"`
	Type            string `yaml:"type"` // single-node, multi-node
	Parallelism     string `yaml:"parallelism"` // openmp, mpi, serial, gpu
	UpstreamVersion string `yaml:"upstream_version"`
}

// ComputeSpec defines compute requirements
type ComputeSpec struct {
	Architectures []Architecture `yaml:"architectures"`
	Batch         BatchConfig    `yaml:"batch"`
}

// Architecture defines a CPU architecture configuration
type Architecture struct {
	Name          string      `yaml:"name"`
	Family        string      `yaml:"family"` // amd, intel, arm
	Generation    string      `yaml:"generation"` // zen4, sapphirerapids, neoverse-v2
	InstanceTypes []string    `yaml:"instance_types"`
	CompilerFlags []string    `yaml:"compiler_flags"`
	MathLibrary   MathLibrary `yaml:"math_library"`
	BaseImage     string      `yaml:"base_image"`
}

// MathLibrary defines math library configuration
type MathLibrary struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	BLAS    string `yaml:"blas,omitempty"`
	LAPACK  string `yaml:"lapack,omitempty"`
}

// BatchConfig defines AWS Batch configuration
type BatchConfig struct {
	MinVCPUs           int     `yaml:"min_vcpus"`
	MaxVCPUs           int     `yaml:"max_vcpus"`
	SpotBidPercentage  int     `yaml:"spot_bid_percentage"`
	Queues             []Queue `yaml:"queues"`
}

// Queue defines a job queue configuration
type Queue struct {
	Name                 string                 `yaml:"name"`
	Priority             int                    `yaml:"priority"`
	ComputeEnvironments  []ComputeEnvironment   `yaml:"compute_environments"`
}

// ComputeEnvironment defines a compute environment
type ComputeEnvironment struct {
	Type          string   `yaml:"type"` // spot, on-demand
	Architectures []string `yaml:"architectures"`
	MaxVCPUs      int      `yaml:"max_vcpus"`
}

// ContainerSpec defines container configuration
type ContainerSpec struct {
	Registry     string                        `yaml:"registry"` // ecr, dockerhub
	Repository   string                        `yaml:"repository"`
	BuildSystem  string                        `yaml:"build_system"` // docker-buildx
	Variants     map[string]ContainerVariant   `yaml:"variants"`
	Dependencies []string                      `yaml:"dependencies"`
}

// ContainerVariant defines a container variant build configuration
type ContainerVariant struct {
	Dockerfile string            `yaml:"dockerfile"`
	Context    string            `yaml:"context"`
	BuildArgs  map[string]string `yaml:"build_args"`
}

// StorageSpec defines storage requirements
type StorageSpec struct {
	Input   StorageLocation `yaml:"input"`
	Output  StorageLocation `yaml:"output"`
	Scratch ScratchStorage  `yaml:"scratch"`
	Shared  *SharedStorage  `yaml:"shared,omitempty"`
}

// StorageLocation defines a storage location
type StorageLocation struct {
	Type       string           `yaml:"type"` // s3, efs, fsx-lustre
	Bucket     string           `yaml:"bucket,omitempty"`
	Prefix     string           `yaml:"prefix,omitempty"`
	Lifecycle  *LifecyclePolicy `yaml:"lifecycle,omitempty"`
}

// LifecyclePolicy defines S3 lifecycle rules
type LifecyclePolicy struct {
	TransitionIA      int `yaml:"transition_ia,omitempty"`
	TransitionGlacier int `yaml:"transition_glacier,omitempty"`
	Expiration        int `yaml:"expiration,omitempty"`
}

// ScratchStorage defines local scratch storage
type ScratchStorage struct {
	Type   string `yaml:"type"` // ebs, instance-store
	SizeGB int    `yaml:"size_gb"`
	VolumeType string `yaml:"type"` // gp3, io2
	IOPS   int    `yaml:"iops,omitempty"`
}

// SharedStorage defines shared filesystem
type SharedStorage struct {
	Type           string `yaml:"type"` // efs, fsx-lustre
	SizeGB         int    `yaml:"size_gb"`
	ThroughputMode string `yaml:"throughput_mode,omitempty"` // bursting, provisioned
}

// Environment defines a runtime environment configuration
type Environment struct {
	Name        string `yaml:"name"`
	Config      string `yaml:"config"`
	Description string `yaml:"description"`
}

// CostSpec defines cost estimation parameters
type CostSpec struct {
	EstimateMethod string             `yaml:"estimate_method"`
	Baseline       BaselineCost       `yaml:"baseline"`
	ScalingFactors map[string]float64 `yaml:"scaling_factors"`
}

// BaselineCost defines baseline cost parameters
type BaselineCost struct {
	Architecture string  `yaml:"architecture"`
	RuntimeHours float64 `yaml:"runtime_hours"`
	CostPerHour  float64 `yaml:"cost_per_hour"`
}

// LicensingSpec defines license requirements
type LicensingSpec struct {
	Type          string `yaml:"type"` // none, flexlm, rlm, custom
	Server        string `yaml:"server,omitempty"`
	Feature       string `yaml:"feature,omitempty"`
	TokensPerJob  int    `yaml:"tokens_per_job,omitempty"`
}

// GPUSpec defines GPU requirements
type GPUSpec struct {
	Required bool     `yaml:"required"`
	Types    []string `yaml:"types"`
	Count    int      `yaml:"count"`
	MemoryGB int      `yaml:"memory_gb"`
}

// NetworkingSpec defines networking requirements
type NetworkingSpec struct {
	EFA            bool `yaml:"efa"`
	PlacementGroup bool `yaml:"placement_group"`
}

// LoadApplication loads an application specification from app.yaml
func LoadApplication(path string) (*Application, error) {
	appYAMLPath := filepath.Join(path, "app.yaml")

	data, err := os.ReadFile(appYAMLPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read app.yaml: %w", err)
	}

	var app Application
	if err := yaml.Unmarshal(data, &app); err != nil {
		return nil, fmt.Errorf("failed to parse app.yaml: %w", err)
	}

	// Validate application
	if err := app.Validate(); err != nil {
		return nil, fmt.Errorf("invalid application specification: %w", err)
	}

	return &app, nil
}

// Validate validates the application specification
func (a *Application) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("name is required")
	}
	if a.Version == "" {
		return fmt.Errorf("version is required")
	}
	if a.PlatformVersion == "" {
		return fmt.Errorf("platform_version is required")
	}
	if len(a.Variants) == 0 {
		return fmt.Errorf("at least one variant is required")
	}
	if len(a.Compute.Architectures) == 0 {
		return fmt.Errorf("at least one architecture is required")
	}

	// Validate each architecture
	for _, arch := range a.Compute.Architectures {
		if err := arch.Validate(); err != nil {
			return fmt.Errorf("invalid architecture %s: %w", arch.Name, err)
		}
	}

	return nil
}

// Validate validates an architecture configuration
func (a *Architecture) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("name is required")
	}
	if a.Family == "" {
		return fmt.Errorf("family is required")
	}
	if len(a.InstanceTypes) == 0 {
		return fmt.Errorf("at least one instance type is required")
	}
	if a.BaseImage == "" {
		return fmt.Errorf("base_image is required")
	}
	return nil
}

// GetArchitecture returns the architecture configuration by name
func (a *Application) GetArchitecture(name string) (*Architecture, error) {
	for _, arch := range a.Compute.Architectures {
		if arch.Name == name {
			return &arch, nil
		}
	}
	return nil, fmt.Errorf("architecture %s not found", name)
}

// GetVariant returns the variant configuration by name
func (a *Application) GetVariant(name string) (*Variant, error) {
	for _, variant := range a.Variants {
		if variant.Name == name {
			return &variant, nil
		}
	}
	return nil, fmt.Errorf("variant %s not found", name)
}

// GetEnvironment returns the environment configuration by name
func (a *Application) GetEnvironment(name string) (*Environment, error) {
	for _, env := range a.Environments {
		if env.Name == name {
			return &env, nil
		}
	}
	return nil, fmt.Errorf("environment %s not found", name)
}
