# AWS HPC Platform - User Personas

This document defines the primary user personas for the AWS HPC Platform. These personas help guide feature prioritization, design decisions, and documentation structure.

---

## 1. Dr. Sarah Chen - Computational Atmospheric Scientist

### Background
- **Role**: Research scientist at a university
- **Domain**: Atmospheric chemistry and climate modeling
- **Experience**: Expert in GEOS-Chem, moderate AWS knowledge
- **Team Size**: Solo researcher with occasional graduate student collaborators

### Goals
- Run GEOS-Chem Classic simulations cost-effectively
- Quickly iterate on model configurations
- Compare results across different processor architectures
- Minimize cloud infrastructure complexity

### Pain Points
- **Current State**: Runs simulations on aging departmental cluster with wait times
- **Challenges**:
  - Limited compute resources during peak times
  - Manual cost tracking is tedious
  - Unclear which AWS instance types are most cost-effective
  - Container configuration feels like overkill for simple jobs

### Use Cases
- Quick sensitivity studies (hours of simulation time)
- Long-running production simulations (days to weeks)
- Architecture benchmarking to optimize costs
- Teaching graduate students to run simulations

### Success Criteria
- Can submit a job in < 5 minutes without AWS console
- Clear cost estimates before running
- Automatic selection of cheapest architecture that meets performance needs
- Easy debugging when jobs fail

### Preferred Interaction
```bash
# Simple, declarative job submission
aws-hpc job submit geos-chem \
  --config benchmark.yaml \
  --architecture auto \  # Let platform choose cheapest
  --budget 50            # Fail if estimated cost > $50
```

### Quote
> "I want to focus on my science, not on cloud infrastructure. Just tell me which instance type to use and how much it will cost."

---

## 2. Dr. James Rodriguez - Computational Chemistry Lab PI

### Background
- **Role**: Principal Investigator running a 15-person research group
- **Domain**: Quantum chemistry (Gaussian, ORCA)
- **Experience**: Limited AWS knowledge, strong HPC background
- **Team Size**: 5 postdocs, 8 graduate students, 2 undergrads

### Goals
- Deploy Gaussian for entire lab with consistent environments
- Track costs per student/project for grant reporting
- Ensure license compliance (Gaussian site license)
- Enable students to run jobs independently

### Pain Points
- **Current State**: On-premises HPC cluster at 90% capacity
- **Challenges**:
  - Students need to learn AWS IAM and infrastructure
  - Gaussian license management is manual
  - No visibility into which project/student is spending what
  - Concerns about data security and compliance

### Use Cases
- Multi-user lab deployment with cost allocation
- License server integration (FlexLM)
- Student job submission without AWS console access
- Monthly cost reports broken down by user and project

### Success Criteria
- One-time infrastructure deployment for entire lab
- Students can submit jobs via simple CLI (no AWS knowledge required)
- Automatic license checkout/checkin
- Cost dashboard showing spending by student and grant
- Data encryption at rest and in transit

### Preferred Interaction
```bash
# PI: One-time lab setup
aws-hpc lab deploy gaussian \
  --license-server license.chem.university.edu \
  --users students.csv \
  --cost-tracking-by user,project

# Students: Simple job submission
aws-hpc job submit gaussian \
  --input molecule.com \
  --project "NSF-Grant-2024" \
  --user "student1@university.edu"
```

### Quote
> "I need to give my students access to compute without teaching them cloud engineering. And I need to know who's spending what on which grants."

---

## 3. Alex Thompson - HPC Engineer

### Background
- **Role**: HPC engineer at a national research facility
- **Domain**: Infrastructure for diverse scientific applications
- **Experience**: Expert in HPC, AWS, containers, and infrastructure-as-code
- **Team Size**: Supports 200+ researchers across disciplines

### Goals
- Deploy platform that supports multiple scientific applications
- Customize infrastructure for specific performance requirements
- Integrate with existing authentication systems (LDAP, AD)
- Maintain security and compliance standards

### Pain Points
- **Current State**: Managing multiple bespoke cloud deployments
- **Challenges**:
  - Each research group has slightly different requirements
  - No standardization across applications
  - Security reviews for every new deployment
  - Difficulty keeping container images updated

### Use Cases
- Multi-application platform deployment (GEOS-Chem, WRF, VASP, Gaussian)
- Custom VPC and networking configuration
- Integration with institutional SSO
- Centralized monitoring and cost allocation
- Automated security patching

### Success Criteria
- Single platform supporting 5+ scientific applications
- Infrastructure-as-code for reproducibility
- Comprehensive monitoring and logging
- Compliance with institutional security policies
- Easy onboarding of new applications

