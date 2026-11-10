# AGENTS.md — Open Companion Listener

Read this file **first** before modifying any application or Companion module code. Then read `HANDOFF.md`, `docs/ARCHITECTURE.md`, `SECURITY.md`, and `SMOKE_TEST.md`. This repository is designed to be understandable and forkable by human contributors and AI coding agents.

## Mission and scope

- Deliver **free, open-source, local-first** desktop control from **Bitfocus Companion** to Windows 10/11 x64. No VICREO dependency, paid license, account, or cloud connection.
- **Windows is the supported desktop target in v0.1.3.** Keep interfaces adaptable to Linux later, but do not claim Linux input support or add an unreviewed cross-platform implementation.
- The listener is a small **native Go Windows executable** and the Companion connection module is **Node.js/CommonJS**. Avoid Electron, embedded browser frameworks, separate Python interpreters, and large new dependencies unless justified with measurements.
- Improve **real-user behavior**, not just compilation. Windows API correctness and Companion runtime behavior need hands-on tests.

## Compatibility contracts (preserve unless an explicit versioned migration is approved)

1. **Wire format:** `GET /v1/status`, `POST /v1/command` on the configurable API port (default `17992`); request body and HMAC signing algorithm described in `docs/ARCHITECTURE.md`.
2. **Authentication:** HMAC-SHA256, 32-hex-character nonce, UNIX-second timestamp within ±60 seconds, replay rejection, trusted TCP peer IP. Never log, paste, or commit real pairing keys. Keep the UI bound to loopback (`127.0.0.1:17993`) by default.
3. **Companion connection ID:** `community-openlistener` and published action IDs must remain stable unless a tested **Companion API 1.x upgrade script** migrates saved configurations. Do not delete previously published upgrade scripts.
4. **User controls:** Preserve native keyboard/mouse press, hold, release, scrolling, Num Lock, file/URL/app opening, Windows Settings/Control Panel, all released presets, and existing configured Companion buttons.
5. **Tray:** Start silently; no automatic browser dashboard. Open GUI, Restart Listener, Exit; avoid duplicate instances. Per-user Windows autostart must remain opt-in and must not require elevation.
6. **Security boundary:** Never add unrestricted network-to-shell execution or expose the local settings UI on LAN/Internet without a separate security design and explicit tests. Windows input injection still obeys OS security boundaries; do not bypass them.

   Existing `open_file` executable/argument launching already permits interpreters and registered URI handlers. Treat pairing as full desktop-user authority, not a restricted launcher. Preserve existing launching behavior; the rule above prohibits adding new shell endpoints and must not be used to claim existing authenticated actions cannot execute code.

## Source map

| Area | Files |
| --- | --- |
| TCP/API server, dashboard, lifecycle | `listener/main.go`, `listener/ui/` |
| Configuration and pairing key | `listener/config.go` |
| Request signing and replay | `listener/auth.go`, `listener/auth_test.go` |
| Command validation | `listener/commands.go` |
| Windows input, file and URL execution | `listener/desktop_windows.go`, `listener/keys.go` |
| Tray and autostart | `listener/tray_windows.go`, `listener/startup_windows.go` |
| Non-Windows compile stubs only | `listener/desktop_stub.go`, `listener/startup_stub.go` |
| Companion actions and presets | `companion-dev-modules/companion-module-community-openlistener/src/main.cjs`, `src/presets.cjs` |
| Companion signed transport | `.../src/transport.cjs` |
| Companion saved-action migration | `.../src/upgrades.cjs`, `src/migration.test.cjs` |
| Release validation | `SMOKE_TEST.md`, `docs/RELEASING.md` |

## Required workflow

1. Record the issue, reproduce it, and identify the smallest affected module. Do not start by rewriting working code.
2. Make changes on a branch. Keep each change focused. Do not modify saved real-user configuration or a live Companion installation to test unreviewed code.
3. Add or update automated regression tests for any altered protocol, action, migration, authentication, or validation behavior.
4. Run:

   ```sh
   ( cd listener && go test -race ./... && go vet ./... )
   ( cd companion-dev-modules/companion-module-community-openlistener && node --test src/*.test.cjs )
   ( cd listener && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go vet ./... )
   ( cd listener && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags='-s -w -H windowsgui' -o ../OpenCompanionListener.exe . )
   ```

5. Perform the Windows + Companion smoke tests in `SMOKE_TEST.md` for every release candidate. Cross-build tests **do not** prove actual input, tray, or startup behavior.
6. Record changed behavior, compatibility risks, test evidence, unresolved issues, and rollback path in `HANDOFF.md`; update `CHANGELOG.md` and README as needed.
7. For publishable changes, follow `docs/RELEASING.md` and do not claim a release has been shipped until artifacts have been uploaded and independently installed.

## Action and UI design

- Offer **dropdowns, checkboxes, and sliders**, not duplicate manual-entry actions. Keep existing button values working via migrations; don't silently discard legacy advanced combinations.
- Use `src/upgrades.cjs` to migrate older action IDs and option shapes; test raw and wrapped Companion option values, empty fields, old expression values, and unusual shortcuts.
- Keep scroll values clamped to 1–20 notches with separate direction; preserve existing direction and magnitude during migration.
- When adding Windows Settings pages or presets, validate platform behavior and label version-specific entries appropriately. Presets should be **opt-in drag-and-drop**, never alter existing Companion pages on installation.
- Distinguish **accepted by OS**, **child process successfully started**, and **task finished** in status/log text. Do not equate a shell exit code of zero with successful downstream GUI action.

## Security and release guardrails

- The HMAC authenticates requests but **does not encrypt** traffic. Document trusted-LAN/private-overlay and firewall requirements. Never publish personal hostnames, local paths, production passwords, account identifiers, or real secrets in examples or logs.
- Keep per-user keys and `settings.json` outside the repository. Avoid storing real deployment credentials in test fixtures.
- Never overwrite an existing published tag or silently replace release binaries. Make a new patch release with a changelog and checksums.
- Contributions should retain the MIT license and any required upstream attributions.

## Handoff format

At the end of any significant agent session, update `HANDOFF.md` with: branch/commit, objective, files changed, tests run with exact results, verification not performed, security risks, unresolved bugs, and **the single next safe action**. Avoid claiming tasks completed without code or tests supporting them.
