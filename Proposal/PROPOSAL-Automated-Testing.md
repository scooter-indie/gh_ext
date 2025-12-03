# Proposal: Automated Non-Destructive Integration Tests & UAT

**Version:** 1.0
**Date:** 2025-12-03
**Author:** PRD-Analyst
**Status:** Draft

---

## Executive Summary

### Problem Statement

The gh-pmu CLI extension currently has ~36% test coverage, primarily unit tests covering command structure and configuration parsing. The existing test suite lacks:

1. **Integration tests** that verify actual GitHub API interactions
2. **End-to-end workflow tests** covering complete user scenarios
3. **UAT scenarios** validating acceptance criteria from a user perspective
4. **Non-destructive test patterns** that can run safely against real GitHub projects

### Proposed Solution

Implement a comprehensive automated testing strategy using **non-destructive patterns** that:
- Test against dedicated test fixtures (projects, repos, issues)
- Use read-only operations where possible
- Clean up any created resources after tests
- Support both CI/CD automation and local development

### Key Benefits

| Benefit | Impact |
|---------|--------|
| Increased confidence in releases | Catch API integration bugs before users |
| Faster development cycles | Automated regression testing |
| Living documentation | Tests serve as usage examples |
| Safer refactoring | Full coverage enables bold improvements |

---

## Scope

### In Scope

- Integration tests for all GraphQL API operations
- End-to-end tests for all CLI commands
- UAT scenarios covering Epics 1-3 acceptance criteria
- Test fixture management (setup/teardown)
- CI/CD pipeline integration
- Test documentation and runbooks

### Out of Scope

- Load/performance testing
- Security penetration testing
- Template ecosystem testing (Epic 4 - future)
- Cross-platform compatibility testing (Windows/macOS/Linux variations)

---

## Non-Destructive Testing Strategy

### Core Principles

1. **Isolation**: Tests use dedicated test projects/repos, never production data
2. **Idempotency**: Tests can run multiple times without side effects
3. **Cleanup**: Any created resources are removed after test completion
4. **Read-First**: Prefer read operations; only write when necessary to validate
5. **Atomic Transactions**: Each test is independent and self-contained

### Test Environment Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    GitHub Test Organization                  │
│                    (gh-pmu-test-org)                        │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  ┌──────────────────┐  ┌──────────────────┐                │
│  │  Test Repo 1     │  │  Test Repo 2     │                │
│  │  (primary)       │  │  (cross-repo)    │                │
│  │                  │  │                  │                │
│  │  - Seed issues   │  │  - Seed issues   │                │
│  │  - Labels        │  │  - Labels        │                │
│  └──────────────────┘  └──────────────────┘                │
│                                                              │
│  ┌──────────────────────────────────────────────────────┐  │
│  │              Test Project (GitHub Projects v2)        │  │
│  │                                                        │  │
│  │  Fields: Status, Priority, Sprint, Estimate           │  │
│  │  Views: Kanban, Table                                 │  │
│  │  Items: Pre-seeded test issues                        │  │
│  └──────────────────────────────────────────────────────┘  │
│                                                              │
└─────────────────────────────────────────────────────────────┘
```

### Resource Management Strategy

| Resource Type | Strategy | Cleanup Method |
|--------------|----------|----------------|
| Test Project | Persistent fixture | Reset to known state |
| Test Issues | Created per-suite | Delete after suite |
| Project Items | Created per-test | Remove from project |
| Sub-Issues | Created per-test | Unlink and delete |
| Field Values | Modified per-test | Reset to defaults |

---

## Test Categories

### Category 1: API Integration Tests

**Purpose:** Verify GraphQL queries and mutations work correctly with real GitHub API

**Approach:**
- Use `//go:build integration` tag
- Run against test fixtures
- Validate response structure and data integrity

**Coverage Areas:**

| Module | Operations | Test Count (Est.) |
|--------|-----------|-------------------|
| Queries | GetProject, GetFields, GetIssue, GetItems, etc. | 15 |
| Mutations | CreateIssue, AddToProject, SetField, LinkSubIssue, etc. | 20 |
| Error Handling | Auth failures, Not found, Rate limits | 10 |
| Edge Cases | Empty results, Large payloads, Special characters | 8 |

**Example Test Pattern:**

```go
//go:build integration

func TestGetProjectFields_ReturnsAllFieldTypes(t *testing.T) {
    // Arrange: Use test project with known fields
    client := setupTestClient(t)
    projectID := os.Getenv("TEST_PROJECT_ID")

    // Act: Query fields
    fields, err := client.GetProjectFields(projectID)

    // Assert: Verify expected fields exist
    require.NoError(t, err)
    assert.Contains(t, fieldNames(fields), "Status")
    assert.Contains(t, fieldNames(fields), "Priority")

    // Cleanup: None needed (read-only)
}
```

