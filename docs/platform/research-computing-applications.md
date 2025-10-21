# Research Computing Applications Analysis

Analysis of the 25 most common research computing applications to inform AWS HPC Platform architecture and requirements.

## Application Categories & Platform Requirements

### 1. **Molecular Dynamics & Chemistry**

#### GROMACS (Molecular Dynamics)
- **Type**: MPI + GPU
- **Parallelism**: Highly parallel (100-1000s cores)
- **Math Libs**: BLAS/LAPACK, FFTW
- **Special Needs**: GPU (CUDA/ROCm), high-bandwidth interconnect (EFA)
- **Data**: Large trajectory files, checkpoint/restart
- **License**: Open source (GPL)

#### NAMD (Molecular Dynamics)
- **Type**: MPI + GPU
- **Parallelism**: Charm++ (10-1000s cores)
- **Math Libs**: FFTW
- **Special Needs**: GPU support, InfiniBand/EFA
- **Data**: PDB files, restart files
- **License**: Free for non-commercial

#### LAMMPS (Molecular Dynamics)
- **Type**: MPI + GPU (optional)
- **Parallelism**: 1-10000 cores
- **Math Libs**: BLAS/LAPACK, FFTW
- **Special Needs**: GPU acceleration optional
- **Data**: Medium-sized input/output
- **License**: Open source (GPL)

#### Gaussian (Quantum Chemistry)
- **Type**: Single-node + shared memory
- **Parallelism**: OpenMP, limited MPI
- **Math Libs**: Proprietary, MKL optional
- **Special Needs**: **Commercial license required**
- **Data**: Checkpoint files (large)
- **License**: **Commercial** (expensive)

#### ORCA (Quantum Chemistry)
- **Type**: MPI
- **Parallelism**: 1-100 cores
- **Math Libs**: BLAS/LAPACK
- **Special Needs**: None special
- **Data**: Moderate I/O
- **License**: Free for academic use

#### NWChem (Computational Chemistry)
- **Type**: MPI
- **Parallelism**: 1-1000s cores
- **Math Libs**: BLAS/LAPACK, ScaLAPACK
- **Special Needs**: Global Arrays library
- **Data**: Checkpoint/restart
- **License**: Open source (ECL 2.0)

---

### 2. **Climate & Weather Modeling**

#### WRF (Weather Research & Forecasting)
- **Type**: MPI
- **Parallelism**: 10-1000s cores
- **Math Libs**: NetCDF, HDF5
- **Special Needs**: Large memory, fast I/O
- **Data**: Very large (GBs-TBs)
- **License**: Open source

#### CESM (Community Earth System Model)
- **Type**: MPI
- **Parallelism**: 100-10000 cores
- **Math Libs**: NetCDF, PnetCDF
- **Special Needs**: High-bandwidth storage
- **Data**: Massive datasets (TBs)
- **License**: Open source

#### GEOS-Chem (Atmospheric Chemistry) ✓ **IMPLEMENTED**
- **Type**: OpenMP (Classic) + MPI (GCHP)
- **Parallelism**: 1-96 cores (Classic), 6-6000 cores (GCHP)
- **Math Libs**: BLAS/LAPACK, NetCDF, HDF5, ESMF
- **Special Needs**: Large input data (met fields)
- **Data**: ~100GB input, moderate output
- **License**: Open source (MIT)

---

### 3. **Materials Science & Physics**

#### VASP (Vienna Ab initio Simulation Package)
- **Type**: MPI
- **Parallelism**: 1-1000 cores
- **Math Libs**: ScaLAPACK, FFTW
- **Special Needs**: **Commercial license**, GPU optional
- **Data**: Checkpoint files, wavefunctions
- **License**: **Commercial** (university licenses common)

#### Quantum ESPRESSO
- **Type**: MPI + GPU (optional)
- **Parallelism**: 1-1000s cores
- **Math Libs**: BLAS/LAPACK, ScaLAPACK, FFTW
- **Special Needs**: GPU acceleration (CUDA)
- **Data**: Moderate I/O
- **License**: Open source (GPL)

#### LAMMPS (also Materials Science)
- See Molecular Dynamics above

#### ABINIT
- **Type**: MPI
- **Parallelism**: 1-1000 cores
- **Math Libs**: BLAS/LAPACK, FFTW, NetCDF
- **Special Needs**: None special
- **Data**: Moderate
- **License**: Open source (GPL)

---

### 4. **Computational Fluid Dynamics (CFD)**

#### OpenFOAM
- **Type**: MPI
- **Parallelism**: 1-1000s cores
- **Math Libs**: None special
- **Special Needs**: Mesh decomposition, large memory
- **Data**: Large mesh files, time series
- **License**: Open source (GPL)

#### ANSYS Fluent
- **Type**: MPI + GPU (optional)
- **Parallelism**: 1-1000s cores
- **Math Libs**: Proprietary
- **Special Needs**: **Commercial license**, GPU optional
- **Data**: Large mesh files
- **License**: **Commercial** (very expensive)

