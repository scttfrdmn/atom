# AWS HPC Platform - Project Structure

## Directory Layout

```
aws-hpc/
├── README.md                     # Project overview
├── VERSIONING.md                 # Versioning strategy
├── PROJECT_STRUCTURE.md          # This file
├── CONTRIBUTING.md               # Contribution guidelines
├── LICENSE                       # MIT License
├──.gitignore
├── Makefile                      # Build automation
│
├── pkg/                          # Platform code (v1.0.0)
│   ├── version.go               # Platform version
│   ├── aws/                     # AWS service clients
│   │   ├── batch.go            # AWS Batch management
│   │   ├── ec2.go              # EC2 instances
│   │   ├── ecr.go              # Container registry
│   │   ├── efs.go              # Elastic File System
│   │   ├── s3.go               # S3 storage
│   │   └── vpc.go              # Networking
│   ├── config/                  # Configuration management
│   │   ├── application.go      # Application spec loader
│   │   ├── environment.go      # Environment config
│   │   └── state.go            # State management
│   ├── container/               # Container build system
│   │   ├── builder.go          # Multi-arch builds
│   │   ├── base.go             # Base image management
│   │   ├── layer.go            # Layer caching
│   │   └── registry.go         # ECR operations
│   ├── job/                     # Job management
│   │   ├── scheduler.go        # Job scheduling
│   │   ├── queue.go            # Priority queues
│   │   ├── monitor.go          # Job monitoring
│   │   └── array.go            # Array jobs
│   ├── cost/                    # Cost management
│   │   ├── calculator.go       # Cost estimation
│   │   ├── optimizer.go        # Cost optimization
│   │   └── tracker.go          # Usage tracking
│   └── license/                 # License management
│       ├── flexlm.go           # FlexLM integration
│       ├── rlm.go              # RLM integration
│       └── tracker.go          # License usage tracking
│
├── applications/                # Application plugins
│   ├── _template/              # Template for new apps
│   │   ├── app.yaml           # Application specification template
│   │   ├── containers/        # Container definitions
│   │   ├── scripts/           # Runtime scripts
│   │   ├── configs/           # Configuration templates
│   │   └── environments/      # Environment configs
│   │
│   ├── geos-chem/              # GEOS-Chem application
│   │   ├── app.yaml           # Application specification
│   │   ├── README.md          # Application-specific docs
│   │   ├── CHANGELOG.md       # Version history
│   │   ├── containers/
│   │   │   ├── Dockerfile.template
│   │   │   ├── classic/
│   │   │   │   ├── entrypoint.sh
│   │   │   │   └── config-generator.py
│   │   │   └── gchp/
│   │   │       └── entrypoint.sh
│   │   ├── scripts/
│   │   │   ├── preprocess.sh
│   │   │   └── postprocess.sh
│   │   ├── configs/
│   │   │   └── templates/
│   │   │       ├── geoschem_config.yml
│   │   │       ├── HEMCO_Config.rc
│   │   │       └── HISTORY.rc
│   │   └── environments/
│   │       ├── benchmark.yaml
│   │       ├── production.yaml
│   │       └── transport.yaml
│   │
│   ├── gaussian/               # Gaussian (planned)
│   │   └── app.yaml
│   │
│   └── wrf/                    # WRF (planned)
│       └── app.yaml
│
├── base-images/                # HPC base container images
│   ├── README.md              # Base image documentation
│   ├── build.sh               # Build script for all bases
│   ├── amd/
│   │   ├── Dockerfile.zen4    # AMD EPYC Genoa
│   │   ├── Dockerfile.zen3    # AMD EPYC Milan
│   │   └── Dockerfile.zen2    # AMD EPYC Rome
│   ├── intel/
│   │   ├── Dockerfile.spr     # Sapphire Rapids
│   │   ├── Dockerfile.icl     # Ice Lake
│   │   └── Dockerfile.clk     # Cascade Lake
│   ├── arm/
│   │   ├── Dockerfile.graviton4  # Neoverse V2
│   │   ├── Dockerfile.graviton3  # Neoverse V1
│   │   └── Dockerfile.graviton2  # Neoverse N1
│   └── gpu/                   # GPU-enabled bases (future)
│       ├── Dockerfile.cuda-amd
│       └── Dockerfile.cuda-intel
│
├── infrastructure/             # Infrastructure as Code
│   ├── cdk/                   # AWS CDK (TypeScript)
│   │   ├── package.json
│   │   ├── cdk.json
│   │   ├── bin/
│   │   │   └── aws-hpc.ts    # CDK app entry point
│   │   └── lib/
│   │       ├── core-stack.ts  # VPC, networking
│   │       ├── compute-stack.ts # Batch, EC2
│   │       ├── storage-stack.ts # S3, EFS, FSx
│   │       ├── app-stack.ts    # Per-application stacks
│   │       └── license-stack.ts # License servers
│   └── modules/               # Reusable Terraform (future)
│
├── cli/                        # Command-line tool
│   ├── main.go                # CLI entry point
│   ├── go.mod
│   ├── go.sum
│   └── cmd/
│       ├── root.go            # Root command
│       ├── app.go             # Application management
│       ├── job.go             # Job submission
│       ├── cost.go            # Cost analysis
│       ├── environment.go     # Environment management
│       └── base.go            # Base image management
│
├── web/                        # Web interface (future)
│   ├── dashboard/             # Universal dashboard
│   │   ├── package.json
│   │   └── src/
│   └── app-specific/          # Per-app UI extensions
│
├── docs/                       # Documentation
│   ├── platform/
│   │   ├── architecture.md    # Platform architecture
│   │   ├── adding-applications.md
│   │   ├── containers.md      # Container strategy
│   │   ├── versioning.md      # Version management
│   │   ├── research-computing-applications.md
│   │   └── api-reference.md   # Platform API docs
│   ├── applications/
│   │   ├── geos-chem.md      # GEOS-Chem guide
│   │   ├── gaussian.md       # Gaussian guide
│   │   └── wrf.md            # WRF guide
│   └── guides/
│       ├── quick-start.md
│       ├── deployment.md
│       └── troubleshooting.md
│
├── examples/                   # Example configurations
│   ├── job-submissions/
│   │   ├── geos-chem-benchmark.yaml
│   │   └── array-job.yaml
│   └── environments/
│       └── custom-environment.yaml
│
└── scripts/                    # Utility scripts
    ├── init-project.sh        # Initialize new application
    ├── build-bases.sh         # Build all base images
    ├── deploy-app.sh          # Deploy application
    └── test-app.sh            # Test application
```

