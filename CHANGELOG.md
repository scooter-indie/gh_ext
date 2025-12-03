# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Removed
- **Epic 2: Project Templates & Creation** - Removed as redundant with native `gh project` commands
  - `gh pmu project create --from-project` → Use `gh project copy` instead
  - `gh pmu project export` → Use `gh project field-list` + `gh project view` instead
  - Template-based creation blocked by GitHub API (no view creation API exists)
- **Epic 4: Template Ecosystem** - Removed as dependent on Epic 2
- Removed `cmd/project.go` and related template code
- Removed `internal/template/` package
- Removed project copy/export API functions

### Changed
- Updated README.md with current feature set and project status
- Updated PRD to reflect removed epics

## [0.3.0] - 2025-12-02

### Added - Sprint 3: Batch Operations
- **Intake command** (`gh pmu intake`)
  - Find open issues not yet added to the project
  - `--dry-run` to preview what would be added
  - `--apply` to add issues to project with default fields
  - Respects repository filter from config
- **Triage command** (`gh pmu triage`)
  - Bulk update issues matching configurable rules
  - Triage configs defined in `.gh-pmu.yml`
  - Supports applying labels, status, priority changes
  - `--interactive` mode for per-issue confirmation
  - `--dry-run` to preview changes
- **Split command** (`gh pmu split`)
  - Create sub-issues from checklist in issue body (`--from body`)
  - Create sub-issues from external file (`--from file.md`)
  - Create sub-issues from command arguments
  - Auto-links created issues as sub-issues of parent
- `GetRepositoryIssues` API query for fetching issues by state
- `AddLabelToIssue` API mutation stub

### Changed
- Sprint 3 backlog updated with completed status
- Added sprint-3-retro.md and sprint-3-summary.md

## [0.2.0] - 2025-12-02

### Added - Sprint 2: Core Commands & Sub-Issues
- **View command** (`gh pmu view <issue>`)
  - Display issue details with all project field values
  - Show sub-issues if any exist
  - Show parent issue if this is a sub-issue
  - `--json` output format support
- **Create command** (`gh pmu create`)
  - Create issues with project fields pre-populated
  - `--title`, `--body` flags for non-interactive creation
  - `--status`, `--priority` to set project field values
  - `--label` to apply labels
  - Automatically adds issue to configured project
- **Move command** (`gh pmu move <issue>`)
  - Update issue project fields from command line
  - `--status`, `--priority` flags
  - Supports field aliases from config
- **Sub-issue commands** (`gh pmu sub`)
  - `sub add <parent> <child>` - Link existing issue as sub-issue
  - `sub create --parent <id> --title <title>` - Create new sub-issue
  - `sub list <parent>` - List sub-issues with completion count
  - `sub remove <parent> <child>` - Unlink sub-issue
- `CreateIssue` API mutation with label support
- `AddIssueToProject` API mutation
- `SetProjectItemField` API mutation for single-select, text, and number fields
- `AddSubIssue` and `RemoveSubIssue` API mutations
- `GetSubIssues` and `GetParentIssue` API queries

### Changed
- Renamed command from `gh pm` to `gh pmu` to avoid conflict with existing extension
- Updated module path to `github.com/scooter-indie/gh-pmu`
- Config filename changed to `.gh-pmu.yml`
- All documentation updated with new command name

## [0.1.0] - 2025-12-01

### Added - Sprint 1: Foundation
- **Project scaffolding**
  - Go module with Cobra CLI framework
  - Makefile with build, test, lint targets
  - GitHub Actions for CI/CD
  - .goreleaser.yml for releases
- **Configuration package** (`internal/config`)
  - Load `.gh-pmu.yml` with Viper
  - Validate required fields
  - Support field aliases (e.g., `backlog` → `Backlog`)
  - Cache project metadata from GitHub API
- **GitHub API client** (`internal/api`)
  - GraphQL client using go-gh library
  - Support for `sub_issues` and `issue_types` feature headers
  - `GetProject` query for user and organization projects
  - `GetProjectFields` query with single-select options
  - `GetProjectItems` query with field values
  - `GetIssue` query with full metadata
- **Init command** (`gh pmu init`)
  - Interactive project configuration setup
  - Auto-detect current repository from git remote
  - Validate project exists before saving
  - Cache project field metadata
- **List command** (`gh pmu list`)
  - Display issues with project field values
  - `--status`, `--priority` filters
  - `--json` output format
  - Repository filtering from config

### Attribution
- Added attribution to [@yahsan2](https://github.com/yahsan2) for original [gh-pm](https://github.com/yahsan2/gh-pm) and [gh-sub-issue](https://github.com/yahsan2/gh-sub-issue) projects

## [0.0.1] - 2025-11-30

### Added - Project Setup
- Initial repository setup
- README.md with project overview
- GitHub project integration configuration
- Agile PRD and product backlog
- Code integration inventory from source projects
- PROPOSAL.md with project vision

---

## Sprint Summary

| Sprint | Points | Focus |
|--------|--------|-------|
| Sprint 1 | 28 | Foundation, init, list |
| Sprint 2 | 24 | view, create, move, sub-issues |
| Sprint 3 | 21 | intake, triage, split |
| Sprint 4 | - | Epic 2 removed (redundant) |

**Total Completed:** 73 story points (Epic 1 complete)

## API Limitations Discovered

### GitHub Projects V2 API
- **No `createProjectV2View` mutation** - Views cannot be created programmatically
- **Status field reserved** - New projects have a default Status field that cannot be replaced
- **Workflows not accessible** - Project automation workflows cannot be read or written via API

These limitations led to removing Epic 2 (Project Templates) from scope.

[Unreleased]: https://github.com/scooter-indie/gh-pmu/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/scooter-indie/gh-pmu/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/scooter-indie/gh-pmu/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/scooter-indie/gh-pmu/compare/v0.0.1...v0.1.0
[0.0.1]: https://github.com/scooter-indie/gh-pmu/releases/tag/v0.0.1
