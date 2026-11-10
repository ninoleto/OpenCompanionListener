# Development handoff — Open Companion Listener

## 2026-11-11 publication-fix continuation — local candidates only

This section supersedes the audit-only and historical next actions below.
Publication remains **on HOLD**; the owner explicitly prohibited GitHub publishing.

- **Branch/commit:** continuing `fix/v0.1.3-publication` from audit commit `5faf3f1` (source baseline `cd2f043`). This session preserves and completes the prior session's uncommitted fixes; the implementation commit containing this section is the packaging checkpoint. Exact candidate commits are recorded in each BUILD-INFO.json.
- **Objective:** retain completed security/migration/transport fixes, finish dependency and packaging work, verify clean builds and document release gates without changing live deployments or preserved snapshots.
- **Files changed:** listener config/main/dashboard and security regressions; module main/transport/upgrades/tests, pinned package/lock and real-base manifest validator; deterministic archive builder/toolchain pins, CI, Git attributes/ignores, Go license/third-party notices, README/security/architecture/install/release/smoke/changelog/release draft and audit/handoff docs.
- **Verified so far:** fresh Windows Go 1.27.4 race run: 15 top-level tests plus 20 key subtests, PASS exit 0; Windows vet with CGO=1 and CGO=0 PASS exit 0. Node 22.23.4 clean `npm ci` installed 17 packages; base 1.13.6 real manifest validation and all 28 Node tests PASS exit 0. npm audit: zero advisories, exit 0. govulncheck v1.9.0 source scan: no vulnerabilities, exit 0 (database updated 2026-11-05). Packaging/binary verification is the next step after the reviewed source commit; later evidence below will record results.
- **Compatibility:** wire format, pairing keys, ports, connection ID, action IDs and all 38 presets retained. Upgrade index 0 is retained with expression-loss prevention; index 1 repairs stranded scroll actions. Expressions already destroyed by an old migration need a pre-upgrade export. Dashboard uses only the canonical numeric loopback URL. Existing app launch behavior is preserved and documented as broad user-session authority.
- **Not performed:** live Companion install/import-export, Windows desktop input, tray observation, Explorer restart, sign-in autostart/registry mutation, Windows 10/11 matrix, independent installation, signing or publication. Tests only use temporary settings and ephemeral network fixtures.
- **Security risks/unresolved:** HMAC is not encryption; executable arguments allow interpreters; replay cache is memory-only across full process exits; in-flight requests can finish under their captured settings. Artwork rights remain unconfirmed. Actual repository/security URLs and prior-publication/version status remain owner decisions. No antivirus clean claim is made. Physical desktop smoke evidence is still required.
- **Rollback:** original v0.1.3 archives/snapshots remain untouched outside this checkout. New builds go under ignored `dist/` directories and refuse overwrite. No real settings, registry values, installed module or Companion configuration was changed. For future deployments restore both old module and matching exported configuration.

**Single next safe action:** build and verify two local candidate sets from the committed packaging checkpoint, preserving the publication hold.

## 2026-11-11 audit handoff — publication on HOLD

This section supersedes the release-readiness claims and next actions in the historical handoff below. The owner requested findings before corrective implementation and prohibited publication. The source audit is complete; the candidate is **not ready for public release**.

- **Branch/commit:** local branch `audit/v0.1.3`; supplied source baseline `cd2f0430d968794007f89facb64cb62f6f08c855`. The incoming workspace had no Git history. This checkout was extracted from the supplied GitHub source ZIP, whose SHA-256 is recorded in [the audit report](docs/AUDIT-v0.1.3.md).
- **Objective:** audit completeness, security, privacy, licensing, buildability, tests, public documentation, and release readiness while preserving working behavior.
- **Changes:** added `docs/AUDIT-v0.1.3.md` and this handoff section. Generated a local untracked module `package-lock.json` during dependency verification; it is evidence, not an approved dependency change. No application/module implementation changed. Parent audit directory contains isolated tools, logs, executable, and defect probes; workspace-root HANDOFF points here.
- **Verification:** Windows/amd64 Go 1.27.4 `go test -race -v ./...` passed (7 top-level tests, 20 key subtests; exit 0). `go vet ./...` passed with both CGO=1 and CGO=0 (exit 0 each). Required GUI Windows build passed (exit 0). Node 22.23.4 original tests passed 16/16. Dependency install, installed-base manifest validation, and npm pack dry run passed. `npm audit` reported zero advisories. Icon-resource regeneration was byte-identical. Existing source/Windows ZIP and inner checksums match; advertised module ZIP is missing.
- **Defect evidence:** added probes fail 3 Go, 3 migration, and 1 transport assertions. Confirmed foreign-Host pairing-key disclosure, replay acceptance after unchanged settings save, active-key change on persistence failure, broken legacy scroll migration/expression metadata loss, and pending transport promises after response abort. See report A01–A06 for precise evidence and fixes.
- **Security/licensing risks:** original EXE embeds Go 1.23.2; binary vulnerability scan reports 67 standard-library advisories (not 67 proven exploits). Fresh Go 1.27.4 source/binary scans report none. The file-launch authority is broader than the README's no-shell claim. Go binary redistribution notices are missing; artwork provenance remains unconfirmed. Public source scan found no credential/personal-path matches; surrounding historical workspace has private deployment material and must not be uploaded wholesale.
- **Not verified:** actual Windows input, tray, sign-in startup, launch completion, real Companion host migration, all presets, and Windows 10/11 matrix. No live user configuration or registry writes performed. Defender scan failed to run (`0x80508023`, exit 2); no clean antivirus claim. Original and rebuilt EXEs are unsigned. No historical upstream commit, release tag, public GitHub URL, or byte-identical reproduction of the original EXE was established.
- **Unresolved:** audit A01–A10, including dependency/engine drift, missing module archive, notices/provenance, public install/build instructions, and manual smoke evidence. No application fixes have been applied because findings are being reported first.
- **Rollback:** original snapshots, EXE, archives, and checksums are untouched. Audit artifacts are isolated and unapproved; no live deployment or publication to revert.

**Single next safe action:** after reviewing the audit, implement and regression-test A01 dashboard authority validation on this branch. Keep publication blocked until the other findings and manual release gates are resolved.

## Historical handoff (retained for context; not current release approval)

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
