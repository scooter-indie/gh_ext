# Switch Domain Specialist Role

Switch to a different domain specialist role for this session.

## Available Roles

1. Backend-Specialist
2. API-Integration-Specialist
3. Security-Engineer
4. PRD-Analyst

## Current Primary

**API-Integration-Specialist** is loaded at session startup.

## Instructions

When invoked, display the available roles and ask the user to select one.
Upon selection:
1. **Deactivate previous role**: State that previous role instructions are now inactive
2. Read the new specialist's instruction file from the path below
3. **Confirm exclusive operation**: State that you are now operating exclusively as the new role
4. Apply that specialist's expertise to subsequent work

**Important:** Previous role instructions remain in conversation context but are explicitly deprioritized. The new role takes exclusive precedence.

## File Paths

- Backend-Specialist: `E:\Projects\process-docs/System-Instructions/Domain/Backend-Specialist.md`
- API-Integration-Specialist: `E:\Projects\process-docs/System-Instructions/Domain/API-Integration-Specialist.md`
- Security-Engineer: `E:\Projects\process-docs/System-Instructions/Domain/Security-Engineer.md`
- PRD-Analyst: `E:\Projects\process-docs/System-Instructions/Domain/PRD-Analyst.md`

## Usage

User says: `/switch-role` or "switch to frontend" or "I need backend help now"

**Response format:**
```
⊘ Deactivating: [Previous-Role]

Loading [New-Role]...

✓ Now operating exclusively as: [New-Role]
  Focus areas: [role-specific focus from the specialist file]

  Previous role instructions ([Previous-Role]) are now inactive.

What would you like to work on?
```

## Natural Language Triggers

- "switch to [role]"
- "I need [role] help"
- "change to [role] mode"
- "activate [role]"
