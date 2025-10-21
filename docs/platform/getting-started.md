# Getting Started with AWS HPC Platform

## Overview

AWS HPC Platform is a flexible framework for running scientific computing applications on AWS with architecture-specific optimizations. This guide will walk you through setting up and using the platform.

## Installation

### Prerequisites

- Go 1.21 or later
- Docker Desktop with buildx support
- AWS CLI v2 configured with appropriate credentials
- Git

### Build the CLI

```bash
cd aws-hpc
make build
```

This creates `bin/aws-hpc` binary.

### Install Globally (Optional)

```bash
make install
```

This installs to `$GOPATH/bin/aws-hpc`.

## Quick Start

### 1. Explore Available Applications

```bash
aws-hpc app list
```

### 2. View Application Details

```bash
aws-hpc app info geos-chem
```

### 3. Validate an Application

```bash
aws-hpc app validate applications/geos-chem
```

### 4. Build Base Images

Base images contain pre-compiled compilers and optimized math libraries:

```bash
# Build a specific architecture
aws-hpc base build amd/zen4

# Build all AMD images
aws-hpc base build amd/all
```

### 5. Build Application Containers

```bash
# Build for specific architecture
aws-hpc app build geos-chem --arch c7a

# Build for all supported architectures
aws-hpc app build geos-chem --all-arch

# Build and push to registry
aws-hpc app build geos-chem --arch c7a --push
```

### 6. Deploy Infrastructure

```bash
aws-hpc app deploy geos-chem --env production
```

This creates:
- AWS Batch compute environments
- Job queues with spot/on-demand instances
- IAM roles and policies
- S3 buckets for data

### 7. Submit a Job

```bash
aws-hpc job submit geos-chem \
    --env benchmark \
    --input s3://my-bucket/input/ \
    --output s3://my-bucket/output/
```

### 8. Monitor Job

```bash
# Check status
aws-hpc job status JOB_ID

# View logs
aws-hpc job logs JOB_ID

# Follow logs
aws-hpc job logs JOB_ID --follow
```

### 9. Cost Analysis

```bash
# Estimate cost before running
aws-hpc cost estimate geos-chem --arch c7a --runtime 4h

# Compare costs across architectures
aws-hpc cost estimate geos-chem --compare --runtime 4h

# Analyze historical costs
aws-hpc cost analyze --days 30

# Get optimization recommendations
aws-hpc cost optimize geos-chem
```

## Architecture Overview

### Container Layering

```
Application Container (geos-chem:c7a-v0.1.0)
    ↓ FROM
Base Image (hpc-base-amd-zen4:20251018)
    ↓ FROM
Amazon Linux 2023
```

**Benefits:**
- Faster application builds (base layers cached)
- Shared library updates propagate easily
- Consistent compiler/library versions

### Supported Architectures

| Family | Generation     | AWS Instances | Math Library |
|--------|---------------|---------------|--------------|
| AMD    | Zen 4 (znver4) | c7a          | AMD AOCL 4.2 |
| AMD    | Zen 3 (znver3) | c6a          | AMD AOCL 4.2 |
| AMD    | Zen 2 (znver2) | c5a          | AMD AOCL 4.2 |
| Intel  | Sapphire Rapids| c7i          | Intel MKL 2024.2 |
| Intel  | Ice Lake       | c6i          | Intel MKL 2024.2 |
| Intel  | Cascade Lake   | c5           | Intel MKL 2024.2 |
| ARM    | Graviton 4 (V2)| c8g          | ARM PL 24.04 |
| ARM    | Graviton 3 (V1)| c7g          | ARM PL 24.04 |
| ARM    | Graviton 2 (N1)| c6g          | ARM PL 24.04 |

## Adding a New Application

See [Research Computing Applications](research-computing-applications.md) for application examples.

### Quick Steps

1. **Copy template:**
   ```bash
   cp -r applications/_template applications/myapp
   ```

2. **Edit app.yaml:**
   ```yaml
   name: "myapp"
   display_name: "My Application"
   version: "0.1.0-alpha"
   ```

3. **Create Dockerfile:**
   ```dockerfile
   ARG BASE_IMAGE={{BASE_IMAGE}}
   FROM ${BASE_IMAGE}

   # Build your application
   ```

4. **Test:**
   ```bash
   aws-hpc app validate applications/myapp
   aws-hpc app build myapp --arch c7a
   ```

## Configuration

### Environment Variables

- `AWS_PROFILE` - AWS profile to use
- `AWS_REGION` - AWS region (default: us-east-1)
- `OMP_NUM_THREADS` - OpenMP thread count

### Configuration Files

Application environments are defined in `environments/*.yaml`:

```yaml
name: "production"
description: "Production configuration"

compute:
  architecture: "c7a"
  instance_type: "c7a.4xlarge"
  vcpus: 16
  memory_mb: 32768

runtime:
  timeout_hours: 24
  retry_attempts: 2
```

## Troubleshooting

### Container Build Fails

**Problem:** Spack package not found

**Solution:** Check base image has required libraries, or add to `app.yaml` dependencies

---

**Problem:** Cross-compilation fails

**Solution:** Ensure you're using buildx multi-platform builds

### Job Fails

**Problem:** Out of memory

**Solution:** Increase memory allocation in job submission

---

**Problem:** Input data not found

**Solution:** Verify S3 bucket permissions and paths

### Performance Issues

**Problem:** Slow execution

**Solution:**
- Check OpenMP thread count matches vCPUs
- Consider switching to Graviton for cost/performance
- Use spot instances for non-urgent work

## Best Practices

### Cost Optimization

1. **Use Spot Instances** for non-urgent workloads (70% savings)
2. **Choose Graviton** when possible (30% better price/performance)
3. **Right-size instances** - don't over-provision
4. **Enable S3 lifecycle policies** for old data

### Performance

1. **Match thread count to vCPUs:** `OMP_NUM_THREADS=$(nproc)`
2. **Use architecture-specific builds** for maximum performance
3. **Test scaling** before large production runs
4. **Monitor metrics** to identify bottlenecks

### Security

1. **Never commit credentials** to Git
2. **Use IAM roles** for EC2/Batch instances
3. **Enable S3 encryption** for sensitive data
4. **Restrict security group** access

## Next Steps

- [Research Computing Applications](research-computing-applications.md)
- [User Personas](../PERSONAS.md)
- [Project Roadmap](https://github.com/scttfrdmn/aws-hpc/blob/main/ROADMAP.md)
- [Contributing Guide](https://github.com/scttfrdmn/aws-hpc/blob/main/CONTRIBUTING.md)

## Getting Help

- GitHub Issues: https://github.com/your-org/aws-hpc/issues
- Documentation: https://docs.aws-hpc.org
- Examples: `examples/` directory

## License

Copyright 2025 Scott Friedman

Licensed under the Apache License, Version 2.0
