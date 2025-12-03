# Claude Code - Project Instructions

**Purpose:** Automatic initialization with IDPF Framework integration
**Process Framework:** IDPF-Agile
**Domain Specialists:** Backend-Specialist, API-Integration-Specialist, Security-Engineer, PRD-Analyst
**Primary Specialist:** API-Integration-Specialist

---

## Framework Configuration

This project uses the IDPF Framework ecosystem.
**Configuration:** See `framework-config.json` for framework location and project type.

---

## Startup Procedure

When starting a new session in this repository, **IMMEDIATELY** perform these steps:

### Step 1: Confirm Date

State the date from your environment information and ask the user to confirm it is correct. **Wait for the user to respond before proceeding to Step 2.**

```
"According to my environment information, today's date is YYYY-MM-DD. Is this correct?"
```

If the user responds "no", prompt for the correct date in YYYY-MM-DD format.

This ensures accurate timestamps in commits and documentation.

### Step 2: Load Configuration

Read `framework-config.json` to get the `frameworkPath`.

### Step 3: Load Startup Instructions and Framework Core

Read these files in order:
1. `STARTUP.md` - Condensed essential rules and guidelines
2. `E:\Projects\process-docs/IDPF-Agile/Agile-Core.md` - Core framework workflow


### Step 4: Load Primary Domain Specialist

Read the primary specialist instructions to activate this role:

`E:\Projects\process-docs/System-Instructions/Domain/API-Integration-Specialist.md`

**Active Role:** API-Integration-Specialist

### Step 5: Configure GitHub Integration (if needed)

If `.gh-pm.yml` does not exist, ask user if they have a GitHub repo and project.
If yes, run `gh pm init`. If no, skip.

If `.claude/commands/gh-workflow.md` has unreplaced placeholders, prompt user for values.

### Step 6: Confirm Ready

Confirm initialization is complete and ask the user what they would like to work on.
Include the active role in your ready message: "Active Role: API-Integration-Specialist"

**Do NOT proceed with any other work until the startup sequence is complete.**

---

## Available Commands

After completing the startup procedure, display available commands:

| Command | Purpose |
|---------|---------|
| `/expand-rules` | Load complete Anti-Hallucination Rules |
| `/expand-domain` | Load full Domain Specialist instructions |
| `/switch-role` | Switch active domain specialist mid-session |
| `/gh-workflow` | Activate GitHub workflow integration |

### When to Suggest Commands

**Proactively suggest `/expand-rules` when:**
- User asks about code quality or best practices
- Reviewing code for potential issues
- User mentions concerns about accuracy or hallucination

**Proactively suggest `/expand-domain` when:**
- Working on domain-specific tasks (backend, frontend, DevOps, etc.)
- User needs specialized technical guidance
- Implementing complex features in a specific domain

---

## Project-Specific Instructions

<!-- Add your project-specific instructions below this line -->

paint a pretty ascii picture on session startup

<!-- These will be preserved during framework updates -->

---

**End of Claude Code Instructions**
