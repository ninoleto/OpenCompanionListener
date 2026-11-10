# Changelog

## 0.1.3 — 2026-11-10

- Start silently in the Windows notification area. Open the dashboard only from tray (or local URL); prevent a second instance from launching via the existing local-dashboard port.
- Add **Start with Windows** toggle for the current user's Windows Run entry. No administrator privileges or background service required.
- Simplify the UI: only slider-based mouse wheel configuration; checkbox/key dropdown builders for Press, Hold and Release shortcuts. An advanced mode retains complex combinations.
- Add an API 1.x migration script for previously saved button actions using `shortcut_builder`, `scroll`, and text-based shortcuts.
- Include **38 user-selectable drag-and-drop Companion presets** across Windows, Chrome, Editing, Mouse and Settings.
- Replace third-party comparison text in the dashboard with free/open-source product description.
- Preserve the signed command protocol, default API/UI ports, desktop input engine, existing pairing keys and source-IP restriction.

**Verification:** Cross-compiled Windows x64 EXE; Go tests/race checks/static analysis and Companion module tests passed. Windows tray/startup behavior and upgrade migration require a real-user smoke test before public deployment.

## 0.1.2 — 2026-11-10

- Working independent listener/module baseline with native input, keyboard and mouse actions, file and URL launching, 20-step scroll limit and dedicated Num Lock control.
- Added native tray, Windows utilities dropdowns, scroll slider and first shortcut builder.