### Preferred Interaction
```typescript
// Infrastructure as Code (CDK)
const hpcPlatform = new HpcPlatformStack(app, 'ResearchFacility', {
  vpc: existingVpc,
  applications: ['geos-chem', 'wrf', 'gaussian', 'vasp'],
  authentication: {
    type: 'ldap',
    server: 'ldap.facility.gov'
  },
  monitoring: {
    cloudwatch: true,
    datadog: true
  }
});
```

### Quote
> "I need a platform I can trust, customize, and deploy repeatedly. Show me the CDK code, not a web console."

---

## 4. Maria Santos - Graduate Student

### Background
- **Role**: PhD student in atmospheric science
- **Domain**: Climate modeling with GEOS-Chem High Performance (GCHP)
- **Experience**: Intermediate Linux, beginner AWS
- **Team Size**: Part of 8-person research group

### Goals
- Run multi-node MPI simulations (GCHP)
- Learn cloud computing for future career
- Work within limited research budget
- Complete dissertation simulations on time

### Pain Points
- **Current State**: Limited to small GCHP simulations on university cluster
- **Challenges**:
  - Need to run 100+ sensitivity experiments
  - Cluster queue times are 3-7 days
  - Budget constraints ($3000 for entire dissertation)
  - Afraid of accidentally spending too much on AWS

### Use Cases
- Array jobs (100 variations of same simulation)
- Multi-node MPI jobs (24-96 cores)
- Cost-optimized spot instances
- Budget alerts and auto-stop

### Success Criteria
- Can run 10x more simulations than on-premises
- Budget safety net (hard limits, alerts)
- Clear documentation for cloud beginners
- Faster turnaround than on-premises cluster

### Preferred Interaction
```bash
# Array job for sensitivity study
aws-hpc job submit-array gchp \
  --config-template sensitivity.yaml \
  --parameter-file params.csv \  # 100 rows
  --budget 500 \
  --alert-threshold 400  # Email at $400
```

### Quote
> "I have 100 simulations to run and $3000 for my entire dissertation. Help me not blow my budget on accident."

---

## 5. Dr. Priya Patel - DevOps Engineer for Research Cloud

### Background
- **Role**: DevOps engineer supporting research computing
- **Domain**: Cloud infrastructure, CI/CD, security
- **Experience**: Expert AWS, Kubernetes, Terraform, security
- **Team Size**: 3-person DevOps team supporting 500+ researchers

### Goals
- Automate deployment and updates
- Implement security best practices
- Monitor costs across research groups
- Maintain compliance with data protection regulations

### Pain Points
- **Current State**: Researchers deploy one-off solutions that become tech debt
- **Challenges**:
  - Difficult to audit what's running in AWS accounts
  - Researchers open security groups to 0.0.0.0/0
  - No standardization for container images
  - Cost overruns are frequent

### Use Cases
- Centralized platform deployment via CI/CD
- Automated security scanning of containers
- Cost anomaly detection and alerts
- Compliance reporting
- Infrastructure drift detection

### Success Criteria
- All HPC workloads use standardized platform
- Automated vulnerability scanning
- Cost visibility and forecasting
- Infrastructure changes go through peer review
- Zero manual deployments

### Preferred Interaction
```yaml
# .github/workflows/deploy-hpc.yml
name: Deploy HPC Platform
on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: aws-actions/configure-aws-credentials@v4
      - run: |
          aws-hpc platform deploy \
            --config production.yaml \
            --security-scan \
            --cost-estimate \
            --require-approval
```

### Quote
> "Researchers shouldn't have AWS console access. Give them a safe, compliant platform with guard rails."

---

## 6. Dr. Michael Zhang - Bioinformatics Data Scientist

### Background
- **Role**: Data scientist at biotech company
- **Domain**: Genomics, machine learning on biological data
- **Experience**: Strong Python/R, moderate cloud experience
- **Team Size**: 4-person computational biology team

### Goals
- Run embarrassingly parallel genomics pipelines
- GPU acceleration for ML model training
- S3 integration for large datasets (multi-TB)
- Jupyter notebooks for interactive analysis

### Pain Points
- **Current State**: Mix of local workstations and ad-hoc AWS
- **Challenges**:
  - Datasets too large for local storage
  - Need hundreds of parallel jobs for batch processing
  - GPU instances are expensive (want spot)
  - Want to prototype in Jupyter, scale to Batch

### Use Cases
- Array jobs processing 1000s of samples
- GPU-accelerated deep learning
- Interactive development (Jupyter) + production batch jobs
- Data pipeline orchestration