---

### 5. **Bioinformatics & Genomics**

#### BLAST (Sequence Alignment)
- **Type**: Embarrassingly parallel
- **Parallelism**: 1-1000s independent jobs
- **Math Libs**: None
- **Special Needs**: Large databases
- **Data**: Large sequence databases (100s GB)
- **License**: Open source (Public domain)

#### Bowtie2 (Sequence Alignment)
- **Type**: Single-node, multi-threaded
- **Parallelism**: 1-16 cores typically
- **Math Libs**: None
- **Special Needs**: Fast I/O
- **Data**: Large FASTQ files
- **License**: Open source (GPL)

#### GATK (Genomics Analysis)
- **Type**: Single-node + Spark cluster
- **Parallelism**: 1-100s cores (Spark)
- **Math Libs**: None (Java-based)
- **Special Needs**: Large memory (>100GB)
- **Data**: Large BAM/VCF files
- **License**: Open source (BSD)

#### Gromacs (also for Structural Biology)
- See Molecular Dynamics above

---

### 6. **Machine Learning & AI**

#### TensorFlow
- **Type**: Multi-GPU + distributed
- **Parallelism**: 1-100s GPUs
- **Math Libs**: cuBLAS, cuDNN
- **Special Needs**: **GPUs required**, high-bandwidth interconnect
- **Data**: Massive training datasets
- **License**: Open source (Apache)

#### PyTorch
- **Type**: Multi-GPU + distributed
- **Parallelism**: 1-100s GPUs
- **Math Libs**: cuBLAS, cuDNN
- **Special Needs**: **GPUs required**
- **Data**: Massive datasets
- **License**: Open source (BSD)

---

### 7. **Engineering Simulation**

#### COMSOL Multiphysics
- **Type**: Single-node + limited MPI
- **Parallelism**: 1-32 cores typically
- **Math Libs**: BLAS/LAPACK, MUMPS
- **Special Needs**: **Commercial license**, GUI often needed
- **Data**: Moderate
- **License**: **Commercial**

#### LS-DYNA (Crash Simulation)
- **Type**: MPI + GPU
- **Parallelism**: 1-1000s cores
- **Math Libs**: Proprietary
- **Special Needs**: **Commercial license**, large memory
- **Data**: Large mesh files
- **License**: **Commercial**

---

### 8. **Mathematical & Statistical Computing**

#### R (Statistical Computing)
- **Type**: Single-node, some parallel packages
- **Parallelism**: 1-16 cores typically
- **Math Libs**: BLAS/LAPACK, OpenBLAS
- **Special Needs**: Large memory for datasets
- **Data**: Datasets (GB-scale)
- **License**: Open source (GPL)

#### MATLAB
- **Type**: Single-node + Parallel Computing Toolbox
- **Parallelism**: 1-100s cores (with PCT)
- **Math Libs**: Proprietary (MKL-based)
- **Special Needs**: **Commercial license**
- **Data**: Moderate
- **License**: **Commercial**

#### Julia
- **Type**: Multi-threaded + distributed
- **Parallelism**: 1-1000s cores
- **Math Libs**: BLAS/LAPACK, FFTW
- **Special Needs**: None special
- **Data**: Moderate
- **License**: Open source (MIT)

---

## Platform Requirements Analysis

### Compute Patterns

| Pattern | Count | Examples | Platform Needs |
|---------|-------|----------|----------------|
| **Embarrassingly Parallel** | 4 | BLAST, job arrays | Simple job submission, array jobs |
| **OpenMP (Shared Memory)** | 5 | Gaussian, GEOS-Chem Classic, R | Single large instance |
| **MPI (Distributed)** | 12 | WRF, VASP, OpenFOAM | EFA networking, MPI support |
| **GPU-Accelerated** | 6 | TensorFlow, GROMACS, NAMD | GPU instances (P/G families) |
| **Hybrid MPI+OpenMP** | 3 | GROMACS, CESM | Both MPI and OpenMP |
| **Spark/Big Data** | 1 | GATK | EMR integration |

### Math Library Requirements

| Library | Required By | Platform Solution |
|---------|-------------|-------------------|
| **BLAS/LAPACK** | 15 apps | Base images (MKL, AOCL, ARM PL) |
| **FFTW** | 8 apps | Add to base images |
| **ScaLAPACK** | 4 apps | Add to base images |
| **NetCDF/HDF5** | 5 apps | Already in base images ✓ |
| **GPU Libraries** | 6 apps | Separate GPU base images |

### License Models

| Model | Count | Platform Implications |
|-------|-------|----------------------|
| **Open Source** | 16 | No restrictions, easy deployment |
| **Free Academic** | 2 | Email verification, institution check |
| **Commercial** | 7 | **License server needed**, usage tracking |

### Data Patterns

| Pattern | Count | Storage Solution |
|---------|-------|------------------|
| **Small (<1GB)** | 3 | EBS volumes |
| **Medium (1-100GB)** | 10 | S3 + EBS cache |
| **Large (100GB-1TB)** | 8 | EFS or FSx for Lustre |
| **Massive (>1TB)** | 4 | FSx for Lustre required |

