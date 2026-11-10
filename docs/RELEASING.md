# Release workflow (maintainer and contributors)

This project uses `vMAJOR.MINOR.PATCH` Git tags. Use **new tags for fixes**; do not overwrite a published release asset or retag an existing version. For example, after v0.1.3 issue fixes become v0.1.4, not a silent replacement of v0.1.3.

## Prepare

1. Start from reviewed source; tag only a tested commit. Preserve a known-good release and backup any real Companion module/config before upgrades.
2. Update the version in the Go API status, dashboard, module `package.json`, Companion `manifest.json`, README, CHANGELOG and other versioned docs.
3. Inspect for secrets: config exports, private IPs/hostnames that identify deployments, actual pairing keys, `.env`, logs, passwords, caches, personal files and local path references. Do not publish these.
4. Verify the complete source, licensing and icon sources are included. The public repository must not contain user settings or private release archives.

## Automated checks

```
( cd listener && go test -race ./... && go vet ./... )
( cd listener && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go vet ./... )
( cd companion-dev-modules/companion-module-community-openlistener && node --test src/*.test.cjs )
( cd listener && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags='-s -w -H windowsgui' -o ../OpenCompanionListener.exe . )
```

Run Windows Defender or another trusted antivirus scan on the built asset if available. Calculate SHA-256 for release archives. A reproducible build process and Windows code-signing are future improvements, **not provided** by v0.1.3.

## Windows + Companion smoke test

Follow `SMOKE_TEST.md` with the candidate EXE and updated module installed on separate test machines if possible. Confirm the native tray, quiet startup, Windows autostart, keyboard/mouse including Num Lock, URL/app/file launching, all new builders/presets, and backwards compatibility for existing saved actions. No Windows desktop or real Companion upgrade claim may be based only on Go/Node unit tests.

## Publishing

1. Push full source, tests and documentation to the public GitHub repository.
2. Create an annotated release tag (`v0.1.3` for initial publication) pointing to the reviewed source commit.
3. Attach `OpenCompanionListener-v0.1.3-Windows-x64.zip` and `OpenCompanionListener-v0.1.3-Companion-Module.zip` as binary/download assets and their published checksums. GitHub itself provides source-code archive downloads.
4. Include concise installation instructions for Windows and the **Companion developer/extra module path**, with rollback notes and the real status of Windows manual verification.
5. Enable Issues and (if possible) GitHub private vulnerability reports. Use issue templates and `CONTRIBUTING.md`.
6. On bugs, publish a new patch version and changelog. Never amend the prior release tag or silently swap an existing archive.

## Public release text

Use `GITHUB_RELEASE.txt` as a starting point; avoid saying 'fully tested' until the Windows smoke test actually passes. If a feature is known to be unverified, state it explicitly.
