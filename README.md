# Open Companion Listener v0.1.3

## Send keystrokes and mouse commands to a Windows PC using Bitfocus Companion.

Open Companion Listener is a free, open-source Windows application that receives commands from Bitfocus Companion and executes them on your computer.

Send individual keystrokes, keyboard shortcuts, mouse clicks, and scroll commands from a Stream Deck, Loupedeck, or any control surface supported by Companion.

You can also launch applications, open files, and execute scripts.

The listener runs quietly in the Windows system tray. Companion can run on the same computer or another machine on your local network.

No subscriptions. No cloud accounts. No license activation.

Free and open source under the MIT license.

**Windows 10/11 x64 · First public release.** Download the Windows listener and Companion module from [GitHub Releases](https://github.com/ninoleto/OpenCompanionListener/releases/tag/v0.1.3). See the [smoke-test record](SMOKE_TEST.md), [audit](docs/AUDIT-v0.1.3.md) and [handoff](HANDOFF.md) for verification and limitations.

## Features

| Feature | Details |
| --- | --- |
| Keyboard | F1–F24, numpad, arrows, navigation, media keys, single key press/hold/release |
| Shortcut builders | Ctrl / Shift / Alt / Win checkboxes and a key dropdown for Press, Hold and Release; advanced fallback for uncommon old combinations |
| Mouse | Click, hold, release and relative movement |
| Mouse wheel | Vertical/horizontal, 1–20 steps slider, Reverse checkbox; 1 is a standard wheel notch |
| Num Lock | Dedicated Toggle / ON / OFF action |
| Files and apps | Open file, folder, EXE or registered URI through Windows Shell |
| Windows utilities | Settings (`ms-settings:`) and Control Panel (`.cpl`) dropdowns |
| Presets | 38 optional drag-and-drop buttons (Windows, Chrome, Editing, Mouse, Settings) |
| System tray | Double-click/Open GUI, Restart Listener, Exit |
| Windows startup | ON/OFF switch in dashboard Settings; per-user startup, no administrator rights needed |
| Connection | HMAC-SHA256 signed commands, request timestamps, nonce replay checks, trusted source-IP restriction |

## Windows installation

1. Download `OpenCompanionListener-v0.1.3-Windows-x64.zip` and `SHA256SUMS.txt` from the release page. Compare `Get-FileHash -Algorithm SHA256 .\OpenCompanionListener-v0.1.3-Windows-x64.zip` with the listed checksum. Extract to a **permanent folder** (for example `C:\Apps\OpenCompanionListener`). Do not run the EXE from inside the ZIP.
2. Run `OpenCompanionListener.exe`. **No browser window opens automatically.** Find the icon in the Windows notification area, right-click **Open GUI**, or double-click the icon. The dashboard is `http://127.0.0.1:17993/` on this PC only.
3. In Settings, set the Windows PC's **Listen IP**, the **Trusted Companion server IP**, and API port (default **17992**). Save settings.
4. Copy the 64-character pairing key from Settings to the custom Companion connection. Do not publish or screenshot it.
5. Use Windows Firewall to allow inbound TCP **17992 only from the trusted Companion host**, on your private network. Never forward the port to the Internet.
6. Optionally enable **Start with Windows**. It creates/removes an entry under the current Windows user's `Run` registry key. This launches silently when you sign in, does not require admin rights, and uses the current EXE's full path. Move the EXE to its final location **before** enabling autostart.

Closing the browser tab does not stop the listener. Use the system tray's **Exit** command to shut it down.

### Upgrading Windows from 0.1.2

Exit the previous application from its tray menu, save its EXE as a backup, extract the new build and run the new EXE. Your existing settings and pairing key under `%LOCALAPPDATA%\OpenCompanionListener\settings.json` are retained. Startup remains opt-in; verify the existing entry and its path before changing it.

**Important:** This version is not code-signed. Windows SmartScreen may warn about an unrecognized publisher. Review the source/checksum. The new startup registry switch requires a real Windows desktop smoke test before production use.

## Bitfocus Companion module installation

The distribution includes the **module source folder and dependency lockfile**:

`companion-dev-modules/companion-module-community-openlistener/`

Copy this folder to the *developer / extra module parent path* configured in Bitfocus Companion. It is not a ZIP-based in-app upload.

Use Companion 4.1+ with API 1.13/node22 support and Node 22.8+ within the 22.x line. This is the API minimum. The owner reports successful corrected-module installation on ZORIN with 28 passing tests; full host/migration coverage is not claimed. In a staging copy, run `npm ci --omit=dev --ignore-scripts --no-audit --no-fund`, `npm run validate`, then `npm test`, stopping on any failure. Dependencies are downloaded during installation; normal use needs no cloud connection.

Follow [first-install, upgrade and rollback instructions](docs/COMPANION-INSTALL.md). Export the Companion configuration before upgrades and stop the test host before replacing its module. Reload only when safe; restarting Companion interrupts connections.

### Companion actions

Use **Open Companion Listener → Add action**, or the **Presets** tab for ready-made buttons. Keyboard shortcuts use checkbox/dropdown builders, and scrolling uses a slider and Reverse checkbox. The old manual shortcut and scrolling action IDs are converted by the included **API 1.x upgrade script**; an advanced text mode remains available for combinations that cannot be represented by a single key dropdown. Back up the Companion configuration prior to updating and verify existing buttons after restart.

**Browser presets** are shortcuts delivered to the *currently focused window*; they do not focus Chrome automatically. Presets never auto-populate button pages.

## Troubleshooting

- **No tray icon:** Check the Windows hidden notification icons. The dashboard is reachable locally at `http://127.0.0.1:17993/` while the listener is running.
- **Cannot connect:** Check IP, port `17992`, trusted IP, pairing key, firewall and both computers' clocks.
- **Can't send keys to elevated apps:** Windows blocks lower-integrity processes from interacting with elevated windows. This app runs as a normal interactive user and does not control the Windows lock screen.
- **Existing button stopped working after a module update:** Restore your Companion configuration backup and the previous module folder. Report the action ID and sanitized options.
- **A Windows Settings page won't open:** Some `ms-settings:` pages depend on Windows edition, version, hardware or administrator policies.

## Security

Remote commands are sent via HTTP with an HMAC-SHA256 signature, timestamp and nonce. **Payloads are authenticated but not encrypted.** Use a trusted LAN or private overlay and restrict incoming traffic using Windows Firewall. The pairing key is stored in a local user configuration file. Do not expose the API publicly. Pairing grants desktop-user execution authority: the file/app launcher accepts executables and arguments, including interpreters. The absence of a dedicated shell action is not an execution sandbox. See [SECURITY.md](SECURITY.md), including the full-process-restart replay limitation.

The local dashboard is loopback-only (`127.0.0.1:17993`). Windows startup uses the per-user registry and launches only after the user signs in; it does **not** run as a Windows service.

## Developer build

Use Go **1.27.4**, Node **22.23.4**, and the checked-in dependency lockfile. The Go language directive is a language minimum, not the release compiler selection. Windows race tests require CGO and a compatible 64-bit MinGW C compiler on PATH; the release EXE uses `CGO_ENABLED=0`.

PowerShell 7, from the repository root (stop if any command fails):

```powershell
$env:CGO_ENABLED = '1'
Push-Location listener
go test -race ./...
go vet ./...
$env:GOOS = 'windows'; $env:GOARCH = 'amd64'; $env:CGO_ENABLED = '0'
go vet ./...
go build -trimpath -buildvcs=false -ldflags='-s -w -H windowsgui' -o ../OpenCompanionListener.exe .
Pop-Location
Push-Location companion-dev-modules/companion-module-community-openlistener
npm ci --omit=dev --ignore-scripts --no-audit --no-fund
npm run validate
npm test
Pop-Location
```

Bash:

```bash
set -eu
cd listener
go test -race ./...
go vet ./...
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go vet ./...
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -buildvcs=false -ldflags='-s -w -H windowsgui' -o ../OpenCompanionListener.exe .

cd ../companion-dev-modules/companion-module-community-openlistener
npm ci --omit=dev --ignore-scripts --no-audit --no-fund
npm run validate
npm test
```

The Go module has no external dependencies. Python 3.11+ is needed only for deterministic archive packaging or optional icon regeneration, never to run the listener. From clean committed source, run `python tools/package_release.py --output dist/candidate-1`; see [release instructions](docs/RELEASING.md) for reproduction and checksum verification. PowerShell users with the Python launcher can use `py -3.13` in place of `python`.

## License

Project code and project artwork are MIT licensed. See [third-party notices](THIRD_PARTY_NOTICES.md) for the Go binary license and the owner-confirmed artwork provenance. Community contributions are welcome. The future Linux backend and optional visual automation builder are not implemented yet.

## Forking and contributing

The repository contains both the standalone Windows listener (`listener/`) and the separate Companion module (`companion-dev-modules/companion-module-community-openlistener/`). Contributors can fork either component or maintain them together, subject to the MIT license.

- **Start here:** [`AGENTS.md`](AGENTS.md) for coding-agent and developer guardrails; [`HANDOFF.md`](HANDOFF.md) for the current implementation, limitations and next steps.
- **Understand the system:** [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md) for API/signature format and native controls.
- **Contribute:** [`CONTRIBUTING.md`](CONTRIBUTING.md) and [`SECURITY.md`](SECURITY.md).
- **Publish patches:** [`docs/RELEASING.md`](docs/RELEASING.md). Report bugs in GitHub Issues with sanitized logs; never share the pairing key.

Bug fixes are published as **new versioned releases**, preserving previous tags and binaries. Windows is supported today; Linux and a visual automation builder are potential future contributions, not current features.
