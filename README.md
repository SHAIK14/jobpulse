# jobpulse

A CLI tool to track cold outreach during your job hunt. Built because bulk applying on LinkedIn goes nowhere — the real work is targeted DMs to the right person at the right company.

## The problem it solves

When you're sending cold DMs to CTOs, HRs, and team leads, you lose track of:
- Who you already contacted
- Whether they replied
- Who you need to follow up with

jobpulse fixes that.

## Install

```bash
go install github.com/SHAIK14/jobpulse@latest
```

## Commands

### Track outreach

```bash
# Add a new outreach
jobpulse add -company=Zepto -joblink="https://..." -contact="https://linkedin.com/in/someone" -title=CTO

# See all outreaches
jobpulse list

# See full details for one entry (with untruncated links)
jobpulse show -id=1

# Update status after a reply
jobpulse update -id=1 -status=Replied

# Delete an entry
jobpulse del -id=1

# See who you DMed 3+ days ago with no reply
jobpulse followup
```

### Status values
`DMed` → `Replied` → `Interview` → `Offered` → `Rejected` / `Ghosted`

### Manage DM templates

```bash
# Save a template (upserts — no duplicates)
jobpulse template add -role=cto -body="Hey, I saw you're building something interesting..."
jobpulse template add -role=hr -body="Hi, I came across the opening for Backend Engineer..."

# See all templates
jobpulse template list

# Update a template
jobpulse template update -role=cto -body="Updated message..."
```

## Data

Everything is stored locally in your home directory:
- `~/.jobpulse.json` — your outreach history
- `~/.jobpulse_templates.json` — your DM templates

No database, no server, no cloud. Just files on your machine.

## Built with

Go standard library only — `encoding/json`, `flag`, `text/tabwriter`, `os`, `time`. No external dependencies.
