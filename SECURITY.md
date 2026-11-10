# Security policy and deployment cautions

## Supported versions

The maintainer intends to provide fixes for the newest public release. Older releases may receive no security updates. Check the GitHub release page for the currently maintained version.

## What this application can do

A paired Companion host can send keyboard and mouse input and request that the Windows user open applications, files, folders, or URLs. `open_file` accepts executables and arguments, including command interpreters; registered URI handlers and keyboard input also carry broad authority. Treat pairing as giving the trusted host code-execution authority in your active desktop user session. This is not a restricted launcher or an application allowlist. Do **not** run the listener as Administrator.

## Important limits

- **Requests are signed, not encrypted.** The transport is HTTP, not TLS. HMAC-SHA256 provides integrity/authentication but not confidentiality; traffic contents can be observed on an untrusted network.
- The listener checks the connecting TCP peer IP and signed request with timestamp/nonce replay protection. Source-IP restriction is defense in depth, not a replacement for firewall or encryption.
- The pairing key is stored in `%LOCALAPPDATA%\OpenCompanionListener\settings.json` as user-readable JSON and shown in the local loopback dashboard. Protect your Windows account and don't publish that file. Rotate the key if exposed and update the Companion connection.
- **Never port-forward API TCP 17992 or dashboard TCP 17993 to the Internet.** Only permit TCP 17992 from your specific Companion host via Windows Firewall; the dashboard is intended to bind to `127.0.0.1` only.
- The listener cannot and should not bypass Windows UAC, lock screen, integrity levels, or session isolation.
- There is no dedicated shell action, but the existing application launcher can invoke interpreters with arguments. Authentication and trust in the paired host are the security boundary; absence of a shell-labelled action is not a sandbox.
- The dashboard requires the exact authority `127.0.0.1:17993`; alternate hostnames such as `localhost` are intentionally rejected. State reads require the dashboard's custom header; foreign Origins and cross-site state requests are rejected.
- Replay entries survive settings saves and tray/API restarts within the same process. They are memory-only: after a full process exit, captured requests still inside the timestamp window can be replayed by an accepted TCP peer. Use a trusted network/private overlay; rotate a compromised key. A request already being authenticated when rotation occurs may complete using its captured settings.
- Settings updates are serialized and persisted before activation. A failed write or rebind preserves the previous live configuration. File replacement does not promise crash/power-loss durability on every Windows filesystem.

## Reporting vulnerabilities

Please use **GitHub's private security-advisory reporting** for this repository if enabled. Otherwise contact the maintainer privately through an established channel before filing public reproduction details. Do not include real credentials, pairing keys, or personal data in reports. There is no guaranteed response-time SLA.

When reporting: identify the exact version/commit, the affected API or action, a minimal reproduction using test credentials, impact, and any mitigations. Avoid using actual production hosts.