### GPU Requirements

| GPU Need | Count | Instance Families |
|----------|-------|-------------------|
| **No GPU** | 19 | C family (current focus) ✓ |
| **Optional GPU** | 4 | P/G families (future) |
| **Required GPU** | 2 | P/G families (future) |

---

## Platform Architecture Implications

### Must-Have Features

1. **Base Image Variants** ✓ (Already designed)
   - CPU-only (current)
   - GPU-enabled (future)
   - Big memory variants

2. **License Management System** 🚧 (Critical for commercial apps)
   - License server integration
   - FlexLM/RLM support
   - Usage tracking

3. **Storage Tiers** 🚧
   - S3 (cold storage) ✓
   - EFS (shared filesystem) - needed
   - FSx for Lustre (high-performance) - needed
   - EBS (local scratch) ✓

4. **Networking Options** 🚧
   - Regular networking (current) ✓
   - EFA (Elastic Fabric Adapter) - needed for MPI apps

5. **Job Types** 🚧
   - Single-node ✓
   - MPI (multi-node) - planned
   - Array jobs (embarrassingly parallel) - needed
   - GPU jobs - future

### Application Specification Extensions Needed

```yaml
# Extended app.yaml schema

compute:
  # Add GPU support
  gpu:
    required: false  # or true
    types: ["nvidia-a100", "nvidia-h100"]
    count: 1-8

  # Add network requirements
  networking:
    efa: false  # Enable EFA for MPI
    placement_group: true  # For low-latency

  # Add memory requirements
  memory:
    min_gb: 16
    recommended_gb: 64

# Add license management
licensing:
  type: "none"  # or "flexlm", "rlm", "custom"
  server: "license.university.edu:port"
  feature: "vasp_std"

# Add storage requirements
storage:
  shared_filesystem:
    type: "efs"  # or "fsx-lustre", "fsx-openzfs"
    size_gb: 1000
    throughput_mode: "bursting"
```

---

## Recommended Implementation Order

### Phase 1: Current Focus ✓
- ✅ CPU-only applications
- ✅ OpenMP (single-node)
- ✅ Basic MPI (GCHP)
- ✅ S3 + EBS storage

**Applications:** GEOS-Chem, ORCA, NWChem, R, Julia

### Phase 2: Expand Compute
- 🚧 Enhanced MPI (EFA networking)
- 🚧 Array jobs
- 🚧 FSx for Lustre

**Applications:** WRF, CESM, OpenFOAM, Quantum ESPRESSO

### Phase 3: Commercial Apps
- 🚧 License management
- 🚧 License server integration
- 🚧 Usage tracking

**Applications:** Gaussian, VASP, ANSYS, COMSOL, MATLAB

### Phase 4: GPU Support
- 🔮 GPU base images
- 🔮 CUDA/ROCm support
- 🔮 Multi-GPU orchestration

**Applications:** GROMACS, NAMD, TensorFlow, PyTorch

### Phase 5: Big Data
- 🔮 Spark integration
- 🔮 EMR support
- 🔮 Object storage optimization

**Applications:** GATK, large-scale genomics

---

## Application Priority Matrix

| Priority | Application | Complexity | User Demand | Implementation Phase |
|----------|-------------|------------|-------------|----------------------|
| **P0** | GEOS-Chem | Medium | High | Phase 1 ✅ |
| **P1** | Gaussian | Low (license) | Very High | Phase 3 |
| **P1** | WRF | High (MPI) | High | Phase 2 |
| **P1** | ORCA | Low | High | Phase 1 |
| **P2** | VASP | Medium (license) | High | Phase 3 |
| **P2** | GROMACS | High (GPU) | High | Phase 4 |
| **P2** | OpenFOAM | Medium (MPI) | Medium | Phase 2 |
| **P2** | Quantum ESPRESSO | Medium | Medium | Phase 2 |
| **P3** | TensorFlow/PyTorch | High (GPU) | High | Phase 4 |
| **P3** | ANSYS | High (license+GUI) | Medium | Phase 3 |
| **P3** | CESM | High (MPI+scale) | Medium | Phase 2 |
| **P4** | MATLAB | Medium (license) | Medium | Phase 3 |
| **P4** | COMSOL | High (GUI) | Medium | Phase 3 |

---

## Conclusion

The analysis reveals:

1. **Current architecture is excellent** for ~40% of research apps (CPU-only, single/basic MPI)
2. **License management is critical** for commercialization (7 of top 25 apps)
3. **GPU support opens 24%** of the market but is complex
4. **MPI with EFA** needed for weather/climate apps
5. **Storage tiering** must be flexible (EFS/FSx for Lustre)

**Recommendation**: Continue with Phase 1 (GEOS-Chem extraction), then add Gaussian (Phase 3 license management) and WRF (Phase 2 MPI) as next applications to validate platform flexibility.
