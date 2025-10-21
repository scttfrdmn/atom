# ATOM Rebranding Summary

**Date**: October 21, 2025
**Repository**: https://github.com/scttfrdmn/atom
**Website**: https://atomhpc.io

## ⚛️ New Identity

### Name
**ATOM** - Automated Toolkit for Optimized Modeling

### Branding
- **Tagline**: *Cloud-native HPC made simple*
- **Motto**: *The fundamental unit of cloud computing*
- **Logo**: ⚛️ (atom symbol)
- **Domain**: atomhpc.io
- **Focus**: Cloud-native high-performance computing platform

### Acronym Pattern
Following the ORCA model (Orchestration for Research Cloud Access):
- **Technical**: Automated Toolkit for Optimized Modeling (describes function)
- **Memorable**: ATOM (real scientific concept)
- **Visual**: Atom symbol ⚛️ (fundamental building block)
- **Metaphor**: Atoms are fundamental units; ATOM is fundamental unit of cloud computing

## 📦 What Changed

### Repository
- ✅ Renamed from `aws-hpc` to `atom`
- ✅ GitHub automatic redirect: `scttfrdmn/aws-hpc` → `scttfrdmn/atom`
- ✅ Git remote updated
- ✅ Repository description updated
- ✅ Topics added: hpc, cloud-native, scientific-computing, aws, research-computing, batch-computing, containers

### Documentation
- ✅ README.md completely rebranded
- ✅ docs/index.md updated with ATOM branding
- ✅ mkdocs.yml configured for atomhpc.io
- ✅ All command examples changed from `aws-hpc` to `atom`
- ✅ Added "cloud-native HPC" messaging throughout

### GitHub Infrastructure
- ✅ Issue templates updated (bug report, feature request, technical debt)
- ✅ Issue template config URLs updated
- ✅ Pull request template (already generic)
- ✅ Labels (already configured, no project-specific names)

### Custom Domain
- ✅ Domain purchased: atomhpc.io (and lenslab.io)
- ✅ DNS configured
- ✅ CNAME file added to docs/
- ✅ GitHub Pages deployment configured
- ⏳ Waiting for DNS propagation (5-60 minutes typical)

## 🌐 The Research Computing Ecosystem

ATOM now joins a cohesive family of research computing tools:

```
┌─────────────────────────────────────────────────────────────┐
│           Research Computing Platform Ecosystem             │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  🧫 petri.io                                                │
│     Account & Budget Management                             │
│     └─> Creates AWS accounts for research                   │
│                                                              │
│  🖥️  cloudworkspaces.io                                     │
│     Interactive Research Workstations                       │
│     └─> Pre-configured ML, R, Python environments           │
│                                                              │
│  🔬 lenslab.io (LENS)                                       │
│     Lab Environment Notebook System                         │
│     └─> Jupyter, RStudio, VSCode for researchers            │
│                                                              │
│  ⚛️  atomhpc.io (ATOM)                                      │
│     Automated Toolkit for Optimized Modeling               │
│     └─> Cloud-native HPC batch computing                    │
│                                                              │
│  🚢 cargoship                                                │
│     Enterprise Data Archiving                               │
│     └─> S3-optimized long-term storage                      │
│                                                              │
│  🐋 orcapod.io (ORCA)                                       │
│     Orchestration for Research Cloud Access                 │
│     └─> Kubernetes → AWS burst computing                    │
│                                                              │
└─────────────────────────────────────────────────────────────┘
```

### Integration Flow
```
University → petri (request account) → Account created
                                             ↓
                         ┌───────────────────┼──────────────────┐
                         ↓                   ↓                  ↓
              cloudworkspaces            lenslab           atomhpc
                (interactive)        (development)    (batch HPC)
                         ↓                   ↓                  ↓
                         └───────────────────┴────────> cargoship
                                       (archive results)
```

## 🎯 Why ATOM?

### Technical Rationale
1. **Avoids AWS Trademark**: No "aws-" prefix
2. **Platform Agnostic**: Works with any cloud (Azure, GCP, on-prem)
3. **Memorable Acronym**: Real word that's also a meaningful acronym
4. **Scientific Credibility**: Atom = fundamental unit (perfect for HPC!)
5. **Visual Branding**: ⚛️ symbol is universally recognized

### Naming Strategy
- **Real word**: ATOM (fundamental particle)
- **Acronym**: Automated Toolkit for Optimized Modeling
- **Domain combo**: atomhpc.io (ATOM + HPC descriptor)
- **Pattern**: Following orcapod.io model

### Brand Positioning
- **Focus**: Cloud-native HPC (not generic cloud management)
- **Audience**: Computational scientists, research labs, HPC engineers
- **Differentiator**: Application-agnostic platform for scientific computing
- **Value**: Makes HPC simple without traditional cluster management

## 📋 Remaining Tasks

### Immediate (When Development Starts)
- [ ] Update Go module path: `github.com/scttfrdmn/atom`
- [ ] Update CLI binary name in build scripts
- [ ] Create atom icon/logo assets (optional but recommended)

### Future Enhancements
- [ ] Create visual logo (atom symbol with branding)
- [ ] Add atom logo to documentation
- [ ] Consider atom-themed visual elements in docs
- [ ] Update any external references (if they exist)

### Domain Configuration
- ✅ CNAME file committed
- ✅ GitHub Pages deployment configured
- ⏳ Waiting for DNS propagation
- [ ] Verify HTTPS certificate (automatic via GitHub)
- [ ] Test https://atomhpc.io once DNS propagates

## 🔗 Important Links

### Production
- **Repository**: https://github.com/scttfrdmn/atom
- **Website**: https://atomhpc.io (propagating)
- **Documentation**: https://atomhpc.io
- **Issues**: https://github.com/scttfrdmn/atom/issues
- **Discussions**: https://github.com/scttfrdmn/atom/discussions

### Legacy (Redirects)
- **Old Repo**: https://github.com/scttfrdmn/aws-hpc (→ atom)
- **Old Docs**: https://scttfrdmn.github.io/aws-hpc/ (→ atom)

## 📝 Commit History

### Key Commits
1. `a8580be` - Initial commit: AWS HPC Platform foundation
2. `6f16073` - Fix documentation build: simplify nav to existing pages
3. `220f7b5` - Fix broken links in getting-started documentation
4. `1a6c474` - Remove strict mode from documentation builds
5. `fcc5d41` - Trigger documentation build after GitHub Pages enablement
6. `7494728` - Trigger docs build
7. `e419528` - Rebrand to ATOM: Automated Toolkit for Optimized Modeling
8. `f970bfe` - Add custom domain atomhpc.io for GitHub Pages

## 🎉 Success Metrics

- ✅ Repository renamed successfully
- ✅ All documentation updated
- ✅ Custom domain configured
- ✅ GitHub Pages deploying
- ✅ Automatic redirects working
- ✅ Dual versioning system in place
- ✅ Issue templates configured
- ✅ Comprehensive personas documented
- ✅ Roadmap through v1.0 defined

## 📞 Next Steps

Once DNS propagates (check in 5-60 minutes):
1. Visit https://atomhpc.io to verify site is live
2. Verify HTTPS certificate is active
3. Test all navigation and links
4. Share new URL with collaborators

To check DNS propagation:
```bash
# Check DNS resolution
dig atomhpc.io

# Check website availability
curl -I https://atomhpc.io

# Check GitHub Pages status
gh api repos/scttfrdmn/atom/pages
```

---

**Status**: ✅ Rebranding Complete
**Domain**: ⏳ Propagating
**Next Milestone**: Begin platform development with ATOM branding
