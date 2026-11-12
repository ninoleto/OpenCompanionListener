# Contributing to Open Companion Listener

Contributions, bug reports, documentation improvements and independent forks are welcome. Code is released under the MIT license.

## Before you change anything

Read `AGENTS.md`, `HANDOFF.md`, `SECURITY.md`, and `docs/ARCHITECTURE.md`. Windows 10/11 x64 is the only supported desktop target in v0.1.4. Linux support is a future goal, not a working feature.

## Development workflow

1. Fork the repository and create a descriptive feature or fix branch.
2. Open an issue to discuss broad protocol, auth, input, or security changes before implementing them.
3. Make small, reviewable commits with regression tests. Preserve Companion action IDs and old saved buttons; add migrations if an action schema changes.
4. Use the toolchain and PowerShell/Bash commands in [the developer guide](docs/DEVELOPMENT.md). Run `go test -race ./...` and `go vet ./...` in `listener/`; run `npm ci --ignore-scripts --no-audit --no-fund`, `npm run validate`, then `npm test` in the module folder. Build and validate the TGZ for packaging changes. The real-library migration tests require installed dependencies.
5. Cross-build the Windows executable and, where feasible, run the `SMOKE_TEST.md` checklist on Windows and an actual Companion host.
6. In your PR describe what changed, why, tests run, screenshots for UI work, compatibility implications, and known gaps.

## What makes a useful issue report?

Include: listener version, Windows edition/build, Companion version and module version, relevant action settings without secrets, exact reproduction steps, expected vs observed result, and a trimmed/redacted log. Do not paste pairing keys or private hostnames.

## Code review priorities

- Respect Windows input focus/UIPI boundaries; no elevation bypasses or new shell endpoints. Existing executable/argument launching gives the authenticated peer desktop-user execution authority; it is not a restricted launcher.
- Keep dependencies light and startup latency low; benchmark before claiming performance improvements.
- Add explicit tests for authentication, replay rejection, input boundaries, old config migration and user-selected presets.
- Prefer intuitive dropdown/checkbox/slider actions to typed syntax while retaining supported old action IDs.
- Do not rewrite a working listener engine or Companion module for aesthetic changes.

See `docs/RELEASING.md` for versioning and artifact publication rules. The repo owner approves official releases; forks may independently release their own versions consistent with the MIT license.
