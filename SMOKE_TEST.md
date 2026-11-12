# Smoke test — Windows + Companion

## v0.1.4 owner report — 2026-11-12

The owner confirms that the v0.1.4 TGZ imported through Companion's GUI on the
production Companion host, retained existing connection settings, connected,
and ran most existing buttons successfully. Exact installed artifact hashes,
host versions, individual button coverage, all presets and rollback were not
specified. This is real-host evidence; it is not an isolated-profile test.
Do not repeat completed tests or infer that every procedure below passed.
See [v0.1.4 verification](docs/VALIDATION-v0.1.4.md) for current release gates.

## Historical v0.1.3 evidence

**Candidate status: partial owner-reported manual verification.** Automated
regression results are recorded separately in HANDOFF.md. The detailed checklist
remains unchecked where its complete procedure has not been observed.

On 2026-11-11 the limited executable smoke tool passed (exit 0): candidate start,
local dashboard/authority checks, signed authentication, save/restart replay
rejection, duplicate-instance exit, temporary-key rotation and clean shutdown.
The installed EXE, current-user settings and startup value were unchanged.
Visible tray behavior and the detailed interactive procedures below have not
been independently verified by the agent.

## Owner report and artifact identity — 2026-11-11

The owner confirmed the following working; the agent did not repeat these
interactive tests or inspect the ZORIN host during this session:

| Owner-reported check | Reported result and scope |
| --- | --- |
| Num Lock | Working; individual ON/OFF/idempotence cases were not specified. |
| Keyboard shortcuts and mouse controls | Working; no claim that all hold/release, interruption or boundary cases passed. |
| Windows display switching | Working using `DisplaySwitch.exe` arguments 1–4 on the owner's setup. |
| AutoHotkey scripts | Script launching working. |
| Voicemeeter recorder and VBAN feedback | Owner's workflow working; this is not a claim of a new built-in listener integration. |
| Corrected Companion module on ZORIN | Owner confirms installation and 28 passing tests; host logs, installed module hashes and Companion version were not independently collected. |

Read-only process inspection found the currently running Windows listener in
the extracted corrected candidate directory. Its SHA-256 is
`510696b6fba8331bfef6d929013213fea15f43f3dde78f7db34c8370dac97f83`, identical
to the candidate Windows ZIP, BUILD-INFO.json and the EXE used for the earlier
isolated automated smoke test. The preserved original v0.1.3 installation has
a different hash (`d9a346c96849c557cb2f75870dce34c68e157a9b02813426b6ddcef88fc21ff1`).
Current process identity is verified; the owner has been asked to identify the
executable used at manual-test time. A current process hash alone does not prove
which executable was used for every earlier test.

Microsoft Defender completed a custom, non-remediating scan of the corrected
EXE on 2026-11-11, reported no threats, and returned HRESULT/exit 0. The file was
not skipped and its hash was unchanged. Product `4.18.26100.4`, engine
`1.1.26100.4`, signatures `1.463.243.0` (updated 2026-11-11 09:41:12 UTC;
Defender reports not out of date). This is a point-in-time scan of that EXE,
not a guarantee or a scan of all dependencies. See the audit for exact evidence.

Record candidate source commit and ZIP/EXE SHA-256, Windows edition/build,
Companion version, tester and date before testing. Use an isolated test account
and exported/copied Companion configuration; never unreviewed code on live buttons.

## Before testing

- Perform interactive tests in a separate Windows test account/VM and isolated
  Companion profile/host. Launching the candidate normally in the production
  Windows account uses its real settings. Do not do that for this checklist.
- `tools/smoke_listener.cjs` supplies its own temporary LOCALAPPDATA and test key;
  it is suitable for the limited automated check after the working listener has
  been exited by the user. It performs no desktop input or startup writes.
- Changing LOCALAPPDATA does **not** isolate the HKCU startup registry key.
  Test Start with Windows only in the separate test account/VM; leave the
  production user's startup value and paired devices untouched.
- Prepare ZORIN with [the isolated upgrade runbook](docs/ZORIN-VALIDATION.md).

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
- [x] Record a successful trusted antivirus scan: corrected EXE, Defender versions/result above; earlier failed/skipped attempts remain historical.
- [x] Record the artwork origin and redistribution permission: owner confirmation and byte-level verification in THIRD_PARTY_NOTICES.md.
- [x] Owner identifies v0.1.3 as the first public release; preserve earlier development snapshots and never replace a published tag or asset.

