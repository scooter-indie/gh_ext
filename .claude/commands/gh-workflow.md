# GitHub Workflow Integration

This command configures Claude to automatically manage GitHub issues during development sessions.

---

## Project Configuration

**Read from `.gh-pmu.yml`** in the repository root. This file defines:
- Project board connection (owner, number)
- Repositories to track
- **Project field values** (Status, Priority) - these are NOT labels

### Labels vs Project Fields

**Labels** = Issue metadata (e.g., `bug`, `enhancement`, `pm-tracked`)
- Applied via: `gh issue edit --add-label "bug"`

**Project Fields** = Project board columns/values (e.g., Status, Priority)
- Updated via: `gh pmu move [number] --status [value]`
- Values defined in `.gh-pmu.yml` under `fields:`

```yaml
fields:
    status:
        field: Status
        values:
            backlog: Backlog
            in_progress: In progress
            in_review: In review
            done: Done
    priority:
        field: Priority
        values:
            p0: P0
            p1: P1
            p2: P2
```

Use the **alias** (left side) in commands: `gh pmu move 90 --status in_progress`

---

## Critical Rules

**NEVER close issues automatically.** Always wait for explicit "Done" from user.

---

## Workflow Steps

### Step 1: Create Issue (AUTOMATIC)
When user reports bug or requests enhancement, immediately create the issue.
Report: "Created issue #[number]. Let me know when you want me to work on it."

### Step 2: Work Issue (ONLY WHEN USER SAYS)
Wait for: "work issue #X", "fix that", "implement it"
Then: `gh pmu move [number] --status in_progress`

### Step 3: Commit and Review (AFTER WORK COMPLETE)
1. Commit with issue reference
2. `gh pmu move [number] --status in_review`
3. `gh issue comment [number] --body "Implemented in commit [hash]..."`

**STOP**: Do NOT close the issue.
Report: "Issue #[number] ready for review. Say 'Done' to close it."
Then WAIT for user response.

### Step 4: Close Issue (ONLY WHEN USER SAYS "DONE")
Wait for: "done", "close it", "approved", "looks good"
Then:
1. `gh pmu move [number] --status done`
2. `gh issue close [number]`

---

## Trigger Phrases

**Bug:** "I found an issue...", "There's a bug...", "finding:", "This is broken..."
**Enhancement:** "I would like...", "Can you add...", "New feature...", "Enhancement..."
**Sub-Issues:** "Create sub-issues for...", "Break this into phases..."
