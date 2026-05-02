# Publish this fork on GitHub

Work lives on branch `fork/send-text-message` (commit: real `tg_send` via `messages.sendMessage`, plus `tg_save_draft` for drafts).

## 1. Log in (once)

```bash
gh auth login
```

Or set a classic PAT with `repo` scope and:

```bash
export GH_TOKEN=ghp_...
```

## 2. Create the fork (GitHub CLI)

```bash
gh repo fork chaindead/telegram-mcp --remote=false
```

That creates `YOUR_USER/telegram-mcp` under your account.

## 3. Push this branch

From this directory:

```bash
git remote add myfork https://github.com/YOUR_USER/telegram-mcp.git
git push -u myfork fork/send-text-message:main
```

(Use `main` or open a PR from `fork/send-text-message` into your fork’s `main` instead.)

## 4. Optional: git bundle (air‑gap / copy)

```bash
git bundle create ../telegram-mcp-send.bundle fork/send-text-message
# On another machine: git clone telegram-mcp-send.bundle telegram-mcp-send
```