## Detailed operator checks

Mark each result only after observing it. Record Windows build, Companion
version, candidate commit/hash, tester/date, and failed/skipped cases separately.
Use an unsaved scratch document and disposable browser tabs. Keep the focus on
the intended test window and release all held modifiers/buttons after each case.

| Area | Procedure and expected observation | Result |
| --- | --- | --- |
| Tray start | Launch in the test account; one icon, no automatic browser. A second launch leaves one process/icon. | Pending |
| Tray controls | Double-click and Open GUI reach the local dashboard. Restart retains one process, settings and connection; Exit removes icon and closes API/UI ports. | Pending |
| Explorer recovery | In the test VM only, restart Explorer; icon returns and controls still work. | Pending |
| Startup enable | Enable in the test account; refresh confirms enabled. Sign out/in with candidate in a permanent test folder; it starts silently with one tray icon. | Pending |
| Startup disable | Disable; refresh confirms off. Next sign-in does not launch it. Production user's startup value remains untouched. | Pending |
| Num Lock | Record initial state. ON twice stays on; OFF twice stays off; Toggle reverses each time. Observe the indicator and keypad behavior, then restore the initial state. | Pending |
| Keyboard builder | Test Ctrl+A in a scratch document, Ctrl+Shift+F23 in an appropriate key/event viewer, and a plain function/numpad key. Check selected modifiers/key, Press, Hold, Release and reverse-order modifier release. | Pending |
| Advanced shortcuts | Test an old unusual combination, an expression and a migrated simple shortcut. Options survive save/restart; no unexpectedly stuck modifiers. | Pending |
| Mouse buttons/move | Test left/right/middle click, hold/release and relative movement on a scratch target; no held button remains afterward. | Pending |
| Vertical scroll | Test magnitudes 1, 5, 20 and Reverse both ways in a long document. Verify direction reverses and magnitude increases; applications may render notches differently. | Pending |
| Horizontal scroll | Repeat on a horizontally scrollable target; verify axis and both directions. Check an old signed scroll action retains its sign/magnitude. | Pending |
| Scroll validation | Confirm a single scroll action in the selector, slider range 1–20, and migrated advanced invalid values reject without moving the target. | Pending |
| Existing buttons | Before upgrade record IDs/locations/options for shortcut, hold/release, scroll, Num Lock, file/URL and utility buttons. After migration and a second restart, compare and execute those same copied buttons. | Pending |
| Preset opt-in | Exactly 38 presets appear in five categories. Installation adds no buttons/pages. Drag to a scratch page only. | Pending |
| Launch actions | Open a harmless test URL/file/app and Windows Settings/Control Panel entries. Observe the GUI independently; OS accepted does not mean task finished. | Pending |
| Interruption | In isolation, interrupt a held-key/button connection, reconnect and release. Record any manual recovery needed; automatic release is not assumed. | Pending |
| Boundaries | Verify normal focus behavior and expected inability to control elevated/locked desktops. Do not elevate to bypass the restriction. | Pending |
| Rollback | Restore the original module parent and matching test config backup; existing copied buttons and device assignments return. | Pending |

Preset inventory to exercise (record a result for each, including unavailable
version-dependent Settings pages):

| Category | Count | Presets |
| --- | --- | --- |
| Windows | 12 | File Explorer; Show Desktop; Windows Settings; Task Manager; Task View; Screen Snip; Run Dialog; Clipboard History; Snap Left; Snap Right; Project Display; Emoji Picker |
| Chrome | 12 | New Tab; Close Tab; Reopen Tab; Next Tab; Previous Tab; Address Bar; Downloads; History; Incognito; Refresh; Find in Page; New Window |
| Editing | 8 | Copy; Paste; Cut; Undo; Redo; Select All; Save; Find |
| Mouse | 2 | Scroll Up; Scroll Down |
| Settings | 4 | Sound Settings; Volume Mixer; Display Settings; Network Settings |

Chrome presets act on the focused window; they do not focus or launch Chrome.
Use disposable content for Close Tab/Cut/Save and a test display configuration
for Project Display. Do not mark all presets passed after checking only three.

If any upgrade/migration test fails, restore the saved Companion configuration and module backup, and switch the EXE back to the confirmed-working version. Collect sanitized Companion error logs; never share pairing keys.
