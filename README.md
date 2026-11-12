# Open Companion Listener

Open Companion Listener receives commands from Bitfocus Companion and sends keystrokes, keyboard shortcuts, mouse clicks, movement and scrolling commands to a Windows PC. It can also open applications, files, URLs and scripts through Windows file associations or an installed interpreter.

The listener runs in the Windows notification area. Companion can run on the same PC or another supported Companion host. The project is free and open source under the MIT license; normal operation is local and requires no project account or cloud service.

**v0.1.4 is in validation and has not been published.** See [release verification](docs/VALIDATION-v0.1.4.md) for completed and pending checks. Public downloads are listed on [GitHub Releases](https://github.com/ninoleto/OpenCompanionListener/releases).

## Requirements

- Windows 10/11 x64 on the PC being controlled, signed in as a normal user.
- Bitfocus Companion 4.1 or later with the node22 module runtime. This is the API minimum; tested host versions are recorded in the verification document.
- A trusted local network or private overlay if Companion runs on another computer.

End users do not need Node.js, npm, Python or development tools. Companion supplies its module runtime.

Download `SHA256SUMS.txt` with the release files. On Windows, compare each file's `Get-FileHash -Algorithm SHA256 <filename>` result with its entry before installation. `BUILD-INFO.json` records the source commit and build details; the source ZIP is optional for ordinary users.

## Install the Windows listener

1. Download `OpenCompanionListener-v0.1.4-Windows-x64.zip` from the release once available. Extract it to a permanent folder and run `OpenCompanionListener.exe` outside the ZIP.
2. Find the notification-area icon, including under hidden icons. Double-click it or choose **Open GUI** from its menu. The dashboard does not open automatically.
3. Open **Settings**. If Companion runs on this PC, leave both addresses at `127.0.0.1`. Otherwise enter this Windows PC's IPv4 address as **Listen IP address** and the Companion host's IPv4 address as **Trusted Companion server IP**. Leave the API port at `17992` unless you need a different port, then save.
4. For a remote Companion host, allow that API port through Windows Firewall on the private network, restricted to the Companion host's address. Never forward it to the Internet.
5. Copy the pairing key for the Companion connection below. Keep it private.

The dashboard is available on the Windows PC at `http://127.0.0.1:17993/`. Closing its browser tab leaves the listener running; use the tray's **Exit** command to stop it. **Restart Listener** restarts the API. **Start with Windows** is optional and starts the listener when this user signs in. Place the EXE in its permanent location before enabling it.

## Install the Companion module

1. Download `community-openlistener-0.1.4.tgz` from the same release. Leave the file compressed.
2. Open Companion's web interface and select **Modules**.
3. Select **Import module package**, choose the `.tgz` file and confirm the import.
4. Open **Connections** and add **Open Companion Listener** (manufacturer **Community**).
5. Enter **Listener IP address**, **Listener port** and **Pairing key** from the Windows listener settings. Save and check that the connection becomes OK.
6. Add actions to your buttons, or open **Presets** and drag the buttons you want onto a page.

See [module installation, upgrades and troubleshooting](docs/COMPANION-INSTALL.md). Export your Companion configuration before upgrading. Existing v0.1.3 development-folder installations need the separate transition procedure in that guide; keep the existing connection so its buttons retain their references.

## Controls

| Control | Available actions |
| --- | --- |
| Keyboard | Single keys, F1–F24, numpad, navigation and media keys; press, hold and release |
| Shortcuts | Modifier checkboxes and a key dropdown for Press, Hold and Release; advanced text for uncommon existing combinations |
| Mouse | Left/right/middle click, hold/release, relative movement |
| Scrolling | Vertical/horizontal, 1–20 notches, separate Reverse control |
| Num Lock | Toggle, ON, OFF |
| Opening items | Applications with optional arguments, files, folders, URLs and scripts using an installed interpreter or file association |
| Windows utilities | Settings and Control Panel dropdowns; availability depends on Windows version, hardware and policy |
| Presets | 38 optional buttons in Windows, Chrome, Editing, Mouse and Settings categories |

Input goes to the focused Windows application. Browser presets do not launch or focus Chrome. Installing presets adds no buttons automatically. An “OS accepted” event means Windows accepted the request; it does not confirm that an application finished opening or completed its task.

## Troubleshooting and limits

- **Import unavailable:** check the Companion version and use its local web interface on the Companion host. Newer Companion versions restrict package imports from remote clients.
- **Connection failure:** check the two IP addresses, API port, pairing key, firewall and clocks on both computers. Use IPv4 addresses, not computer names. The dashboard port is different from the API port.
- **No input:** focus the intended application. Windows security restrictions prevent control of some elevated windows and locked desktops. Do not run the listener as Administrator to work around them.
- **Held key or mouse button:** send the matching Release action. Automatic release after an interrupted connection is not implemented.
- **Windows warning:** the executable is unsigned. Download from the project release and consult its checksums before deciding whether to run it.
- **Upgrade problem:** use the matching module and configuration backup described in the installation guide. Report sanitized steps and versions in [GitHub Issues](https://github.com/ninoleto/OpenCompanionListener/issues).

## Security

Pairing gives the Companion host authority to control your desktop and launch code as your Windows user. Commands are authenticated with HMAC-SHA256 but **are not encrypted**. Use trusted networks and restrict the API with the firewall. Keep the pairing key, configuration exports and `%LOCALAPPDATA%\OpenCompanionListener\settings.json` private. The settings dashboard stays on loopback. Read [SECURITY.md](SECURITY.md) for replay-protection and other limits.

## Development and project records

- [Developer guide](docs/DEVELOPMENT.md): tools, builds and source-module development.
- [Contributing](CONTRIBUTING.md), [architecture](docs/ARCHITECTURE.md) and [release process](docs/RELEASING.md).
- [Agent instructions](AGENTS.md) and [current handoff](HANDOFF.md) for coding agents continuing the work.
- [Changelog](CHANGELOG.md), [smoke checklist](SMOKE_TEST.md), [MIT license](LICENSE) and [third-party notices](THIRD_PARTY_NOTICES.md).
