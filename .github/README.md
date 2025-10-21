# GitHub Configuration for AWS HPC Platform

This directory contains GitHub-specific configuration for project management, issue tracking, and automation.

## Structure

```
.github/
├── ISSUE_TEMPLATE/           # Issue templates for bugs, features, tech debt
│   ├── config.yml           # Template configuration
│   ├── bug_report.yml       # Bug report template
│   ├── feature_request.yml  # Feature request template
│   └── technical_debt.yml   # Technical debt template
├── workflows/               # GitHub Actions workflows
│   ├── docs.yml            # Documentation build and deployment
│   ├── labels.yml          # Label synchronization
│   └── project-automation.yml # Auto-add issues/PRs to project board
├── pull_request_template.md # PR template
├── labels.yml              # Label definitions
└── README.md               # This file
```

## Issue Templates

### Bug Report (`bug_report.yml`)
For reporting bugs and unexpected behavior. Collects:
- Affected area (platform, application, infrastructure, etc.)
- Application name (if applicable)
- Component (job submission, container build, etc.)
- Version information
- OS and AWS region
- Processor architecture (if relevant)
- Reproduction steps
- Expected vs actual behavior
- Logs and configuration (sanitized)

### Feature Request (`feature_request.yml`)
For suggesting new features. Collects:
- Target persona (scientist, engineer, admin, etc.)
- Affected areas
- Computational use case
- Problem statement
- Proposed solution
- Alternatives considered
- Workflow impact
- Priority level
- Architecture preferences

### Technical Debt (`technical_debt.yml`)
For code improvements and refactoring. Collects:
- Code area (pkg/, applications/, infrastructure/, etc.)
- Type of improvement (refactoring, performance, security, etc.)
- Current state and issues
- Proposed improvement
- Impact assessment
- Testing strategy
- Breaking change status
- Application impact

### Template Configuration (`config.yml`)
- Disables blank issues
- Provides links to:
  - Q&A Discussions
  - Ideas & Feature Brainstorming
  - Documentation (GitHub Pages)
  - Security reporting

## Labels (`labels.yml`)

Comprehensive label system organized hierarchically:

### Type Labels
- `bug`, `enhancement`, `documentation`, `technical-debt`, `question`

### Priority Labels
- `priority: critical/high/medium/low`

### Area Labels
- `area: platform`, `area: applications`, `area: base-images`, `area: infrastructure`
- `area: cli`, `area: web`, `area: containers`, `area: job-management`
- `area: cost`, `area: license`, `area: tests`, `area: build`, `area: docs`

### Application Labels
- `app: geos-chem`, `app: gaussian`, `app: wrf`, `app: vasp`, `app: template`

### Persona Labels
- `persona: computational-scientist`, `persona: research-lab`, `persona: hpc-engineer`
- `persona: graduate-student`, `persona: it-admin`, `persona: devops-engineer`, `persona: data-scientist`

### Use Case Labels
- `use-case: atmospheric-modeling`, `use-case: molecular-dynamics`, `use-case: quantum-chemistry`
- `use-case: machine-learning`, `use-case: genomics`, `use-case: cfd`
- `use-case: materials-science`, `use-case: parallel-computing`
- `use-case: cost-optimization`, `use-case: benchmarking`

### Status Labels
- `triage`, `needs-info`, `blocked`, `ready`, `in-progress`, `in-review`, `awaiting-merge`

### Resolution Labels
- `duplicate`, `wontfix`, `invalid`, `works-as-designed`

### Special Labels
- `good first issue`, `help wanted`, `breaking-change`, `security`, `performance`, `dependencies`

### Phase Labels (Roadmap)
- `phase: 0.1-foundation`, `phase: 0.2-multi-app`, `phase: 0.3-optimization`
- `phase: 0.4-enterprise`, `phase: 1.0-production`

### AWS-Specific Labels
- `aws: batch`, `aws: ec2`, `aws: ecr`, `aws: s3`, `aws: efs`, `aws: fsx`
- `aws: vpc`, `aws: parallelcluster`, `aws: spot`, `aws: iam`, `aws: gpu`