### Success Criteria
- Seamless S3 data access
- GPU spot instances with fault tolerance
- Can submit jobs from Jupyter notebooks
- Pipeline orchestration (Step Functions or similar)

### Preferred Interaction
```python
# Python SDK for job submission from notebooks
import awshpc

job = awshpc.Job(
    application="custom-genomics",
    container="genomics-pipeline:latest",
    input_data="s3://genomics-data/raw/",
    output_data="s3://genomics-data/processed/",
    instance_type="g5.2xlarge",
    spot=True,
    retry_on_spot_interruption=3
)

job_id = job.submit()
results = job.wait()  # Block until complete
```

### Quote
> "I need to process thousands of samples in parallel, use GPUs cost-effectively, and integrate with my existing data in S3."

---

## 7. Tom Wilson - IT Administrator (University Research Computing)

### Background
- **Role**: IT admin managing university research computing services
- **Domain**: System administration, user support, budgets
- **Experience**: Strong Linux/HPC, basic AWS
- **Team Size**: 2-person IT team supporting 300 faculty researchers

### Goals
- Provide self-service cloud access to faculty
- Minimize support burden
- Track and allocate AWS costs to departments
- Maintain institutional compliance

### Pain Points
- **Current State**: Researchers request AWS accounts and spend without oversight
- **Challenges**:
  - No central cost visibility
  - Security incidents from misconfigured resources
  - Support requests for basic AWS usage
  - Budget overruns and surprise bills

### Use Cases
- Centralized platform for all research computing
- Department-level cost allocation
- Self-service with guard rails
- Centralized monitoring and alerting
- Chargeback to departments

### Success Criteria
- 90% reduction in AWS support tickets
- Central cost dashboard
- Departmental chargeback reports
- Compliance with university security policy
- Researchers can self-serve without opening tickets

### Preferred Interaction
```bash
# Admin: Set up department budget
aws-hpc admin create-department chemistry \
  --budget 10000 \
  --alert-threshold 8000 \
  --billing-code "CHEM-AWS-2025"

# Admin: Grant faculty access
aws-hpc admin add-user "dr.chen@university.edu" \
  --department chemistry \
  --role researcher

# Monthly report
aws-hpc admin report costs --by department --month 2025-10
```

### Quote
> "I need researchers to be productive without creating security incidents or budget emergencies."

---

## Persona Summary Matrix

| Persona | Primary Goal | AWS Expertise | HPC Expertise | Main Concern | Preferred Interface |
|---------|--------------|---------------|---------------|--------------|---------------------|
| Sarah (Scientist) | Run simulations efficiently | Medium | High | Cost & simplicity | CLI |
| James (PI) | Lab-wide deployment | Low | High | Cost tracking & licenses | CLI + Dashboard |
| Alex (HPC Engineer) | Multi-app platform | Expert | Expert | Customization & security | IaC (CDK) |
| Maria (Grad Student) | Budget-conscious research | Beginner | Medium | Budget safety | CLI |
| Priya (DevOps) | Automation & compliance | Expert | Medium | Security & standards | CI/CD + IaC |
| Michael (Data Scientist) | Parallel pipelines | Medium | Low | GPU costs & S3 integration | Python SDK |
| Tom (IT Admin) | Central governance | Basic | Medium | Cost visibility & support | Admin CLI + Dashboard |

---

## Using These Personas

### Feature Prioritization
When evaluating new features, ask:
- Which persona(s) benefit most?
- Does this align with their primary goals?
- Does it address their pain points?

### Design Decisions
Consider:
- **For Sarah & Maria**: Optimize for simplicity and clear cost feedback
- **For James & Tom**: Build cost allocation and multi-user management
- **For Alex & Priya**: Provide infrastructure-as-code and customization
- **For Michael**: Focus on data integration and GPU optimization

### Documentation Structure
- **Getting Started**: Target Sarah, Maria (beginners)
- **Advanced Deployment**: Target Alex, Priya (experts)
- **Lab/Multi-user Setup**: Target James, Tom (administrators)
- **API/SDK**: Target Michael, Priya (programmatic usage)

### Success Metrics
- **Sarah & Maria**: Time from idea to results, cost per simulation
- **James & Tom**: Cost visibility, user satisfaction, support burden
- **Alex & Priya**: Deployment repeatability, security compliance
- **Michael**: Data throughput, GPU utilization, cost efficiency

---

## Feedback and Evolution

These personas should evolve based on real user feedback. To suggest updates:
1. Open an issue with label `persona: [name]`
2. Provide specific examples from real usage
3. Link to related feature requests or bugs

Last Updated: October 2025
