# Open Companion Listener

**Version 0.1.3 — Windows release** · Free and open source · MIT License

Open Companion Listener is a lightweight Windows desktop-control application with its **own Bitfocus Companion connection module**. Control Windows keyboard, mouse, scrolling, file and application launching from a Stream Deck, Loupedeck, or any surface supported by Companion. No subscription, cloud account, Python runtime, or license activation is required.

**Windows 10/11 x64** is supported. The listener runs in the Windows notification area and presents a dark configuration dashboard in your browser **only when you request it**. It does not open the dashboard when the application starts. Linux desktop control is not available in this release.

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

1. Download `OpenCompanionListener-v0.1.3-Windows-x64.zip` and extract it to a **permanent folder** (for example `C:\Apps\OpenCompanionListener`). Do not run the EXE from inside the ZIP.
2. Run `OpenCompanionListener.exe`. **No browser window opens automatically.** Find the icon in the Windows notification area, right-click **Open GUI**, or double-click the icon. The dashboard is `http://127.0.0.1:17993/` on this PC only.
3. In Settings, set the Windows PC's **Listen IP**, the **Trusted Companion server IP**, and API port (default **17992**). Save settings.
4. Copy the 64-character pairing key from Settings to the custom Companion connection. Do not publish or screenshot it.
5. Use Windows Firewall to allow inbound TCP **17992 only from the trusted Companion host**, on your private network. Never forward the port to the Internet.
6. Optionally enable **Start with Windows**. It creates/removes an entry under the current Windows user's `Run` registry key. This launches silently when you sign in, does not require admin rights, and uses the current EXE's full path. Move the EXE to its final location **before** enabling autostart.

Closing the browser tab does not stop the listener. Use the system tray's **Exit** command to shut it down.

### Upgrading Windows from 0.1.2

Exit the previous application from its tray menu, save its EXE as a backup, extract the new build and run the new EXE. Your existing settings and pairing key under `%LOCALAPPDATA%\OpenCompanionListener\settings.json` are retained. Windows startup is OFF until you enable it in the dashboard.

**Important:** This version is not code-signed. Windows SmartScreen may warn about an unrecognized publisher. Review the source/checksum. The new startup registry switch requires a real Windows desktop smoke test before production use.

## Bitfocus Companion module installation

The distribution includes the **module source folder**:

`companion-dev-modules/companion-module-community-openlistener/`

Copy this folder to the *developer / extra module parent path* configured in Bitfocus Companion. It is not a ZIP-based in-app upload.

Example for a Linux Companion host with a configured developer module path:

```bash
# Back up the currently installed module before replacing files.
SRC="/path/to/extracted/companion-module-community-openlistener"
DEST="$HOME/companion-module-active/companion-module-community-openlistener"

test "$(node -e "console.log(require(process.argv[1]).version)" "$SRC/package.json")" = "0.1.3" || exit 1
cp -a "$DEST" "$HOME/openlistener-module-backup-$(date +%Y%m%d-%H%M%S)"
cp -a "$SRC/." "$DEST/"
cd "$DEST"
npm install --omit=dev
node --test src/*.test.cjs
```

Reload Companion only when safe: restarting the host service temporarily interrupts your existing Companion connections. On systemd installations it may be `sudo systemctl restart companion.service`. Keep an export/backup of the Companion configuration before updating a module with existing buttons.

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

Remote commands are sent via HTTP with an HMAC-SHA256 signature, timestamp and nonce. **Payloads are authenticated but not encrypted.** Use a trusted LAN or private overlay and restrict incoming traffic using Windows Firewall. The pairing key is stored in a local user configuration file. Do not expose the API publicly. Arbitrary remote shell execution is not supported.

The local dashboard is loopback-only (`127.0.0.1:17993`). Windows startup uses the per-user registry and launches only after the user signs in; it does **not** run as a Windows service.

## Developer build

The Go module has no external Go dependencies.

```bash
cd listener
go test -race ./...
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go vet ./...
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags='-s -w -H windowsgui' -o ../OpenCompanionListener.exe .

cd ../companion-dev-modules/companion-module-community-openlistener
node --test src/*.test.cjs
```

The Windows PE icon resource is generated from the included `.ico` file. The full application and module source, MIT license, and developer handoff are included.

## License

MIT License. Free to use, fork, modify, and redistribute under the license terms. Community contributions are welcome. The future Linux backend and optional visual automation builder are not implemented yet.

## Forking and contributing

The repository contains both the standalone Windows listener (`listener/`) and the separate Companion module (`companion-dev-modules/companion-module-community-openlistener/`). Contributors can fork either component or maintain them together, subject to the MIT license.

- **Start here:** [`AGENTS.md`](AGENTS.md) for coding-agent and developer guardrails; [`HANDOFF.md`](HANDOFF.md) for the current implementation, limitations and next steps.
- **Understand the system:** [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md) for API/signature format and native controls.
- **Contribute:** [`CONTRIBUTING.md`](CONTRIBUTING.md) and [`SECURITY.md`](SECURITY.md).
- **Publish patches:** [`docs/RELEASING.md`](docs/RELEASING.md). Report bugs in GitHub Issues with sanitized logs; never share the pairing key.

Bug fixes are published as **new versioned releases**, preserving previous tags and binaries. Windows is supported today; Linux and a visual automation builder are potential future contributions, not current features.
