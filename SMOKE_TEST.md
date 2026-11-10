# v0.1.3 smoke test — Windows + Companion

**Candidate status: not yet manually verified.** Automated regression results
are recorded separately in HANDOFF.md; they do not check these boxes.

Record candidate source commit and ZIP/EXE SHA-256, Windows edition/build,
Companion version, tester and date before testing. Use an isolated test account
and exported/copied Companion configuration; never unreviewed code on live buttons.

## Before testing

- Keep the original v0.1.3 snapshots (and any confirmed-working older EXE) and an export of your Companion configuration.
- Back up the existing Companion module directory.
- Do not overwrite an existing release in place during first testing.

## Windows listener

- [ ] Launch OpenCompanionListener.exe; verify **no browser window automatically opens**.
- [ ] Verify one system-tray icon. Launch the EXE a second time; verify no duplicate process/tray icon.
- [ ] Double-click tray icon; dashboard opens. Right-click menu → Open GUI / Restart Listener / Exit.
- [ ] In dashboard Settings, enable Start with Windows; verify checkbox persists after refresh.
- [ ] With EXE in a permanent location, sign out/in; listener starts with tray icon and **no browser window**.
- [ ] Disable Start with Windows; confirm autostart entry removed using Windows Task Manager Startup apps or Registry Editor.
- [ ] Verify network/pairing settings from previous version remain unchanged.
- [ ] Confirm standard keyboard/mouse, scroll, URL/file launching and Num Lock continue working.
- [ ] Confirm dashboard loads via `127.0.0.1:17993`, state refresh works, settings persist, and key rotation requires updating the Companion key.
- [ ] Restart from tray and confirm the API reconnects without a second process.
- [ ] Restart Explorer on an isolated test desktop; tray icon returns.
- [ ] Test modifier/mouse hold-release and interrupted connections; record any held inputs requiring manual release.
- [ ] Verify focus/UIPI/locked-desktop limits; no elevation bypass is expected.
- [ ] Check launch logs say OS accepted, and independently observe the application opening; do not treat API success as task completion.

## Companion module

- [ ] Create one new shortcut with Ctrl+Shift+F23 builder; verify Press/Hold/Release sends expected event types.
- [ ] Check one old shortcut button after upgrade; verify it still works.
- [ ] Check one old scroll button; verify slider action retained movement direction and magnitude.
- [ ] Import/export a copied legacy configuration, including raw/wrapped values, numeric-looking expressions, unusual shortcuts and signed scroll expressions. Verify migration and execution after a second restart.
- [ ] Verify the additive repair on an old `scroll` action whose upgrade index is already 0; preserve expression metadata and direction/magnitude.
- [ ] Verify only one Mouse wheel scroll action appears in new action selector.
- [ ] Verify 38 presets appear by categories (Windows, Chrome, Editing, Mouse, Settings), do not automatically populate any page.
- [ ] Drag File Explorer, Next Tab and Scroll Down presets and verify behavior in target applications.
- [ ] Confirm Windows Settings and classic Control Panel dropdowns still work.
- [ ] Exercise all 38 presets where applicable; record OS-specific unavailable Settings pages and expected labels.
- [ ] Restore the backed-up module **and** matching configuration export and verify rollback.

## Release evidence

- [ ] Repeat applicable desktop checks on Windows 10 x64 and Windows 11 x64; record builds and deviations.
- [ ] Record a successful trusted antivirus scan (engine/signature versions and result). A failed scan is not a clean result.
- [ ] Record the artwork origin and redistribution permission.
- [ ] Confirm whether v0.1.3 has already been distributed; use a new patch version when required.

If any upgrade/migration test fails, restore the saved Companion configuration and module backup, and switch the EXE back to the confirmed-working version. Collect sanitized Companion error logs; never share pairing keys.
