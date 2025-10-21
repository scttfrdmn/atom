# Contributing to AWS HPC Platform

Thank you for your interest in contributing to the AWS HPC Platform! This document provides guidelines for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Adding a New Application](#adding-a-new-application)
- [Submitting Changes](#submitting-changes)
- [Coding Standards](#coding-standards)
- [Testing Guidelines](#testing-guidelines)

## Code of Conduct

This project adheres to professional standards of conduct. Be respectful, inclusive, and constructive in all interactions.

## Getting Started

### Prerequisites

- Docker Desktop with buildx support
- AWS CLI v2
- Go 1.21 or later
- Node.js 18 or later (for CDK)
- Git

### Initial Setup

```bash
# Clone the repository
git clone https://github.com/your-org/aws-hpc.git
cd aws-hpc

# Install CLI dependencies
cd cli
go mod download

# Install CDK dependencies
cd ../infrastructure/cdk
npm install

# Build base images (optional, for testing)
cd ../../base-images
./build.sh amd/zen4
```

## Development Workflow

### Branch Strategy

- `main` - Stable releases only
- `develop` - Integration branch for development
- `feature/*` - Feature development branches
- `app/*` - Application-specific work

### Making Changes

1. Create a feature branch from `develop`:
   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following coding standards

3. Test your changes locally

4. Commit with descriptive messages:
   ```bash
   git commit -m "Add support for Gaussian application"
   ```

5. Push and create a pull request to `develop`

## Adding a New Application

### Step 1: Copy Template

```bash
cp -r applications/_template applications/your-app
cd applications/your-app
```

### Step 2: Create Application Specification

Edit `app.yaml` with your application details:

```yaml
name: "your-app"
display_name: "Your Application Name"
version: "0.1.0-alpha"
platform_version: ">=1.0.0-dev"

metadata:
  description: "Brief description"
  homepage: "https://your-app.org"
  license: "MIT"

variants:
  - name: "standard"
    type: "single-node"
    parallelism: "openmp"

compute:
  architectures:
    - name: "c7a"
      family: "amd"
      generation: "zen4"
      compiler_flags: ["-march=znver4", "-O3"]
      base_image: "hpc-base-amd-zen4:latest"
```

### Step 3: Create Dockerfiles

Choose your approach:

**Option A: Use Base Images (Recommended)**

Create `containers/Dockerfile.template`:

```dockerfile
ARG BASE_IMAGE={{BASE_IMAGE}}
FROM ${BASE_IMAGE}

# Application metadata
LABEL org.hpc.app.name="your-app"
LABEL org.hpc.app.version="{{VERSION}}"

# Download and build your application
WORKDIR /tmp/build
RUN git clone https://github.com/your-org/your-app.git && \
    cd your-app && \
    . /etc/profile.d/hpc-base.sh && \
    mkdir build && cd build && \
    cmake .. && \
    make -j$(nproc) && \
    make install
```

**Option B: Standalone Dockerfile**

Create architecture-specific Dockerfiles that install everything from scratch.

### Step 4: Add Runtime Scripts

Create `scripts/entrypoint.sh`:

```bash
#!/bin/bash
set -e

# Parse arguments
INPUT_DATA="$1"
OUTPUT_PATH="$2"

# Download input data from S3
aws s3 sync "$INPUT_DATA" /data/input/

# Run your application
cd /opt/run-dir
your-app --input /data/input --output /data/output

# Upload results to S3
aws s3 sync /data/output/ "$OUTPUT_PATH"
```

### Step 5: Add Configuration Templates

Create `configs/templates/` with your application's configuration files.

### Step 6: Document Your Application

Create `README.md` with:
- Installation instructions
- Usage examples
- Configuration options
- Troubleshooting tips

### Step 7: Test Your Application

```bash
# Build container
aws-hpc app build your-app --arch c7a

# Test locally
docker run your-app:c7a-latest --help

# Test on AWS Batch
aws-hpc app deploy your-app --env test
aws-hpc job submit your-app --env test
```

### Step 8: Submit for Review

1. Update `CHANGELOG.md` in your application directory
2. Create pull request with description of application
3. Request review from platform maintainers

## Submitting Changes

### Pull Request Guidelines

- **Title**: Use clear, descriptive titles
  - Good: "Add Gaussian quantum chemistry application"
  - Bad: "Update files"

- **Description**: Include:
  - What changes were made
  - Why the changes were needed
  - How to test the changes
  - Any breaking changes

- **Size**: Keep PRs focused and reasonably sized
  - Large applications can be split into multiple PRs
  - First PR: basic structure and app.yaml
  - Second PR: containers and build system
  - Third PR: documentation and examples

### Commit Message Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Test additions or changes
- `chore`: Build system or tooling changes

**Examples:**
```
feat(app): Add Gaussian quantum chemistry application

- Created app.yaml with single-node configuration
- Added Dockerfile templates for AMD, Intel, and ARM
- Implemented license server integration
- Added configuration generator for Gaussian input files

Closes #123
```

```
fix(container): Resolve BLAS library linking for AMD AOCL

The AMD AOCL libraries were not being found at runtime due to
incorrect library path configuration in CMake build.

Changed from BLA_VENDOR=AOCL to explicit BLAS_LIBRARIES path.

Fixes #456
```

## Coding Standards

### Go Code (Platform)

- Follow standard Go conventions
- Use `gofmt` for formatting
- Add doc comments for exported functions
- Write table-driven tests

```go
// LoadApplication loads an application specification from app.yaml
func LoadApplication(path string) (*Application, error) {
    data, err := os.ReadFile(filepath.Join(path, "app.yaml"))
    if err != nil {
        return nil, fmt.Errorf("failed to read app.yaml: %w", err)
    }

    var app Application
    if err := yaml.Unmarshal(data, &app); err != nil {
        return nil, fmt.Errorf("failed to parse app.yaml: %w", err)
    }

    return &app, nil
}
```

### TypeScript Code (CDK)

- Use TypeScript strict mode
- Follow AWS CDK best practices
- Use constructs for reusable components
- Document props interfaces

```typescript
export interface AppStackProps extends StackProps {
  readonly appSpec: ApplicationSpec;
  readonly vpcId: string;
}

export class AppStack extends Stack {
  constructor(scope: Construct, id: string, props: AppStackProps) {
    super(scope, id, props);

    // Implementation
  }
}
```

### Dockerfile Standards

- Use multi-stage builds where appropriate
- Minimize layer count
- Clean up in the same RUN command
- Use specific versions for dependencies
- Add labels for metadata

```dockerfile
# Install dependencies and clean up in one layer
RUN yum install -y \
        gcc gcc-c++ gcc-gfortran \
        make cmake git && \
    yum clean all && \
    rm -rf /var/cache/yum
```

### Shell Script Standards

- Use `#!/bin/bash` shebang
- Enable strict mode: `set -euo pipefail`
- Quote variables: `"$VARIABLE"`
- Add help text for scripts
- Use descriptive variable names

```bash
#!/bin/bash
set -euo pipefail

# Script: build-app.sh
# Description: Build application containers for specified architecture
# Usage: ./build-app.sh <app-name> <architecture>

APP_NAME="${1:?App name required}"
ARCH="${2:?Architecture required}"

echo "Building ${APP_NAME} for ${ARCH}..."
```

## Testing Guidelines

### Platform Code Tests (Go)

```bash
cd pkg
go test ./...
go test -race ./...  # Check for race conditions
go test -cover ./... # Check coverage
```

### Application Validation

```bash
# Validate app.yaml syntax
aws-hpc app validate applications/your-app

# Test container builds
aws-hpc app build your-app --arch c7a --no-push

# Test job submission (dry run)
aws-hpc job submit your-app --dry-run
```

### Infrastructure Tests (CDK)

```bash
cd infrastructure/cdk
npm test
npm run synth  # Verify CloudFormation synthesis
```

### Integration Tests

Create test cases in `applications/your-app/tests/`:

```bash
#!/bin/bash
# Test basic application execution

docker run your-app:c7a-latest \
    s3://test-bucket/input \
    s3://test-bucket/output

# Verify output
aws s3 ls s3://test-bucket/output/results.nc
```

## Documentation Standards

### Application README Template

```markdown
# Application Name

Brief description of the application.

## Overview

Detailed description of what the application does.

## Supported Architectures

- AMD EPYC (Zen 2, Zen 3, Zen 4)
- Intel Xeon (Cascade Lake, Ice Lake, Sapphire Rapids)
- AWS Graviton (2, 3, 4)

## Prerequisites

- Input data requirements
- AWS resources needed
- License requirements (if applicable)

## Quick Start

\`\`\`bash
# Build container
aws-hpc app build your-app

# Submit job
aws-hpc job submit your-app \
    --input s3://bucket/input \
    --output s3://bucket/output
\`\`\`

## Configuration

Description of configuration options.

## Troubleshooting

Common issues and solutions.

## References

- Application homepage
- Documentation links
- Citation information
```

## Version Management

### Platform Versioning

Platform version in `pkg/version.go`:

```go
package pkg

const Version = "1.0.0"
```

Tag releases as: `pkg/v1.0.0`

### Application Versioning

Application version in `app.yaml`:

```yaml
version: "0.2.0"
```

Tag releases as: `your-app/v0.2.0`

### Semantic Versioning Rules

**Platform:**
- MAJOR: Breaking API changes
- MINOR: New features (backward compatible)
- PATCH: Bug fixes

**Applications:**
- MAJOR: Breaking configuration changes
- MINOR: New features or container updates
- PATCH: Bug fixes or documentation

## License

By contributing to this project, you agree that your contributions will be licensed under the MIT License.

## Questions?

- Open an issue for bugs or feature requests
- Start a discussion for questions about the platform
- Join our community Slack channel (link)

## Acknowledgments

Thank you for contributing to AWS HPC Platform!