### Category 2: Command Integration Tests

**Purpose:** Verify CLI commands work end-to-end with real API

**Approach:**
- Execute actual `gh pmu` commands via subprocess
- Capture stdout/stderr
- Validate output format and content
- Verify side effects via API queries

**Coverage Areas:**

| Command | Scenarios | Test Count (Est.) |
|---------|-----------|-------------------|
| `init` | New config, existing config, auto-detect | 4 |
| `list` | Filter by status, JSON output, empty results | 6 |
| `view` | Valid issue, invalid issue, with fields | 4 |
| `create` | With fields, minimal, validation errors | 5 |
| `move` | Change status, change priority, multiple fields | 5 |
| `intake` | Find untracked, add to project, dry-run | 4 |
| `triage` | Apply rules, skip processed, dry-run | 5 |
| `sub add` | Link existing, cross-repo, invalid parent | 4 |
| `sub create` | New sub-issue, with fields, inherit parent | 4 |
| `sub list` | With children, no children, recursive | 3 |
| `sub remove` | Unlink, delete, not found | 3 |
| `split` | From checklist, from args, validation | 4 |

**Example Test Pattern:**

```go
//go:build integration

func TestListCommand_FiltersByStatus(t *testing.T) {
    // Arrange: Ensure test issues exist with known statuses
    setupTestIssues(t, []TestIssue{
        {Title: "Test-Todo", Status: "Todo"},
        {Title: "Test-Done", Status: "Done"},
    })
    defer cleanupTestIssues(t)

    // Act: Run list command with filter
    output, err := runCommand("gh", "pmu", "list", "--status", "Todo")

    // Assert: Only Todo issues in output
    require.NoError(t, err)
    assert.Contains(t, output, "Test-Todo")
    assert.NotContains(t, output, "Test-Done")
}
```

### Category 3: User Acceptance Tests (UAT)

**Purpose:** Validate complete user workflows and acceptance criteria from PRD

**Approach:**
- Scenario-based tests using Given-When-Then structure
- Cover complete user journeys (multi-command workflows)
- Validate business value delivery
- Human-readable test names and documentation

**UAT Scenarios by Epic:**

#### Epic 1: Core Unification

| ID | Scenario | Acceptance Criteria |
|----|----------|---------------------|
| UAT-1.1 | Initialize new project | User can run `gh pmu init` and get working config |
| UAT-1.2 | List and filter issues | User can list issues filtered by status/priority |
| UAT-1.3 | Create issue with fields | User can create issue with status/priority pre-set |
| UAT-1.4 | Move issue through workflow | User can update issue status/priority via CLI |
| UAT-1.5 | Intake untracked issues | User can find and add issues not yet in project |
| UAT-1.6 | Triage with rules | User can apply configurable rules to categorize issues |
| UAT-1.7 | Manage sub-issues | User can create, link, list, and remove sub-issues |
| UAT-1.8 | Split issue into tasks | User can split issue into sub-issues from checklist |

#### Epic 2: Project Templates

| ID | Scenario | Acceptance Criteria |
|----|----------|---------------------|
| UAT-2.1 | Create from template | User can create project from YAML template |
| UAT-2.2 | Export project | User can export existing project to YAML |
| UAT-2.3 | Validate template | User can validate template before use |
| UAT-2.4 | List templates | User can discover available templates |

#### Epic 3: Enhanced Integration

| ID | Scenario | Acceptance Criteria |
|----|----------|---------------------|
| UAT-3.1 | Cross-repo sub-issues | User can link issues across repositories |
| UAT-3.2 | Progress tracking | User can see sub-issue completion percentage |
| UAT-3.3 | Recursive operations | User can bulk update issue hierarchies |

**Example UAT Test:**

```go
//go:build uat

func TestUAT_1_3_CreateIssueWithFields(t *testing.T) {
    /*
    Scenario: Create issue with project fields pre-populated

    Given I have initialized gh-pmu with a valid configuration
    And the project has Status and Priority fields
    When I run: gh pmu create --title "New Feature" --status "Todo" --priority "High"
    Then a new issue should be created
    And it should be added to the project
    And the Status field should be "Todo"
    And the Priority field should be "High"
    */

    // Given
    cfg := setupTestConfig(t)

    // When
    output, err := runCommand("gh", "pmu", "create",
        "--title", "UAT-Test-Issue-"+randomSuffix(),
        "--status", "Todo",
        "--priority", "High",
    )
    require.NoError(t, err)

    // Then
    issueNum := extractIssueNumber(output)
    defer deleteTestIssue(t, issueNum)

    issue := getIssueWithFields(t, issueNum)
    assert.Equal(t, "Todo", issue.Fields["Status"])
    assert.Equal(t, "High", issue.Fields["Priority"])
}
```

