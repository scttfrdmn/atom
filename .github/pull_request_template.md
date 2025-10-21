## Description

<!-- Provide a clear and concise description of your changes -->

## Related Issues

<!-- Link to related issues using #issue_number or "Closes #123" -->

Closes #

## Type of Change

<!-- Check all that apply -->

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Technical debt (code refactoring, optimization, or cleanup)
- [ ] Documentation update
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Security fix

## Computational Use Case Impact

<!-- Which computational use cases does this affect? Check all that apply -->

- [ ] Atmospheric Modeling
- [ ] Molecular Dynamics
- [ ] Quantum Chemistry
- [ ] Machine Learning / AI
- [ ] Genomics / Bioinformatics
- [ ] Computational Fluid Dynamics
- [ ] Materials Science
- [ ] Parallel Computing (MPI)
- [ ] Cost Optimization
- [ ] Performance Benchmarking
- [ ] N/A (infrastructure/build/tests)

## Testing

<!-- Describe the testing you performed -->

- [ ] Manual testing performed
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] End-to-end tests with real AWS Batch jobs (if applicable)
- [ ] Container builds tested (if applicable)
- [ ] `make test` passes locally
- [ ] `make lint` passes locally
- [ ] Tested on multiple platforms (if applicable)
- [ ] Tested on multiple architectures (ARM/AMD/Intel if applicable)

**Test Details:**
```
Describe manual testing steps here
```

## Code Review Checklist

<!-- Verify these before requesting review -->

- [ ] My code follows the project's style guidelines
- [ ] I have performed a self-review of my code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings or errors
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] Any dependent changes have been merged and published

## Documentation Updates

- [ ] Updated CHANGELOG.md
- [ ] Updated relevant documentation (docs/ or README)
- [ ] Updated application README (if application-specific change)
- [ ] Added/updated code comments
- [ ] Updated mkdocs configuration (if adding new pages)

## Performance Impact

<!-- Describe any performance implications -->

- [ ] No significant performance impact
- [ ] Performance improvement (describe below)
- [ ] Potential performance impact (describe below)

**Performance Notes:**


## Cost Impact

<!-- Describe any AWS cost implications -->

- [ ] No cost impact
- [ ] Cost reduction (describe below)
- [ ] Potential cost increase (describe below)

**Cost Notes:**


## Breaking Changes

<!-- If this includes breaking changes, describe migration path -->

- [ ] No breaking changes
- [ ] Breaking changes included (describe below)

**Migration Path:**


## Screenshots / Demo

<!-- Add screenshots, terminal output, or demo videos if applicable -->


## Reviewer Notes

<!-- Any additional context for reviewers -->


## Components Affected

<!-- Check all that apply -->

- [ ] Platform core (pkg/)
- [ ] Applications
- [ ] Base images
- [ ] Infrastructure (CDK/Terraform)
- [ ] CLI (aws-hpc)
- [ ] Web dashboard
- [ ] Container build system
- [ ] Job management
- [ ] Cost tracking
- [ ] License management

## Architecture Impact

<!-- If this affects specific processor architectures, note here -->

- [ ] All architectures
- [ ] ARM Graviton specific
- [ ] AMD EPYC specific
- [ ] Intel specific
- [ ] GPU specific
