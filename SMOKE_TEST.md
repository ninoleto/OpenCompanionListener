# v0.1.3 smoke test — Windows + Companion

## Before testing

- Keep working v0.1.2 EXE and an export of your Companion configuration.
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

## Companion module

- [ ] Create one new shortcut with Ctrl+Shift+F23 builder; verify Press/Hold/Release sends expected event types.
- [ ] Check one old shortcut button after upgrade; verify it still works.
- [ ] Check one old scroll button; verify slider action retained movement direction and magnitude.
- [ ] Verify only one Mouse wheel scroll action appears in new action selector.
- [ ] Verify 38 presets appear by categories (Windows, Chrome, Editing, Mouse, Settings), do not automatically populate any page.
- [ ] Drag File Explorer, Next Tab and Scroll Down presets and verify behavior in target applications.
- [ ] Confirm Windows Settings and classic Control Panel dropdowns still work.

If any upgrade/migration test fails, restore the saved Companion configuration and module backup, and switch the EXE back to the confirmed-working version. Collect sanitized Companion error logs; never share pairing keys.