---

## Test Infrastructure

### Test Fixtures

#### Required GitHub Resources

1. **Test Organization:** `gh-pmu-test-org` (or user account)
2. **Test Repository 1:** `gh-pmu-test-repo` (primary)
3. **Test Repository 2:** `gh-pmu-test-repo-2` (cross-repo tests)
4. **Test Project:** GitHub Projects v2 with standard fields

#### Fixture Configuration

```yaml
# testdata/fixtures/test-project.yml
project:
  title: "gh-pmu Integration Test Project"
  owner: "gh-pmu-test-org"
  fields:
    - name: Status
      type: single_select
      options: [Todo, In Progress, In Review, Done]
    - name: Priority
      type: single_select
      options: [Low, Medium, High, Critical]
    - name: Sprint
      type: iteration
    - name: Estimate
      type: number
  views:
    - name: Kanban
      type: board
      group_by: Status
    - name: All Items
      type: table

seed_issues:
  - title: "[TEST] Seed Issue 1"
    status: Todo
    priority: Medium
  - title: "[TEST] Seed Issue 2"
    status: In Progress
    priority: High
  - title: "[TEST] Seed Issue 3 (Parent)"
    status: Todo
    priority: Low
    sub_issues:
      - title: "[TEST] Sub-Issue 3.1"
      - title: "[TEST] Sub-Issue 3.2"
```

### Test Utilities Package

Create `internal/testutil/` package with:

```go
// internal/testutil/testutil.go
package testutil

// Client setup
func SetupTestClient(t *testing.T) *api.Client
func GetTestProjectID() string
func GetTestRepoOwner() string
func GetTestRepoName() string

// Issue management
func CreateTestIssue(t *testing.T, title string, opts ...IssueOption) int
func DeleteTestIssue(t *testing.T, issueNum int)
func CleanupTestIssues(t *testing.T)

// Project item management
func AddIssueToProject(t *testing.T, issueNum int) string
func RemoveItemFromProject(t *testing.T, itemID string)
func ResetProjectItem(t *testing.T, itemID string)

// Assertions
func AssertIssueHasField(t *testing.T, issueNum int, field, value string)
func AssertIssueInProject(t *testing.T, issueNum int)

// Command execution
func RunCommand(t *testing.T, args ...string) (string, error)
func RunCommandWithConfig(t *testing.T, cfg *config.Config, args ...string) (string, error)
```

### CI/CD Integration

#### GitHub Actions Workflow

```yaml
# .github/workflows/integration-tests.yml
name: Integration Tests

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
  schedule:
    - cron: '0 6 * * *'  # Daily at 6 AM UTC

jobs:
  integration:
    runs-on: ubuntu-latest
    environment: integration-tests

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Run Integration Tests
        env:
          GH_TOKEN: ${{ secrets.TEST_GH_TOKEN }}
          TEST_PROJECT_ID: ${{ vars.TEST_PROJECT_ID }}
          TEST_REPO_OWNER: ${{ vars.TEST_REPO_OWNER }}
          TEST_REPO_NAME: ${{ vars.TEST_REPO_NAME }}
        run: |
          go test -v -tags=integration ./...

      - name: Run UAT Tests
        env:
          GH_TOKEN: ${{ secrets.TEST_GH_TOKEN }}
          TEST_PROJECT_ID: ${{ vars.TEST_PROJECT_ID }}
          TEST_REPO_OWNER: ${{ vars.TEST_REPO_OWNER }}
          TEST_REPO_NAME: ${{ vars.TEST_REPO_NAME }}
        run: |
          go test -v -tags=uat ./...

      - name: Cleanup Test Resources
        if: always()
        env:
          GH_TOKEN: ${{ secrets.TEST_GH_TOKEN }}
        run: |
          go run ./cmd/testcleanup/main.go
```

### Required Secrets/Variables

| Name | Type | Purpose |
|------|------|---------|
| `TEST_GH_TOKEN` | Secret | GitHub PAT with repo, project, admin:org scopes |
| `TEST_PROJECT_ID` | Variable | Project node ID for test project |
| `TEST_REPO_OWNER` | Variable | Owner of test repositories |
| `TEST_REPO_NAME` | Variable | Primary test repository name |
| `TEST_REPO_NAME_2` | Variable | Secondary test repository (cross-repo) |

---

## Implementation Roadmap

### Phase 1: Foundation (Story Points: 13)

**Deliverables:**
- [ ] Create test organization/repos/project on GitHub
- [ ] Implement `internal/testutil/` package
- [ ] Create fixture seed scripts
- [ ] Set up CI/CD workflow skeleton
- [ ] Document test environment setup

**Dependencies:** None

### Phase 2: API Integration Tests (Story Points: 21)

