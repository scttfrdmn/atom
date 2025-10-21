# Versioning Strategy

AWS HPC Platform uses a **dual versioning system** inspired by [aws-ide](https://github.com/yourusername/aws-ide) to track platform stability and application features independently.

## Overview

```
geos-chem v0.2.0 (platform: v1.0.0)
          ^^^^^            ^^^^^
          │                └─ Platform/Infrastructure Version
          └─ Application Version
```

## Current Versions

| Component | Version | Status |
|-----------|---------|--------|
| Platform (`pkg/`) | v1.0.0-dev | In Development |
| GEOS-Chem | v0.1.0-alpha | Extracting |
| Gaussian | - | Planned |
| WRF | - | Planned |

## Platform Version (`pkg/` module)

**Current**: `v1.0.0-dev`
**Git Tag Format**: `pkg/v1.0.0`

### Semantic Versioning

- **MAJOR** (`X.0.0`): Breaking API changes
  - Changed function signatures in `pkg/`
  - Removed public APIs
  - Modified application specification schema (breaking)

- **MINOR** (`1.X.0`): New features (backward compatible)
  - New packages under `pkg/`
  - New application spec fields (optional)
  - New compute architectures

- **PATCH** (`1.0.X`): Bug fixes
  - Bug fixes in platform code
  - Performance improvements
  - Documentation updates

### Platform v1.0.0 Target APIs

- `pkg/aws/batch.go`: AWS Batch management
- `pkg/aws/ec2.go`: EC2 instance management
- `pkg/aws/ecr.go`: Container registry
- `pkg/aws/s3.go`: S3 storage management
- `pkg/aws/vpc.go`: Networking
- `pkg/container/builder.go`: Multi-architecture container builds
- `pkg/container/optimizer.go`: Architecture-specific optimizations
- `pkg/job/scheduler.go`: Job scheduling and queuing
- `pkg/config/application.go`: Application specification loader
- `pkg/cost/calculator.go`: Cost estimation and tracking

## Application Versions

**Format**: `<app>/vX.Y.Z`
**Git Tag Format**: `geos-chem/v0.2.0`

### Semantic Versioning

- **MAJOR** (`X.0.0`): Breaking changes to application
  - Incompatible configuration changes
  - Major algorithm changes
  - Removed features

- **MINOR** (`0.X.0`): New features
  - New environments
  - New container architectures
  - New configuration options

- **PATCH** (`0.2.X`): Bug fixes
  - Bug fixes
  - Performance optimizations
  - Documentation updates

## Base Image Versions

**Format**: `hpc-base-<family>-<arch>:YYYYMMDD`
**Examples**:
- `hpc-base-amd-zen4:20251018`
- `hpc-base-intel-spr:20251018`
- `hpc-base-arm-graviton4:20251018`

### Versioning Strategy

- **Date-based**: Images tagged with build date
- **Latest**: Rolling `latest` tag for current stable
- **Frozen**: Specific dates for reproducibility

### Update Cadence

- **Security patches**: As needed (CVEs)
- **Library updates**: Monthly
- **Compiler updates**: Quarterly
- **Major versions**: Annually

## Version Compatibility

### Platform ↔ Application

Applications specify minimum platform version:

```yaml
# applications/geos-chem/app.yaml
platform_version: ">=1.0.0"
```

Platform checks compatibility at runtime:
```go
if app.PlatformVersion > platform.Version {
    return errors.New("application requires newer platform")
}
```

### Application ↔ Base Images

Applications specify base image requirements:

```yaml
# applications/geos-chem/app.yaml
containers:
  base_images:
    amd-zen4: "hpc-base-amd-zen4:20251018"
    intel-spr: "hpc-base-intel-spr:latest"
    arm-graviton4: "hpc-base-arm-graviton4:>=20251001"
```

## Release Process

### Platform Release

1. Update `pkg/version.go`
2. Tag: `git tag pkg/v1.0.0`
3. Build and test all applications
4. Update CHANGELOG.md
5. Create GitHub release

### Application Release

1. Update `applications/<app>/app.yaml`
2. Build containers for all architectures
3. Tag: `git tag geos-chem/v0.2.0`
4. Test deployment
5. Update application CHANGELOG

### Base Image Release

1. Update Dockerfiles in `base-images/`
2. Build for all architectures
3. Tag: `hpc-base-amd-zen4:YYYYMMDD`
4. Push to ECR
5. Test with sample applications

## Migration Strategy

When platform APIs change:

### Non-Breaking (Minor/Patch)

- Applications continue working
- Update documentation
- Deprecation warnings if needed

### Breaking (Major)

- Announce in advance (30 days)
- Provide migration guide
- Support N-1 version for 90 days
- Update all applications in repo
- Release notes with upgrade path

## Version File Locations

```
pkg/version.go                           # Platform version
applications/geos-chem/app.yaml          # App version
base-images/amd/Dockerfile               # Base image version
```

## Checking Versions

```bash
# Platform version
aws-hpc version

# Application version
aws-hpc app version geos-chem

# Base image versions
aws-hpc base-images list
```

## Example Version Timeline

```
2025-10-18: Platform v1.0.0-dev, GEOS-Chem v0.1.0-alpha
2025-11-01: Platform v1.0.0-beta.1, GEOS-Chem v0.1.0-beta
2025-12-01: Platform v1.0.0, GEOS-Chem v0.1.0
2026-01-15: GEOS-Chem v0.2.0 (new environments)
2026-02-01: Gaussian v0.1.0 (first release)
2026-03-01: Platform v1.1.0 (new features)
```
