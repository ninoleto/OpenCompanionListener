# Development handoff — Open Companion Listener

**Last verified source baseline:** `v0.1.3`, released package built 2026-11-10.  
**Scope:** Windows 10/11 x64 listener + Bitfocus Companion `community-openlistener` module.  
**Project purpose:** A free, independent, lightweight desktop-control system. Public contributions and forks are explicitly welcomed (MIT).

## What already works and must stay working

- Windows-native Go listener with silent system-tray startup, Open GUI / Restart Listener / Exit, local settings dashboard, opt-in **Start with Windows** (HKCU Run), and single-instance behavior anchored by local dashboard port.
- Go HTTP API: configurable trusted peer IP and port 17992; 64-character hex pairing key; signed commands, nonce replay protection and clock-skew checking. Local UI on 127.0.0.1:17993.
- Windows input: single keys including F1–F24 and numpad, shortcut press/hold/release, mouse movement/buttons/hold/release, vertical/horizontal wheel scrolling, Num Lock toggle/on/off.
- Direct Windows Shell launching: HTTP/HTTPS URLs, application executables, files/folders, Windows Settings URIs and Control Panel tools. No generic remote shell action.
- Companion API 1.x CommonJS module, checked into `companion-dev-modules/companion-module-community-openlistener/` with `@companion-module/base` dependency. 38 curated presets and GUI-oriented action builders.
- Action migrations (older scroll and keyboard IDs) via `src/upgrades.cjs`. **Never remove published migrations**; migration behavior must be verified on a backed-up real Companion configuration, not just unit-tested.

## Verification: distinguish tested from untested

- **Automated:** On 2026-11-10, Linux-host Go tests passed; Windows cross-target static checks passed; Companion Node tests passed (16 total). Go race checks were reported as passed in the build session that produced v0.1.3; rerun on any changed code.
- **User testing:** Prior Windows-native input and application-launch functions from the earlier 0.1.x releases were exercised on an actual PC. The user subsequently installed v0.1.3 and reported it looked good. That is useful feedback, **not evidence that every 0.1.3 feature passed the full smoke test**.
- **Still requires explicit release verification:** Windows sign-in autostart registry toggle, silent launch and duplicate instance prevention; real installed Companion upgrade behavior including legacy button migration; all 38 presets; corner cases for modifier hold/release and interruption.

## Architecture and files

- `listener/main.go` — UI HTTP server, authenticated API, lifecycle, key rotation, startup setting handlers.
- `listener/config.go` — local user settings and pairing key storage.
- `listener/auth.go` — signature checks, timestamp + nonce replay window.
- `listener/commands.go` — accepted command schema and bounds.
- `listener/desktop_windows.go`, `listener/keys.go` — native input and Shell launching; non-Windows stubs are NOT a Linux implementation.
- `listener/tray_windows.go`, `listener/startup_windows.go` — native Windows user session lifecycle.
- `listener/ui/index.html` — browser-rendered loopback dashboard.
- `companion-dev-modules/companion-module-community-openlistener/src/main.cjs` — Companion actions/config, status pings.
- `src/transport.cjs` — canonical request signing; `src/upgrades.cjs` — saved-action migrations; `src/presets.cjs` — 38 optional presets.

Detailed flow, command schema and threat model: `docs/ARCHITECTURE.md` and `SECURITY.md`.

## Known limitations and engineering risks

1. **HMAC is not confidentiality.** Requests use HTTP without TLS; source IP + HMAC do not make Internet exposure safe. User must firewall to the specific trusted Companion host and use LAN/private overlay.
2. Pairing key is stored in a per-user JSON file. Do not commit real config, print keys in logs, or misrepresent it as Windows credential-vault protected.
3. GUI automation follows Windows focus and integrity restrictions. Elevated/locked desktops cannot be promised as supported.
4. No Windows service, no background input when logged out. Autostart takes effect at **user sign-in** only.
5. `key_down`/`key_up` may need extra fail-safe handling for interrupted connections and unreleased modifiers; verify and design targeted tests rather than assuming automatic release.
6. Migration in `src/upgrades.cjs`: non-integer or expression-based legacy scroll values may stay unchanged; test against real Companion exported configurations and decide explicit recovery steps.
7. Existing `example.org/open-companion-listener` Go module path is a neutral placeholder; update only as part of a reviewed repository-bootstrap change and rerun builds. Module manifest `repository`/`bugs` are currently empty—populate them **after the real GitHub URL exists**.
8. **Linux:** No usable Linux input backend. Build tags allow non-Windows compilation only. Wayland/X11 require researched, consent-aware backend choices and explicit UI privilege constraints.
9. No built-in AHK builder yet. The roadmap considers optional external automation with clear permission boundaries, not unrestricted remote shell execution.

## Working agreements for future forks / AI agents

- Keep `v0.1.3` as a rollback checkpoint. Do not overwrite tags, rewrite history, or publish an untested EXE in place of a stable download.
- Change one subsystem per PR (listener core, GUI, Companion actions, migration, or docs). Keep action IDs + signed transport compatible, or create and test a migration.
- Use the exact test commands in `AGENTS.md`; gather **Windows host** and **Companion host** test logs before claiming interoperability.
- Back up installed Companion module and Companion settings before upgrades; don't run test module changes against active streaming sessions.
- For a bug report include version, OS build, Companion version, steps, expected/actual behavior, and **redacted** logs. Never share pairing key.

## Suggested next work, ordered

1. Have the maintainer complete `SMOKE_TEST.md` on v0.1.3, with special attention to autostart and action migration.
2. Create the first public GitHub repository (no public repository URL was present when this handoff was authored), populate `repository` and `bugs` manifest values with that actual URL in a follow-up patch, and publish v0.1.3 artifacts with SHA-256 checksums.
3. Triage any user-reported bugs and issue **new patch tags** (`v0.1.4`, etc.). Do not break existing Companion buttons.
4. Evaluate optional automation workflows/AHK builder and future Linux support as separately scoped follow-up milestones.

**Next safe action:** finish v0.1.3 on-device smoke tests; upload the reviewed source to an owner-created public GitHub repository and attach tested Windows/module ZIPs as release assets. If a smoke test fails, keep the public release notes accurate and fix in a new patch version.
