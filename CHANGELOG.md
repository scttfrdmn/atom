# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project structure and documentation
- Platform API for application management (`pkg/config`)
- CLI tool with application, job, cost, and base image commands
- Application template with comprehensive examples
- Base image definitions for AMD, Intel, and ARM architectures
- GEOS-Chem as reference application implementation
- Container layering strategy (base images + application layers)
- Dual versioning system (platform + applications)
- Support for 9 CPU architectures across 3 families
- Documentation for adding new applications
- Research analysis of 25 common HPC applications

### Changed
- N/A (initial release)

### Deprecated
- N/A

### Removed
- N/A

### Fixed
- N/A

### Security
- N/A

## [1.0.0-dev] - 2025-10-18

### Added
- Initial development release
- Core platform infrastructure
- GEOS-Chem reference implementation

[Unreleased]: https://github.com/your-org/aws-hpc/compare/v1.0.0-dev...HEAD
[1.0.0-dev]: https://github.com/your-org/aws-hpc/releases/tag/v1.0.0-dev
