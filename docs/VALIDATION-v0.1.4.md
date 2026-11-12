# v0.1.4 release verification

Status: **not published; validation in progress**. Do not publish until final packages are verified. Preserve v0.1.3 unchanged.

Scope: finish module packaging and documentation, integrate the owner's supplied logo in the Windows EXE, tray and dashboard only, and correct version labels. Preserve dashboard layout/colors, listener behavior, protocol, action/preset IDs and migration scripts. Approved ICO/128px PNG supplied on 2026-11-12 and integrated unchanged; the SYSO was regenerated with the existing tool.

## Evidence

| Check | Result |
| --- | --- |
| Windows Go 1.27.4 race checks | PASS: 15 top-level tests and 20 key subtests, exit 0 |
| Go vet, CGO=1 and Windows/amd64 CGO=0 | PASS, exit 0 each |
| Windows GUI executable build | PASS, exit 0; PE x64/GUI and embedded API/dashboard version verified; approved artwork resource checks pass; committed-source package verification follows the source freeze |
| Node 22.23.4 module suite | PASS: 28/28, exit 0 |
| Installed base 1.13.6 manifest validation | PASS, exit 0 |
| Full npm dependency audit | PASS: zero advisories, exit 0 |
| Official module build | PASS, exit 0; tools 2.9.1 bundled entrypoint, manifest, help and notices validated; runtime byte-identical to prior local candidate |
| GUI import and connection configuration | Owner confirms GUI TGZ import on production Companion host, retained settings and working connection (2026-11-12) |
| Real Windows input and launch checks through installed package | Owner reports most existing buttons working; individual action coverage unspecified |
| All 38 presets and existing saved buttons | Unit coverage passes; most existing buttons owner-tested; complete preset/edge-case coverage not claimed |
| Committed-source archives, checksums and independent downloads | Pending |
| GitHub Actions for release commit | Pending |

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

Do not repeat the owner's completed GUI import and working-button tests. The report concerns the production Companion host, not the isolated test profile. Exact installed TGZ/EXE hashes, host versions, all-preset coverage and rollback were not supplied; do not infer them. Compare the final module runtime payload with the existing candidate, and record final package hashes after committed-source builds. Approved artwork passes PE resource equality, native resource loading at 16/32/48/256px and mocked dashboard rendering in the unchanged layout. Actual visible-tray testing remains pending while the production listener runs. Remaining applicable cases in [SMOKE_TEST.md](../SMOKE_TEST.md) stay unverified until observed.

The current working listener and its settings must remain untouched. No private keys, account data, development snapshots or test profiles belong in public archives.