### Architecture Labels
- `arch: graviton`, `arch: amd`, `arch: intel`, `arch: multi-arch`

## Workflows

### Documentation (`docs.yml`)
- **Trigger**: Push to main or PR changing docs/ or mkdocs.yml
- **Check job** (PRs): Validates documentation builds successfully
- **Deploy job** (main): Builds and deploys to GitHub Pages
- Uses MkDocs with Material theme

### Labels (`labels.yml`)
- **Trigger**: Push to main changing labels.yml, or manual dispatch
- **Action**: Syncs labels from `.github/labels.yml` to repository
- Uses `micnncim/action-label-syncer`

### Project Automation (`project-automation.yml`)
- **Trigger**: Issue or PR opened/reopened
- **Action**: Automatically adds to GitHub Project board
- Requires project URL to be updated with your project number

## Pull Request Template

Comprehensive PR template with sections for:
- Description and related issues
- Type of change (bug, feature, tech debt, docs, breaking, security)
- Computational use case impact
- Testing checklist (manual, unit, integration, e2e, architecture-specific)
- Code review checklist
- Documentation updates
- Performance impact
- Cost impact
- Breaking changes and migration path
- Components affected
- Architecture impact

## Setup Instructions

### 1. Enable GitHub Pages
1. Go to repository Settings → Pages
2. Source: GitHub Actions
3. Workflow will deploy automatically on push to main

### 2. Create GitHub Project (Optional)
1. Go to your profile → Projects → New project
2. Note the project number from URL
3. Update `project-automation.yml` with your project URL

### 3. Enable Discussions (Optional)
1. Go to repository Settings → Features
2. Check "Discussions"
3. Create categories: Q&A, Ideas, General

### 4. Sync Labels
After first push to main:
```bash
# Labels will sync automatically via workflow
# Or manually trigger:
gh workflow run labels.yml
```

### 5. Test Issue Templates
1. Go to Issues → New issue
2. Verify templates appear correctly
3. Submit a test issue to verify all fields work

## Customization

### Update Repository Links
Replace `scttfrdmn/aws-hpc` in:
- `config.yml` (issue template links)
- `mkdocs.yml` (documentation site URL)
- `project-automation.yml` (project URL)

### Add New Labels
Edit `.github/labels.yml` and push to main. Workflow will sync automatically.

### Modify Issue Templates
Edit YAML files in `ISSUE_TEMPLATE/` directory. Changes take effect immediately.

### Update Workflows
Edit workflow files in `workflows/` directory. Test in a feature branch before merging.

## Best Practices

### For Maintainers
1. **Triage regularly**: Review new issues and add appropriate labels
2. **Use milestones**: Link issues to release milestones (v0.1.0, v0.2.0)
3. **Reference roadmap**: Tag issues with phase labels (phase: 0.1-foundation)
4. **Close with context**: Always explain why when closing issues

### For Contributors
1. **Search first**: Check if issue already exists before creating
2. **Use templates**: Fill out all required fields
3. **Be specific**: Provide clear reproduction steps and examples
4. **Sanitize data**: Remove account IDs, credentials, and sensitive info
5. **Link PRs**: Reference related issues with `Closes #123` in PR description

## Related Documentation

- [ROADMAP.md](../ROADMAP.md) - Product roadmap aligned with phase labels
- [PERSONAS.md](../docs/PERSONAS.md) - User personas referenced in issue templates
- [CONTRIBUTING.md](../CONTRIBUTING.md) - Contribution guidelines
- [GitHub Pages Site](https://scttfrdmn.github.io/aws-hpc) - Published documentation

## Questions?

- **Setup issues**: [Open a discussion](https://github.com/scttfrdmn/aws-hpc/discussions/new?category=q-a)
- **Template improvements**: [Feature request](https://github.com/scttfrdmn/aws-hpc/issues/new?template=feature_request.yml)
- **Workflow problems**: [Bug report](https://github.com/scttfrdmn/aws-hpc/issues/new?template=bug_report.yml)
