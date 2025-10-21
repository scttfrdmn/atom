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

// Package pkg provides the AWS HPC Platform core functionality
package pkg

const (
	// Version is the platform version (semantic versioning)
	Version = "1.0.0-dev"

	// APIVersion is the API compatibility version
	// Change this when making breaking API changes
	APIVersion = "v1"

	// MinAppSpecVersion is the minimum app.yaml version supported
	MinAppSpecVersion = "0.1.0"
)

// VersionInfo contains detailed version information
type VersionInfo struct {
	Version    string `json:"version"`
	APIVersion string `json:"api_version"`
	GitCommit  string `json:"git_commit,omitempty"`
	BuildDate  string `json:"build_date,omitempty"`
	GoVersion  string `json:"go_version,omitempty"`
}

// GetVersionInfo returns detailed version information
func GetVersionInfo() VersionInfo {
	return VersionInfo{
		Version:    Version,
		APIVersion: APIVersion,
		// GitCommit, BuildDate, GoVersion will be set at build time via ldflags
	}
}
