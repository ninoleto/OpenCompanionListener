# v0.1.4 release verification

Status: **owner live Windows PASS; final release authorized**. The owner confirms the EXE, tray, dashboard and existing Companion buttons work. Public release/download verification must pass before deleting the v0.1.3 GitHub release and tag. No rebuild is authorized.

Scope: finish module packaging and documentation, integrate the owner's supplied logo in the Windows EXE, tray and dashboard only, and correct version labels. Preserve dashboard layout/colors, listener behavior, protocol, action/preset IDs and migration scripts. Approved ICO/128px PNG supplied on 2026-11-12 and integrated unchanged; the SYSO was regenerated with the existing tool.

## Evidence

| Check | Result |
| --- | --- |
| Windows Go 1.27.4 race checks | PASS: 15 top-level tests and 20 key subtests, exit 0 |
| Go vet, CGO=1 and Windows/amd64 CGO=0 | PASS, exit 0 each |
| Windows GUI executable build | PASS, exit 0; PE x64/GUI and embedded API/dashboard version verified; approved artwork resource checks pass; final committed-source packages and approved artwork verified |
| Node 22.23.4 module suite | PASS: 28/28, exit 0 |
| Installed base 1.13.6 manifest validation | PASS, exit 0 |
| Full npm dependency audit | PASS: zero advisories, exit 0 |
| Official module build | PASS, exit 0; tools 2.9.1 bundled entrypoint, manifest, help and notices validated; runtime byte-identical to prior local candidate |
| GUI import and connection configuration | Owner confirms GUI TGZ import on production Companion host, retained settings and working connection (2026-11-12) |
| Live Windows EXE, tray, dashboard and existing Companion buttons | OWNER PASS (2026-11-12); detailed individual action/OS-matrix coverage unspecified |
| All 38 presets and existing saved buttons | Unit coverage passes; most existing buttons owner-tested; complete preset/edge-case coverage not claimed |
| Committed-source archives and checksums | PASS: two identical local builds, source/CRC/PE/module/hash checks |
| Independent public downloads | Pending publication and verification |
| GitHub Actions for release commit | Pending |

## Final local packages — 2026-11-12

Packaged source: `30be74c8841594d67a4aaf49ee5cc4f6db36c827`.
`dist/release-v0.1.4/` and `dist/release-v0.1.4-reproduction/` are byte-identical
across all five files: Windows ZIP, source ZIP, unchanged Companion TGZ,
BUILD-INFO.json and SHA256SUMS.txt. Both build runs exit 0. All four checksums,
ZIP CRCs, exact git-archive source contents/hashes, safe paths, notices and module
manifest/runtime hashes pass. The module was copied, not rebuilt or reinstalled.

The final extracted EXE is at
`dist/release-v0.1.4-extracted/OpenCompanionListener.exe`, SHA-256
`b8f74decb94e9e60ec6b4015cb1df7a95e2cdd806add3f72de0ff5dd0d6847b5`.
All ten embedded icon frames and group 1 match the supplied approved ICO.
Windows loads resource 1 at 16/32/48/256px without executing listener code.
The exact approved 128px PNG is embedded and renders in the unchanged 44px
dashboard logo slot; the mocked-data screenshot was visually inspected.
No dashboard HTML/style or functional code changed during logo integration.

Fresh branded Go race tests: 15 top-level + 20 subtests PASS, exit 0.
CGO=1 and Windows/amd64 CGO=0 vet and GUI build: PASS, exit 0 each.
Three negative packaging-argument/hash checks reject invalid reuse requests.
The existing module files and original TGZ match their pre-integration fingerprints.
Prior 28/28 module tests remain applicable and were not repeated.

At the local packaging checkpoint, live smoke was pending; the owner now
confirms live Windows EXE/tray/dashboard/existing-button PASS. No live listener/Companion settings or startup registry were changed.
No new antivirus result, remote CI, public download or independent installation
is claimed. The owner now explicitly authorizes publication and conditional old-release removal.
This post-build evidence update is separate from the frozen packaged source.

## Continuation checks — 2026-11-12

- Reused the downloaded Go/Node/compiler and existing npm dependencies; no tool downloads, new Companion environment or repeated security audit.
- `go test -race -count=1 -v ./...`: PASS, 15 top-level tests plus 20 key subtests, exit 0. `go vet ./...` with CGO=1 and Windows/amd64 CGO=0: PASS, exit 0 each. GUI build: PASS, exit 0.
- Installed-base manifest validation: PASS. `npm test`: 28/28 PASS, exit 0. Official package build: PASS, exit 0. The upstream base notice is supplied separately by the existing build configuration; the upstream builder's missing-license warning does not indicate a missing packaged notice.
- Package inspection and normalized payload equality: PASS. Five malformed-package probes (missing entrypoint, wrong runtime entrypoint, missing license, traversal path and external npm dependency) all rejected. Normalizing twice is byte-identical.
- `main.js` SHA-256 remains `e4ae10c1018d1f5bd2cda8432c1fbb72974453b353ff2dac8b37b6ed9d55ac06`. Only `pkg/companion/HELP.md` differs from the previous local TGZ. This comparison does not establish the exact file installed on the owner's host.
- Verified module: `dist/module-v0.1.4-verified/community-openlistener-0.1.4.tgz`, SHA-256 `949ab82e84ebf9cf0ccdaff3d811bc4486af41adfcf0a20c3f823928a299b725`; adjacent `SHA256SUMS.txt` covers that TGZ only. This is not the final five-file release set.
- Dashboard comparison against HEAD permits only `0.1.3` → `0.1.4` text: PASS. Module `src/` unchanged. All 34 relative documentation links resolve. Draft GitHub download links name future v0.1.4 assets and cannot be verified live before publication.
- Initial Node test workers were denied by the sandbox (`spawn EPERM`); the authorized local rerun passed. Initial Go telemetry access was redirected to a workspace-only disabled telemetry directory. PowerShell treated the builder's known license warning as a terminating error; a subprocess runner captured it and confirmed exit 0. An initial binary-check assertion normalized HTML line endings; checking the actual embedded bytes passed. These were tooling/probe issues, not application fixes.
- No executable smoke or visible tray check: the existing listener is running and was not stopped. No settings, startup entry or live Companion installation was changed. No new antivirus result is claimed for the rebuilt EXE.

The first Windows build attempt returned exit 0 despite a blocked PowerShell `npx` subprocess. Its archive lacked `main.js`. Process-scoped execution-policy configuration resolves that local build failure; archive validation now explicitly rejects missing entrypoints. Tools 3.2.0 requires base 2.x and was rejected without changing the runtime dependency. Tools 2.9.1 supports the pinned API 1.13 runtime.

## Required before publishing

Do not repeat the owner's completed GUI import and working-button tests. The report concerns the production Companion host, not the isolated test profile. Exact installed TGZ/EXE hashes, host versions, all-preset coverage and rollback were not supplied; do not infer them. Compare the final module runtime payload with the existing candidate, and record final package hashes after committed-source builds. Approved artwork passes PE resource equality, native resource loading at 16/32/48/256px and mocked dashboard rendering in the unchanged layout. The owner now confirms live tray and dashboard PASS. Remaining applicable cases in [SMOKE_TEST.md](../SMOKE_TEST.md) stay unverified until observed.

The current working listener and its settings must remain untouched. No private keys, account data, development snapshots or test profiles belong in public archives.
