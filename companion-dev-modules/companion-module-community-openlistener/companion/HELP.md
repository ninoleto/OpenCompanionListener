# Open Companion Listener 0.1.3

Free, open-source desktop control for Windows via Bitfocus Companion.

1. Install and run the Windows x64 listener on the controlled PC. It starts in the system tray without opening its dashboard.
2. Right-click tray icon → **Open GUI**; configure the Windows Listen IP and Trusted Companion server IP. API default port is `17992`.
3. Copy the pairing key into this Companion connection (treat the key as a secret).
4. Restrict Windows Firewall inbound TCP 17992 to the Companion host's IP. Never expose the port on the public Internet.
5. Use the action builders for keyboard shortcuts, mouse wheel scrolling and Windows Settings / Control Panel utilities.
6. Optional ready-made buttons are available in **Presets → Open Companion Listener**. Drag them to any surface button.

## Actions

Keyboard Press/Hold/Release shortcut actions use Ctrl/Shift/Alt/Windows checkboxes plus a key dropdown. Advanced shortcut text is supported for uncommon combinations. Single Key actions include F1–F24, numpad and navigation keys. Num Lock is a separate state action.

Mouse wheel: choose direction with Reverse, speed with a 1–20 slider, axis Vertical/Horizontal. File/URL open actions use native Windows Shell.

Existing v0.1.2 actions are upgraded using the module's included upgrade script. Export your Companion configuration before updating an existing installation.

## Status and security

The API requires HMAC-SHA256 authenticated commands, timestamps, nonces and a trusted source IP. Payloads are not encrypted. Use private networks only.
