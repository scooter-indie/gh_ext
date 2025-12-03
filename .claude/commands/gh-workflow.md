# GitHub Workflow Integration

This command configures Claude to automatically manage GitHub issues during development sessions.

---

## Project Configuration

**Read from `.gh-pm.yml`** in the repository root:

```yaml
project:
    owner: {owner}      # GitHub username or org
    number: {number}    # Project board number
repositories:
    - {owner}/{repo}    # Repository in owner/repo format
```

If `.gh-pm.yml` doesn't exist, run `gh pm init` to create it.

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
Then: `gh pm move [number] --status in_progress`

### Step 3: Commit and Review (AFTER WORK COMPLETE)
1. Commit with issue reference
2. `gh pm move [number] --status in_review`
3. `gh issue comment [number] --body "Implemented in commit [hash]..."`

⚠️ **STOP**: Do NOT close the issue.
Report: "Issue #[number] ready for review. Say 'Done' to close it."
Then WAIT for user response.

### Step 4: Close Issue (ONLY WHEN USER SAYS "DONE")
Wait for: "done", "close it", "approved", "looks good"
Then:
1. `gh pm move [number] --status done`
2. `gh issue close [number]`

---

## Trigger Phrases

**Bug:** "I found an issue...", "There's a bug...", "finding:", "This is broken..."
**Enhancement:** "I would like...", "Can you add...", "New feature...", "Enhancement..."
**Sub-Issues:** "Create sub-issues for...", "Break this into phases..."

---

**Note:** Replace {owner}, {repo}, {number} placeholders after running `gh pm init`.
