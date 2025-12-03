# Sprint 5 Retrospective: gh-pm-unified

**Date:** 2025-12-03

---

## What Went Well ✅

- Exceeded sprint goal by adding Story 3.4 mid-sprint (21 points vs 18 planned)
- Story 3.1 was already implemented - quick verification and close
- Cross-repo testing revealed important design insight about GitHub Projects
- TDD approach continued to be effective for flag validation tests
- Recursive move implementation was clean and well-tested on first attempt
- Progress bar visualization works correctly with edge cases (0%, overflow)

---

## What Could Be Improved 🔄

- Story point estimate for 3.1 was too high (estimated 5, actual ~1 since already done)
- Initial acceptance criteria for 3.2 included invalid requirement ("validate repos in same project")
- Had to research GitHub Projects behavior during sprint - could have been done in planning
- Config file naming inconsistency (.gh-pm.yml vs .gh-pmu.yml) caused testing friction
- Status field values vary by project ("in_progress" vs "In progress") - no normalization

---

## Key Learnings 📚

### GitHub Projects v2 Architecture
- Projects can contain issues from **any** repository
- No built-in validation for repo membership - projects are cross-org capable
- This changes assumptions about multi-repo workflows

### Testing Approach
- Created dedicated test repos (gh-pmu-repo-1, gh-pmu-repo-2) and projects (27, 28)
- Separating test infrastructure from production project (11) was valuable
- Manual verification complemented unit tests well

---

## Action Items for Next Sprint 🎯

- [ ] Standardize config file naming (.gh-pmu.yml everywhere)
- [ ] Consider adding status value normalization in config resolution
- [ ] Document cross-repo sub-issue capabilities in README
- [ ] Review remaining Epic 3 stories for completion
- [ ] Consider adding `--tree` view to list command

---

## Process Adjustments

**Planning:** Research platform capabilities during backlog refinement, not during sprint execution.

**Testing:** Continue using dedicated test repos/projects for safe experimentation.

**Story Sizing:** Verify implementation state before estimating - some stories may already be done.

---

## Velocity Trends

| Sprint | Points | Stories | Notes |
|--------|--------|---------|-------|
| 1 | 28 | 5 | Foundation |
| 2 | 24 | 4 | Core commands |
| 3 | 21 | 3 | Complex stories |
| 4 | - | - | Skipped |
| 5 | 21 | 4 | Added story mid-sprint |

**Cumulative Velocity:** 94 points across 4 sprints
**Average Velocity:** 23.5 points per sprint

---

## Epic Status

### Epic 1: Core Unification ✅
All 12 commands implemented and tested.

### Epic 3: Enhanced Sub-Issue Integration ✅
All 4 stories (3.1, 3.2, 3.3, 3.4) complete:
- Native sub-issue handling in split
- Cross-repository sub-issues
- Progress tracking
- Recursive operations

---

## Team Notes

Sprint 5 demonstrated good adaptability - Story 3.4 was pulled in from backlog when capacity allowed. The recursive move feature adds significant value for managing large issue hierarchies.

The investigation into GitHub Projects' cross-repo capabilities was valuable. Removing the invalid acceptance criterion rather than implementing unnecessary validation was the right call.

---

## Next Steps

1. Update product backlog with remaining work
2. Plan Sprint 6 if more stories exist
3. Consider release preparation if MVP is complete
