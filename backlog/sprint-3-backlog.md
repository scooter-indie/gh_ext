# Sprint 3 Backlog: gh-pm-unified

**Sprint Goal:** Complete Core Unification (Epic 1) - Finish all remaining core project management features
**Sprint Duration:** 2025-12-02 to TBD
**Total Story Points:** 21

---

## Selected Stories

### Story 1.6: Issue Intake - Find Untracked Issues

**As a** project manager
**I want** to find issues not yet added to the project
**So that** I can ensure all work is tracked on the project board

**Acceptance Criteria:**
- [x] `gh pmu intake` finds open issues not in the project
- [x] Shows list of untracked issues with titles
- [x] `--apply` flag adds them to project with default fields
- [x] `--dry-run` shows what would be added
- [x] Respects repository filter from config

**Story Points:** 5
**Status:** Done

---

### Story 1.7: Triage - Bulk Process Issues

**As a** project manager
**I want** to bulk update issues matching certain criteria
**So that** I can efficiently maintain project hygiene

**Acceptance Criteria:**
- [x] `gh pmu triage <config-name>` runs named triage config
- [x] Triage configs defined in `.gh-pmu.yml` with query and apply rules
- [x] Supports applying labels, status, priority changes
- [x] `--interactive` flag prompts for each issue
- [x] `--dry-run` shows what would be changed
- [x] Reports summary of changes made

**Story Points:** 8
**Status:** Done

---

### Story 1.12: Split Issue into Sub-Issues

**As a** developer breaking down an epic
**I want** to split an issue's checklist into sub-issues
**So that** I can convert task lists into trackable issues

**Acceptance Criteria:**
- [x] `gh pmu split <issue> --from=body` parses checklist from issue body
- [x] `gh pmu split <issue> --from=file.md` parses from external file
- [x] `gh pmu split <issue> "Task 1" "Task 2"` creates from arguments
- [x] Each checklist item becomes a sub-issue
- [x] Sub-issues linked to parent automatically
- [x] Shows summary of created sub-issues

**Story Points:** 8
**Status:** Done

---

## Sprint Progress

**Completed:** 21 story points
**Remaining:** 0 story points
**Velocity:** 21 points

---

## Notes & Blockers

*Sprint started 2025-12-02*

---
