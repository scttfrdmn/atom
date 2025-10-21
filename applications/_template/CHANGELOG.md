# Changelog

All notable changes to this application will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial application structure

## [0.1.0-alpha] - 2025-10-18

### Added
- Initial alpha release
- Support for AMD EPYC Zen 4 (c7a instances)
- Support for Intel Sapphire Rapids (c7i instances)
- Support for AWS Graviton 4 (c8g instances)
- Basic container builds with architecture-specific optimizations
- S3 integration for input/output data
- Test environment configuration

### Known Issues
- Performance not yet optimized for all architectures
- Limited testing on production workloads
- Documentation incomplete

[Unreleased]: https://github.com/your-org/aws-hpc/compare/your-app/v0.1.0-alpha...HEAD
[0.1.0-alpha]: https://github.com/your-org/aws-hpc/releases/tag/your-app/v0.1.0-alpha
