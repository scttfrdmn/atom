# Your Application Name

Brief description of your application.

## Overview

Detailed description of what the application does and its scientific domain.

## Version Information

- **Application Version**: 0.1.0-alpha
- **Platform Version**: >=1.0.0-dev
- **Upstream Version**: 1.0.0

## Supported Architectures

- **AMD EPYC**
  - Zen 4 (c7a instances)
  - Zen 3 (c6a instances)
- **Intel Xeon**
  - Sapphire Rapids (c7i instances)
- **AWS Graviton**
  - Graviton 4 (c8g instances)

## Prerequisites

### Input Data

Describe what input data is required and where to obtain it:

- Dataset 1: Available at https://...
- Dataset 2: Contact provider@example.com

### AWS Resources

- **S3 Buckets**: For input data and results storage
- **AWS Batch**: For job execution
- **IAM Roles**: With appropriate permissions

### Licenses (if applicable)

Describe any license requirements:

- License type: Commercial / Academic / Open Source
- How to obtain: Contact vendor or download from...
- License server setup: See docs/license-setup.md

## Quick Start

### 1. Build Container

```bash
# Build for specific architecture
aws-hpc app build your-app --arch c7a

# Or build for all supported architectures
aws-hpc app build your-app --all-arch
```

### 2. Prepare Input Data

```bash
# Upload input data to S3
aws s3 sync ./local-data/ s3://your-app-input-data/datasets/v1/
```

### 3. Submit Job

```bash
# Submit using predefined test environment
aws-hpc job submit your-app \
    --env test \
    --input s3://your-app-input-data/datasets/v1/ \
    --output s3://your-app-results/run-001/

# Or submit with custom parameters
aws-hpc job submit your-app \
    --arch c7a \
    --vcpus 8 \
    --memory 16384 \
    --input s3://your-app-input-data/custom/ \
    --output s3://your-app-results/custom-run/
```

### 4. Monitor Job

```bash
# Check job status
aws-hpc job status JOB_ID

# View logs
aws-hpc job logs JOB_ID

# Download results
aws s3 sync s3://your-app-results/run-001/ ./results/
```

## Configuration

### Environment Files

Pre-configured environments are available in `environments/`:

- **test.yaml**: Quick test with minimal resources
- **production.yaml**: Full production configuration

### Custom Configuration

Create a custom environment file:

```yaml
# environments/custom.yaml
name: "custom"
description: "Custom configuration"

compute:
  architecture: "c7a"
  instance_type: "c7a.4xlarge"
  vcpus: 16
  memory_mb: 32768

parameters:
  # Application-specific parameters
  input_param1: "value1"
  input_param2: "value2"

runtime:
  timeout_hours: 24
  retry_attempts: 2
```

Then submit with:

```bash
aws-hpc job submit your-app --config environments/custom.yaml
```

## Application-Specific Documentation

### Input File Format

Describe the expected input file formats and structure.

### Output Files

Describe what output files are produced and their formats.

### Performance Tuning

Tips for optimizing performance:

- **Threading**: Set OMP_NUM_THREADS to match vCPUs
- **Memory**: Recommended 2GB per vCPU
- **I/O**: Use gp3 EBS volumes for scratch space

## Architecture-Specific Notes

### AMD EPYC (Zen 4)

- Best overall performance
- Optimized with AMD AOCL libraries
- Recommended for production workloads

### Intel Xeon (Sapphire Rapids)

- Good performance with Intel MKL
- Wide AVX-512 support

### AWS Graviton 4

- Best cost/performance ratio
- ARM-optimized with ARM Performance Libraries
- Recommended for cost-sensitive workloads

## Troubleshooting

### Common Issues

**Issue**: Container build fails with "package not found"

**Solution**: Ensure all dependencies are listed in `app.yaml` under `containers.dependencies`

---

**Issue**: Job fails with out-of-memory error

**Solution**: Increase memory allocation or reduce problem size

---

**Issue**: Input data not found

**Solution**: Verify S3 bucket permissions and object paths

### Getting Help

- Check application logs: `aws-hpc job logs JOB_ID`
- Review CloudWatch logs in AWS Console
- Open an issue on GitHub

## Performance Benchmarks

Include performance data if available:

| Architecture | vCPUs | Runtime | Cost | Cost/hour |
|--------------|-------|---------|------|-----------|
| c7a.4xlarge  | 16    | 2.5h    | $0.87| $0.3468   |
| c8g.4xlarge  | 16    | 3.0h    | $0.62| $0.2076   |
| c7i.4xlarge  | 16    | 2.8h    | $0.76| $0.2720   |

## Citation

If using this application in research, please cite:

```
Author, A., et al. (2024). Application Name: Description.
Journal Name, Volume(Issue), Pages. DOI: 10.xxxx/xxxxx
```

## References

- Application homepage: https://your-app.org
- User manual: https://your-app.readthedocs.io
- GitHub repository: https://github.com/your-org/your-app
- Scientific paper: DOI link

## License

This application implementation is licensed under MIT License.

The upstream application may have a different license - see upstream documentation.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history.