**Deliverables:**
- [ ] Query tests (GetProject, GetFields, GetIssue, etc.)
- [ ] Mutation tests (CreateIssue, SetField, LinkSubIssue, etc.)
- [ ] Error handling tests
- [ ] Edge case coverage

**Dependencies:** Phase 1

### Phase 3: Command Integration Tests (Story Points: 21)

**Deliverables:**
- [ ] Init command tests
- [ ] List/View command tests
- [ ] Create/Move command tests
- [ ] Sub-issue command tests
- [ ] Intake/Triage command tests
- [ ] Split command tests

**Dependencies:** Phase 2

### Phase 4: UAT Implementation (Story Points: 13)

**Deliverables:**
- [ ] Epic 1 UAT scenarios (8 tests)
- [ ] Epic 2 UAT scenarios (4 tests)
- [ ] Epic 3 UAT scenarios (3 tests)
- [ ] UAT documentation and runbook

**Dependencies:** Phase 3

### Phase 5: Polish & Documentation (Story Points: 5)

**Deliverables:**
- [ ] Test coverage report integration
- [ ] Test documentation (TESTING.md)
- [ ] Local development testing guide
- [ ] CI/CD optimization (parallelization, caching)

**Dependencies:** Phase 4

---

## Success Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| Integration test coverage | 80% of API operations | Query/mutation coverage |
| Command test coverage | 100% of commands | At least 1 test per command |
| UAT coverage | 100% of PRD acceptance criteria | Traced to PRD |
| Test execution time | < 5 minutes | CI/CD duration |
| Test reliability | < 1% flaky rate | Failure tracking |
| Cleanup success rate | 100% | No orphaned resources |

---

## Risks & Mitigations

| Risk | Impact | Likelihood | Mitigation |
|------|--------|------------|------------|
| GitHub API rate limiting | Tests fail or slow down | Medium | Use caching, batch operations, retry logic |
| Test fixture corruption | Tests become unreliable | Low | Reset fixtures before each run |
| Flaky tests from network | False failures | Medium | Retry logic, longer timeouts |
| Token exposure | Security breach | Low | Use GitHub secrets, rotate regularly |
| Test org costs | Unexpected charges | Low | Monitor usage, use free tier limits |

---

## Open Questions

| # | Question | Impact | Owner |
|---|----------|--------|-------|
| 1 | Use existing org or create dedicated test org? | Setup complexity | TBD |
| 2 | Should UAT tests run on every PR or just main? | CI time vs coverage | TBD |
| 3 | How to handle GitHub API preview feature changes? | Test stability | TBD |
| 4 | Should we mock any API calls for speed? | Test fidelity vs speed | TBD |

---

## Appendix A: Test File Structure

```
gh-pm-unified/
├── internal/
│   ├── api/
│   │   ├── queries_integration_test.go     # //go:build integration
│   │   └── mutations_integration_test.go   # //go:build integration
│   ├── testutil/
│   │   ├── testutil.go                     # Test utilities
│   │   ├── fixtures.go                     # Fixture management
│   │   └── cleanup.go                      # Resource cleanup
│   └── config/
│       └── config_integration_test.go      # //go:build integration
├── cmd/
│   ├── list_integration_test.go            # //go:build integration
│   ├── create_integration_test.go          # //go:build integration
│   ├── move_integration_test.go            # //go:build integration
│   └── ...
├── test/
│   ├── uat/
│   │   ├── epic1_test.go                   # //go:build uat
│   │   ├── epic2_test.go                   # //go:build uat
│   │   └── epic3_test.go                   # //go:build uat
│   └── fixtures/
│       ├── test-project.yml                # Project fixture definition
│       └── seed-issues.yml                 # Issue seed data
├── testdata/
│   └── configs/
│       ├── valid-config.yml                # Test configs
│       └── invalid-config.yml
└── TESTING.md                              # Test documentation
```

---

## Appendix B: Makefile Targets

```makefile
# Test targets
.PHONY: test test-unit test-integration test-uat test-all

test: test-unit                              ## Run unit tests only (default)

test-unit:                                   ## Run unit tests
	go test -v ./...

test-integration:                            ## Run integration tests (requires TEST_* env vars)
	go test -v -tags=integration ./...

test-uat:                                    ## Run UAT tests (requires TEST_* env vars)
	go test -v -tags=uat ./...

test-all: test-unit test-integration test-uat ## Run all tests

test-coverage:                               ## Run tests with coverage
	go test -coverprofile=coverage.out ./...
	go test -tags=integration -coverprofile=coverage-integration.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-setup:                                  ## Setup test fixtures
	go run ./cmd/testsetup/main.go

test-cleanup:                                ## Cleanup test resources
	go run ./cmd/testcleanup/main.go
```

---

## Revision History

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 1.0 | 2025-12-03 | PRD-Analyst | Initial proposal |

