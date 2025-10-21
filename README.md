# AWS HPC Platform

A flexible, application-agnostic platform for deploying high-performance computing applications on AWS with architecture-optimized containers and cost-efficient job scheduling.

## Overview

AWS HPC Platform provides a reusable infrastructure for running computational applications (GEOS-Chem, Gaussian, WRF, VASP, etc.) on AWS with:

- **Architecture-optimized containers** - Automatic builds for AMD (Zen 2/3/4), Intel (Cascade/Ice/Sapphire), ARM (Graviton 2/3/4)
- **Containerization strategy** - Layered base images (compilers/libraries) + application layers
- **Single-node & MPI support** - OpenMP (Classic) and distributed (GCHP-style) workloads
- **Cost optimization** - Spot instances, priority queues, Graviton savings
- **Job management** - AWS Batch integration with monitoring and retry logic
- **Multi-application** - Add new applications without changing platform code

## Architecture

```
Platform (v1.0.0) - Stable infrastructure APIs
    ↓
Applications - Pluggable application definitions
    ├── GEOS-Chem (Classic & GCHP)
    ├── Gaussian
    ├── WRF
    └── [Your app here]
```

## Project Status

**Current Version**: v0.1.0-alpha
**Platform API**: v1.0.0-dev

🚧 **In Active Development** - Platform extraction from [aws-geos-chem](https://github.com/yourusername/aws-geos-chem) in progress.

### Implemented
- [ ] Core platform abstractions
- [ ] Application specification format
- [ ] Container build system
- [ ] Base image layers
- [ ] GEOS-Chem application plugin

### Planned
- [ ] Gaussian application plugin
- [ ] WRF application plugin
- [ ] CLI tool (`aws-hpc`)
- [ ] Web dashboard
- [ ] Cost tracking system

## Quick Start

*Coming soon - platform under development*

```bash
# Install CLI
brew install aws-hpc  # or download from releases

# Deploy an application
aws-hpc app deploy geos-chem --env benchmark

# Submit a job
aws-hpc job submit geos-chem \
  --architecture c7a \
  --input s3://my-bucket/input/ \
  --output s3://my-bucket/output/

# Monitor job
aws-hpc job status <job-id>
```

## Container Strategy

### Layered Architecture

```
Base Images (maintained separately):
├── hpc-base-amd-zen4:latest
│   └── GCC, OpenMPI, AMD AOCL (BLIS + libFLAME)
├── hpc-base-intel-spr:latest
│   └── GCC, OpenMPI, Intel MKL
└── hpc-base-arm-graviton4:latest
    └── GCC, OpenMPI, ARM Performance Libraries

Application Layers (versioned independently):
├── geos-chem:14.4.3-c7a
│   └── FROM hpc-base-amd-zen4:latest
├── geos-chem:14.4.3-c7i
│   └── FROM hpc-base-intel-spr:latest
└── gaussian:16-c7a
    └── FROM hpc-base-amd-zen4:latest
```

**Benefits:**
- Base images shared across applications
- Faster application builds (no recompiling libraries)
- Version application independently from infrastructure
- Security updates to base propagate to all apps

## Application Types

### Classic (Single-Node, OpenMP)
- GEOS-Chem Classic
- Gaussian
- ORCA
- NWChem (single-node mode)

**Characteristics:**
- Single EC2 instance
- Shared memory parallelism (OpenMP)
- Instance families: c7a, c7i, c8g
- Optimal for: Small to medium workloads

### MPI (Multi-Node, Distributed)
- GEOS-Chem High Performance (GCHP)
- WRF
- VASP
- NWChem (MPI mode)

**Characteristics:**
- Multiple EC2 instances
- Message passing (MPI)
- EFA networking for low latency
- Optimal for: Large-scale simulations

## Documentation

- [Platform Architecture](docs/platform/architecture.md)
- [Adding Applications](docs/platform/adding-applications.md)
- [Container Build System](docs/platform/containers.md)
- [Versioning Strategy](docs/platform/versioning.md)
- [Application: GEOS-Chem](docs/applications/geos-chem.md)

## Related Projects

- **[aws-geos-chem](https://github.com/yourusername/aws-geos-chem)** - Original GEOS-Chem deployment (reference implementation)
- **[aws-ide](https://github.com/yourusername/aws-ide)** - IDE platform that inspired this architecture

## Contributing

Contributions welcome! See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT License - See [LICENSE](LICENSE)

---

**Status**: 🏗️ Platform under active development
**First Application**: GEOS-Chem (extracting from aws-geos-chem project)
**Target Release**: Q1 2026