## Key Design Principles

### 1. **Separation of Concerns**

- **Platform** (`pkg/`): Stable, reusable AWS infrastructure
- **Applications** (`applications/`): Self-contained, pluggable apps
- **Base Images** (`base-images/`): Shared compiler/library layers
- **Infrastructure** (`infrastructure/`): Deployment automation

### 2. **Versioning Independence**

- Platform has its own version (`pkg/v1.0.0`)
- Each application has its own version (`geos-chem/v0.2.0`)
- Base images are date-tagged (`hpc-base-amd-zen4:20251018`)

### 3. **Container Layering**

```
Application Container
    ↓ (FROM)
Base Image (compilers + libraries)
    ↓ (FROM)
OS Base (Amazon Linux 2023)
```

**Benefits:**
- Faster application builds
- Shared library updates propagate easily
- Clear separation of infrastructure vs. application

### 4. **Application Self-Description**

Each application defines its own requirements in `app.yaml`:
- Compute architectures
- Math libraries
- Parallelism type (OpenMP/MPI)
- Storage requirements
- License requirements

### 5. **Platform API Stability**

`pkg/` provides stable APIs that applications depend on:
- `pkg/config.LoadApplication()` - Load app.yaml
- `pkg/container.Build()` - Build containers
- `pkg/job.Submit()` - Submit jobs
- `pkg/cost.Estimate()` - Estimate costs

Breaking changes require major version bump.

## File Naming Conventions

### Application Files
- `app.yaml` - Application specification (required)
- `README.md` - Application documentation
- `CHANGELOG.md` - Version history
- `Dockerfile.template` - Container template with placeholders

### Base Images
- `Dockerfile.{arch}` - e.g., `Dockerfile.zen4`, `Dockerfile.graviton3`
- Pattern: `Dockerfile.{generation}`

### Environment Configs
- `{purpose}.yaml` - e.g., `benchmark.yaml`, `production.yaml`
- Lowercase, descriptive names

### Scripts
- `{action}-{target}.sh` - e.g., `build-bases.sh`, `deploy-app.sh`
- Kebab-case

## Configuration File Formats

### Application Specification (`app.yaml`)
- **Format**: YAML
- **Schema**: See `applications/_template/app.yaml`
- **Required Fields**: name, version, platform_version, compute, containers

### Environment Configuration (`environments/*.yaml`)
- **Format**: YAML
- **Purpose**: Runtime configuration per use case
- **Examples**: benchmark, production, debug

### Infrastructure Configuration
- **CDK**: TypeScript (`*.ts`)
- **Terraform**: HCL (`*.tf`) - future

## Development Workflow

### Adding a New Application

1. Copy template: `cp -r applications/_template applications/myapp`
2. Edit `applications/myapp/app.yaml`
3. Create Dockerfile template
4. Add runtime scripts
5. Test build: `aws-hpc app build myapp`
6. Test deployment: `aws-hpc app deploy myapp --env test`

### Building Base Images

```bash
cd base-images
./build.sh amd/zen4
./build.sh --all  # Build all architectures
```

### Testing Changes

```bash
# Test platform changes
cd pkg && go test ./...

# Test application
cd applications/geos-chem
aws-hpc app validate .

# Test infrastructure
cd infrastructure/cdk
npm test
```

## Git Repository Structure

### Branches
- `main` - Stable releases
- `develop` - Development integration
- `feature/*` - Feature branches
- `app/*` - Application-specific work

### Tags
- Platform: `pkg/v1.0.0`
- Applications: `geos-chem/v0.2.0`
- Base images: `base-amd-zen4/20251018`

### Monorepo Strategy

**Single repository** contains:
- Platform code
- All applications
- Base image definitions
- Infrastructure code
- Documentation

**Benefits:**
- Atomic cross-component changes
- Simplified dependency management
- Unified CI/CD
- Easier for contributors

## Build & Release Process

### Base Images
1. Update Dockerfiles
2. Build: `./base-images/build.sh --all`
3. Tag with date: `YYYYMMDD`
4. Push to ECR
5. Update application `app.yaml` references

### Applications
1. Update `app.yaml` version
2. Build containers: `aws-hpc app build geos-chem --all-arch`
3. Run tests: `aws-hpc app test geos-chem`
4. Tag: `git tag geos-chem/v0.2.0`
5. Update CHANGELOG

### Platform
1. Update `pkg/version.go`
2. Run tests: `go test ./...`
3. Build CLI: `make build`
4. Tag: `git tag pkg/v1.0.0`
5. Create GitHub release

## Next Steps

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.
See [docs/platform/adding-applications.md](docs/platform/adding-applications.md) for detailed application development guide.
