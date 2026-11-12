# Changelog

## 0.1.4 — 2026-11-12

- Distribute a bundled Companion TGZ for **Modules → Import module package**; ordinary users no longer need Node.js, npm or a development-module folder.
- Separate user installation and upgrade instructions from developer build instructions; update the dashboard and bundled help version labels.
- Validate the packaged entrypoint, identity, bundled runtime and licenses, and record repeatable package checksums.
- Integrate the supplied owner-approved logo into the Windows executable, tray and dashboard without changing the dashboard layout or colors.
- Retain listener behavior, the signed protocol, connection/action IDs, all 38 presets and both migration scripts. The owner confirms GUI import, retained connection settings, a working connection and most existing buttons on the production Companion host.

Owner live Windows testing PASS: EXE, tray, dashboard and existing Companion buttons. See [v0.1.4 validation](docs/VALIDATION-v0.1.4.md) for artifact evidence and coverage limits.

## 0.1.3 — first public release (2026-11-11)

- Reject foreign dashboard authorities/origins and protect pairing-state reads.
- Preserve replay state across settings saves and in-process restarts; serialize settings updates and persist before activating new keys/network settings.
- Preserve legacy Companion expression options; add an indexed scroll repair migration and retain one scroll action with an advanced compatibility mode.
- Reject aborted/oversized HTTP responses and enforce an overall transport deadline.
- Pin Companion base 1.13.6 with an npm lockfile and Node 22 requirements; validate the real installed API in CI.
- Build deterministic local Windows, source and Companion candidate ZIPs with Go 1.27.4, source-commit metadata, SHA-256 checksums and the Go redistribution notice.
- Correct execution-authority claims, installation/rollback instructions and status wording (`OS accepted` does not mean a task completed).

The first public release incorporates these previously completed corrections; publication preparation adds documentation, repository links and final package labeling only. The corrected EXE passed the isolated automated smoke test and a completed Defender scan reporting no threats. Owner-reported working functions and ZORIN module installation are recorded separately from untested cases in SMOKE_TEST.md. Artwork provenance and MIT inclusion are documented. No claim of complete Windows/Companion configuration coverage is made. Preserve existing snapshots and never replace published tags/assets.

## 0.1.3 development snapshot — 2026-11-10 (before first public release)

- Start silently in the Windows notification area. Open the dashboard only from tray (or local URL); prevent a second instance from launching via the existing local-dashboard port.
- Add **Start with Windows** toggle for the current user's Windows Run entry. No administrator privileges or background service required.
- Simplify the UI: only slider-based mouse wheel configuration; checkbox/key dropdown builders for Press, Hold and Release shortcuts. An advanced mode retains complex combinations.
- Add an API 1.x migration script for previously saved button actions using `shortcut_builder`, `scroll`, and text-based shortcuts.
- Include **38 user-selectable drag-and-drop Companion presets** across Windows, Chrome, Editing, Mouse and Settings.
- Replace third-party comparison text in the dashboard with free/open-source product description.
- Preserve the signed command protocol, default API/UI ports, desktop input engine, existing pairing keys and source-IP restriction.

**Verification:** Cross-compiled Windows x64 EXE; Go tests/race checks/static analysis and Companion module tests passed. Windows tray/startup behavior and upgrade migration require a real-user smoke test before public deployment.

## 0.1.2 development snapshot — 2026-11-10

- Working independent listener/module baseline with native input, keyboard and mouse actions, file and URL launching, 20-step scroll limit and dedicated Num Lock control.
- Added native tray, Windows utilities dropdowns, scroll slider and first shortcut builder.
