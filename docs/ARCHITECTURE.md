# Architecture and protocol (v0.1.3)

## Components and boundaries

```
Stream Deck / Loupedeck / any Companion surface
                |
      Bitfocus Companion (module)
       Node.js 20+ / CommonJS
                |
   HTTP + HMAC-SHA256 over trusted LAN
           default TCP 17992
                |
       Native Go Windows listener
       -- validation and replay guard
       -- Windows SendInput and Shell APIs
       -- app/session lifecycle + tray
                |
         Active Windows desktop

Local browser (on Windows PC only) -> 127.0.0.1:17993 -> settings/status UI
```

The Companion module is independent of any other listener/module. It can run on a different machine from Windows. No cloud account, Python interpreter, or arbitrary shell action is needed.

## Protocol

Requests:
- `GET /v1/status` (empty body), returns JSON with `ok`, `version`, `platform` when paired.
- `POST /v1/command` (UTF-8 JSON), returns `{"ok":true}` when the listener's action function returns no error. **This does not guarantee a spawned application fully initialized.**
- Both endpoints require the same signed authentication headers and a matching trusted TCP peer IP.

Headers:
- `X-OCL-Timestamp`: Unix epoch seconds (decimal), accepted within approximately ±60 seconds.
- `X-OCL-Nonce`: 16 cryptographically random bytes, lowercase hexadecimal (32 characters), not previously seen in the replay cache.
- `X-OCL-Signature`: lowercase hexadecimal HMAC-SHA256 signature (64 characters).

Signature input: construct a single UTF-8 string with each value separated by a newline (`\n`), **no trailing newline**:

```
METHOD
URL_PATH
TIMESTAMP
NONCE
HEX_SHA256(raw_request_body_bytes)
```

Compute `hex(HMAC_SHA256(pairing_key_bytes, UTF8(payload)))`, where the pairing key is exactly 32 raw bytes represented as 64 hexadecimal characters in settings. `GET` hashes an empty byte body. The server checks the source IP, method, request size (maximum 8192 bytes), signature in constant time, timestamp validity, and nonce uniqueness. Replay cache items are retained for a short window.

This is **HTTP without TLS**. The signature is not encryption. Only deploy on a trusted LAN or private overlay with firewall access restricted to the Companion host.

## Commands

All commands are JSON with a `type` field. Only listed fields are accepted (unrecognized fields are rejected). Example shapes:

```json
{"type":"hotkey","keys":"CTRL+SHIFT+F23"}
{"type":"key_down","keys":"CTRL+F13"}
{"type":"key_up","keys":"CTRL+F13"}
{"type":"num_lock","mode":"off"}
{"type":"click","button":"left"}
{"type":"mouse_down","button":"left"}
{"type":"mouse_up","button":"left"}
{"type":"move","dx":10,"dy":-10}
{"type":"scroll","axis":"vertical","steps":-1}
{"type":"open_url","url":"https://example.com/"}
{"type":"open_file","path":"C:\\Windows\\notepad.exe","args":""}
```

- Wheel values: integer `-20..-1` or `1..20`; `0` is rejected. One step maps to 120 Windows wheel units.
- Move values: each signed integer `-10000..10000`.
- `num_lock` modes: `toggle`, `on`, `off`.
- Mouse button: `left`, `right`, `middle`.
- URL: absolute `http`/`https` only, maximum 2048 characters.
- File path: nonempty, maximum 1024 bytes, no CR/LF/NUL; arguments maximum 1024 bytes.
- Key strings: `+`-separated supported key identifiers from `listener/keys.go`, max length 100. See module `src/keys.cjs` for dropdown choices.

## Windows UI, settings, and process lifetime

- UI is served only on loopback `127.0.0.1:17993`; it is not the remote-control API.
- `OpenCompanionListener.exe` starts quietly and places a native icon in the tray; browser dashboard opens only when requested.
- Tray: Open GUI, Restart Listener, Exit. Restart rebinding affects API port while leaving local UI/tray alive.
- Per-user HKCU Run startup is opt-in; launching at Windows sign-in does not create a Windows service.
- Settings live at `%LOCALAPPDATA%\OpenCompanionListener\settings.json`. This contains a private pairing key and **must never be committed**.
- The local UI includes signed-key rotation; rotating requires updating the Companion connection.

## Companion module design

- `src/main.cjs`: registration of actions, settings, and connection-status checks.
- `src/transport.cjs`: HMAC signing and HTTP connection pooling.
- `src/keys.cjs`, `src/windows-tools.cjs`: dropdown choices.
- `src/presets.cjs`: 38 optional presets, never automatically mapped to the user's grid.
- `src/upgrades.cjs`: action migration from previous module versions. Must remain compatible with old stored config until explicitly deprecated and documented.
- `@companion-module/base` is the only external runtime module dependency.

## Future Linux (not implemented)

Share transport/configuration and command semantics, then build a separate OS adapter. X11 XTEST and Wayland RemoteDesktop/libei require different permission/session handling. Linux window focus/input may not match Windows behavior. Do not implement root-owned `/dev/uinput` injection by default. Update documentation and test matrices before advertising Linux support.
