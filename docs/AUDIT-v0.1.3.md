# v0.1.3 pre-publication audit — 2026-11-11

## Final-validation preparation (latest session)

Candidate source remains `58b9f14821326cae0f6d92a9f1a379babff14c69` on
`fix/v0.1.3-publication`. No application, module, built archive, working
installation, live settings or paired device was changed in this follow-up.
Publication still requires the owner's explicit approval.

- Read-only Windows inspection identified the installed `OpenCompanionListener.exe`
  as the owner of port 17993 connections. The first permitted CIM query found no
  LISTENING entry, while netstat subsequently showed established/closing sockets
  owned by the still-running process. The owner received its PID/path and tray
  Exit instructions. No automatic termination or live-dashboard request occurred.
- Existing candidate checksums and the selected smoke EXE were reverified against
  the archive and BUILD-INFO.json: PASS. The source commit and executable hash
  are unchanged from the verified package below.
- The new [ZORIN runbook](ZORIN-VALIDATION.md) prepares backup, staged `npm ci`,
  real-base validation, isolated import/migration tests, future activation using
  a separate module parent and configuration-aware rollback. Its Bash examples
  pass syntax checking (exit 0); they have **not** been executed on ZORIN.
- [SMOKE_TEST.md](../SMOKE_TEST.md) now includes expected tray/startup, Num Lock,
  builder/hold/release, mouse/scroll, existing-button and all-preset observations.
  Manual startup tests explicitly require an isolated account/VM because changing
  LOCALAPPDATA does not isolate the HKCU startup value.
- The executable smoke test is **pending user tray exit** in this session.
  No candidate start/authentication/exit result is claimed yet. Existing Go/Node
  results below were not rerun for these documentation-only changes.

**Decision: HOLD.** Next: confirm the owner's tray exit, recheck process/port,
run the existing isolated executable smoke test, and record only its observed
results. Full manual/host checks and the other release gates remain open.

## Publication-fix continuation (current status)

The initial audit below is retained as historical defect evidence. Its statements
that no implementation changed apply to that earlier session only. Corrections
now exist on `fix/v0.1.3-publication`, continuing the prior session's uncommitted
work. Nothing has been published or installed into a live Companion instance.

| Finding | Current disposition |
| --- | --- |
| A01 | Fixed: exact dashboard authority/Origin and state-read header guards; regression coverage. |
| A02 | Compiler pinned to Go 1.27.4 in CI/tooling. Source and rebuilt binary vulnerability scans report no vulnerabilities; x64 GUI build verified. Original EXE remains untouched. |
| A03 | Fixed in source: preserve expression wrappers, retain migration index 0, additive index 1 repairs stranded scroll actions. Real base 1.13.6 upgrade processor tests pass. Actual Companion import/export remains a gate; already-lost expressions require the backup export. |
| A04 | Fixed for settings saves and in-process restart. Full process exit still clears the memory-only cache; limitation documented. |
| A05 | Fixed: serialize updates, bind/persist before activation, clean temporary files; failed rotation, network save, endpoint bind/rename, and concurrent saves covered. |
| A06 | Fixed: aborted/error/oversized responses reject once, 3.5-second overall deadline bounds trickling peers. Real loopback HTTP tests pass. |
| A07 | Corrected authority claims in README, SECURITY, architecture and contributor/agent guidance; existing launching behavior preserved. Success text now says OS accepted. |
| A08 | Go binary redistribution notice included and verified against toolchain. Module ZIP does not bundle dependencies. Artwork authorship/permission remains unresolved and blocks publication. |
| A09 | Base 1.13.6 pinned with lockfile; Node 22.x aligned; clean install and installed-library validation added to CI. Companion 4.1 is the upstream API minimum, not completed host certification. |
| A10 | Three candidate archives/checksums built twice byte-identically; contents/CRCs/source hashes/Go notice verified. Extracted module clean install/validation/28 tests pass. Guarded install/rollback docs added. Manual desktop/Companion, antivirus and owner publication gates remain. |

Fresh verification: Windows/amd64 Go 1.27.4 race tests pass
15 top-level tests plus 20 key subtests; both vet modes pass. Node 22.23.4 clean
install (17 dependencies), real base manifest validation and 28/28 Node tests
pass, including a clean install and rerun from the extracted module ZIP. npm audit reports zero advisories. govulncheck v1.9.0 source and binary scans report no
vulnerabilities (database updated 2026-11-05). See HANDOFF.md for the current
session and the release instructions for exact commands. Initial Node sandbox
workers failed with EPERM before running; the authorized rerun passed. Python's
private temporary build directory was likewise inaccessible in the sandbox;
the authorized packaging rerun succeeded with two byte-identical build sets.

The EXE is 6,990,848 bytes, SHA-256
`510696b6fba8331bfef6d929013213fea15f43f3dde78f7db34c8370dac97f83`,
Go 1.27.4, Windows/amd64 GUI, CGO=0, GOAMD64=v1. Checkpoint archive build and
content verification used commit `a9b6ae3`; final candidates include the later
evidence/tooling commit identified in their BUILD-INFO.json and are rebuilt with
the same process. Final archives/checksums live under ignored
`dist/publication-candidate/`, with a separate reproduction directory.

The limited binary smoke tool stopped before launching any process because
dashboard port 17993 was occupied (`EADDRINUSE`, exit 1). No live dashboard was
read or changed. Defender returned exit 0 but explicitly **skipped the file**;
this is not a clean scan. The unchanged artwork files, original snapshots, user
configuration, registry startup and installed Companion module were preserved.

**Decision remains HOLD.** Candidate packaging does not establish artwork rights,
actual Windows/Companion behavior, an antivirus result, owner repository links or
whether v0.1.3 has already been distributed. Local candidate filenames preserve
the original snapshots and avoid implying a replacement release.

## Original audit (historical)

**Decision: HOLD publication.** The supplied source builds and its existing tests pass, but the audit reproduced security and compatibility defects and found release-material gaps. No application or Companion implementation was changed. Nothing was published, tagged, uploaded, or installed into a live Companion instance.

## Scope and source baseline

The incoming workspace was a collection of release snapshots, archives, and research, not a Git repository. The public-source candidate is the 49-file `OpenCompanionListener-v0.1.3-GitHub-Source.zip`, SHA-256 `a9c00f97ca006023b938b074cbbf8dfa2d154b4d781d7ffec313f4acbedab3bf`. It was extracted into a separate audit workspace and recorded on branch `audit/v0.1.3`, baseline commit `cd2f0430d968794007f89facb64cb62f6f08c855`. This is a local provenance checkpoint, not a historical upstream commit or release tag.

Read the complete current Go and CommonJS source, tests, embedded dashboard, icon-resource generator, module metadata, release documentation, contributor/agent guidance, and CI/issue templates before making audit additions. Compared the public-source candidate with the extracted v0.1.3 package: implementation and assets match; the source candidate adds public documentation/CI and replaces private-range authentication-test IPs with documentation-range addresses. Historical snapshots were inspected for privacy and compatibility context, not certified as releases.

The audit checkout is the intended starting point for corrections. Do not initialize/publish the entire surrounding workspace: it contains historical deployment information, unrelated research, old binaries, downloaded toolchains, and diagnostic logs.

## Confirmed findings and required corrections

### A01 — High: dashboard discloses the pairing key for an arbitrary Host

`listener/main.go:175` returns all of `Config`, including `Secret`, from `GET /ui/state`. The handler never validates `r.Host`; the Origin check at line 196 applies only after GET requests have already returned.

An isolated request addressed to `http://rebinding.example:17993/ui/state`, with a loopback peer, received HTTP 200 and the generated test pairing key. The test does not print the key. This is a server-side DNS-rebinding weakness: a browser able to reach loopback through an attacker-controlled hostname could expose local state. A complete browser attack was not performed; browser local-network protections and trusted API peer restrictions affect exploitability. The UI is still bound to loopback, not directly to the LAN.

**Correction:** validate the allowed local authority before serving any dashboard route, including GET state, while retaining mutation Origin/header protections. Add positive tests for the supported dashboard URL and negative tests for foreign/malformed authorities. This does not require changing the wire protocol or architecture.

### A02 — High: supplied Windows EXE embeds an obsolete Go standard library

`go version -m` identifies the supplied executable as **Go 1.23.2**, with `CGO_ENABLED=0`, Windows/amd64, and trimpath enabled. `govulncheck v1.9.0 -mode=binary` reported **67 standard-library advisories**, exit 3. Binary scanning does not establish reachability or exploitability of all 67 findings; some concern facilities this application does not use. The result is still a reason to rebuild before release.

The unchanged source rebuilt using **Go 1.27.4** passed source and rebuilt-binary vulnerability scans, both reporting no vulnerabilities, exit 0. Go 1.23 is outside the current support window under the [Go release policy](https://go.dev/doc/devel/release). CI currently derives its compiler from the `go 1.23` language directive and does not select a supported release toolchain separately.

**Correction:** select and record a supported patched release compiler in CI/release instructions, rebuild and smoke-test the candidate, and generate new artifact checksums. A newer compiler changes the embedded runtime and does not by itself prove desktop behavior. Never silently replace an already published asset or tag; establish whether previous distribution constrains version numbering.

### A03 — High: legacy Companion actions can become unusable during migration

`src/upgrades.cjs:51` leaves unconvertible legacy `scroll` values unchanged, but `src/main.cjs` no longer registers `scroll`. A wrapped variable expression therefore retains an action ID with no executable definition. `sameWrap` also forces `isExpression:false`, and migration parses expression contents without first checking the expression flag. Numeric expression wrappers can become literals in scroll and shortcut options.

Reproduced through the **installed base library's actual upgrade processor**, not just the repository's mock: 3 targeted assertions fail, while ordinary signed horizontal scrolling and empty-shortcut handling pass. Dynamic legacy scroll is the demonstrated functional break; literal-looking expression cases demonstrate lost expression metadata, not necessarily a changed keystroke for the same numeric value.

**Correction:** ensure every retained legacy action still has an execution path and preserve expression semantics/wrappers. Add an additive repair migration if the existing script has been distributed; retain its index. Test raw/wrapped values, expressions, empty options, unusual shortcuts, and round trips through a real Companion export/import. Keep one primary user-facing builder without sacrificing saved-button compatibility.

### A04 — Medium: saving settings re-enables recently consumed requests

`listener/main.go:231` replaces the replay cache even when saving identical settings. `start` also creates a new cache during in-process restart. The audit accepted a signed status request, saved identical configuration, then replayed the same request and received HTTP 200 instead of 401.

**Correction:** preserve the replay state across settings saves and in-process restarts when the pairing key remains valid. Test this lifecycle explicitly. This is a bounded replay window, not a way to forge a new signature; the attacker still needs captured traffic and an accepted TCP peer. Cache loss on full process restart is a separate current limitation that needs an explicit policy.

### A05 — Medium: a failed settings write can still change live authentication

`listener/main.go:267` changes the active pairing key before `saveConfig` succeeds. A deliberately unwritable destination produced HTTP 500 while the in-memory key had already changed. The old persisted key will return at the next process start. Network settings are likewise installed before persistence at lines 223–234.

**Correction:** keep active/persisted settings consistent on failure and serialize conflicting settings operations. Add regression cases for disk/rename failure during rotation and network updates. Preserve the existing key on failed rotation; do not log either key.

### A06 — Medium: truncated HTTP responses leave module requests pending

`src/transport.cjs:22` handles response `data` and `end`, but not an aborted/erroring response. A local test server sent an incomplete JSON response then closed the socket. The request promise was still pending after 4.5 seconds despite the 3.5-second request timeout.

**Correction:** reject interrupted responses exactly once, handle response errors/aborts, and terminate oversized responses instead of continuing to consume their data. Add real local HTTP tests for success, non-2xx, malformed JSON, body limits, timeout, and truncated response. Preserve HMAC signing and command payloads.

### A07 — Security documentation overstates the restriction on shell execution

The README says arbitrary remote shell execution is unsupported. There is no dedicated `shell` command type, but `open_file` accepts arbitrary executable paths and arguments (`listener/commands.go:63`) and passes them to `ShellExecuteW`. Validation accepted interpreter executables and command arguments in a non-executing probe. Registered URI handlers and keyboard control also carry substantial authority.

This requires an authenticated trusted peer; it is not an unauthenticated execution bypass. `SECURITY.md` already usefully describes pairing as desktop access. [Microsoft's ShellExecuteW documentation](https://learn.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew) confirms the executable/parameter behavior.

**Correction:** reconcile README, SECURITY, architecture, and agent guidance with the actual full-user-session authority. Preserve working application launching during this audit. If a genuinely restricted launcher is desired, that is a separate compatibility/security design; an interpreter-name denylist would not provide that boundary.

### A08 — Release licensing/provenance is incomplete

Both project LICENSE files contain MIT terms and contributor attribution. The listener has no external Go modules, but the statically linked Go runtime/standard library is still third-party code. The Windows ZIP contains only the project's MIT license, without the Go notice required for binary redistribution by the [Go license](https://go.dev/LICENSE).

The 17 installed Node package roots declare 15 MIT, one BSD-3-Clause, and one 0BSD license. No incompatible declared license was identified. The installed `@companion-module/base` tarball declares MIT but contains no top-level license notice; retain the upstream notice if bundling dependencies. This is an inventory, not proof of authorship of every source line.

The ICO, PNG, and generated resource are included. Regenerating the resource with the included Python tool produced an identical SHA-256. Artwork origin and redistribution rights are not recorded; a maintainer clarification was requested and remains outstanding.

**Correction:** include the appropriate Go and other applicable notices with binary distribution, preserve dependency notices when bundling them, and document artwork provenance/license. Do not invent an author or claim provenance was verified.

### A09 — Build/runtime support and dependency resolution are underspecified

The module declares Node `>=20`, its manifest selects `node22`/API `1.13.0`, and the floating dependency `^1.13.0` installed base **1.15.1**. That version advertises Node `^18.12 || ^22.8 || ^26.5`; installation under the preinstalled Node 24.18.0 emitted an engine warning. Node 22.23.4 tests and real-library manifest validation passed. This is dependency drift and inconsistent support guidance, not proof that the module fails to load on every host.

The source has no lockfile. CI runs mocked Node unit tests without installing the runtime dependency. It cannot detect installation or real API compatibility failures. The public documentation also lacks a tested minimum Companion version. The [official compatibility matrix](https://bitfocus.github.io/companion-module-base/) ties API support to Companion versions; determine support using the chosen dependency version and actual host tests.

**Correction:** choose a tested API 1.x dependency set, preserve its resolved lockfile, align Node and Companion requirements, and add clean dependency installation/manifest validation to CI. The audit-generated lockfile is evidence of what resolved, not an approved dependency update.

### A10 — Packaging, public documentation, and release evidence need reconciliation

- The public checksum list names `OpenCompanionListener-v0.1.3-Companion-Module.zip`, but that file is absent. The source and Windows ZIP checksums and all three inner snapshot checksums match.
- The GitHub source ZIP has the complete public layout. The workspace root and extracted Windows snapshot do not have the same guidance/security/architecture layout. Root `docs/...` links fail if that directory is treated as the repository.
- Historical v0.1.1 material contains a personal home-path example; historical documentation and the research note contain deployment IPs. Those files are excluded from the GitHub source candidate. The snapshot/Windows ZIP still uses private-range test IPs, while the GitHub candidate uses RFC documentation addresses. Private-range test data alone is not proof of leaked deployment credentials.
- README's Linux module-copy example assumes an existing destination, does not stop if backup/copy/cd fails, and mixes first installation with upgrade. Separate and guard those procedures before recommending them to contributors.
- Public development commands use Bash syntax even though Windows is the primary target. Add PowerShell equivalents, Go/Node prerequisites, and Windows race-detector C compiler requirements. Python is optional for icon regeneration, not required to run the product. See [Go race-detector requirements](https://go.dev/doc/articles/race_detector).
- `HANDOFF.md` describes a released baseline and a compound next action including upload. It needs the current HOLD status and test evidence. `AGENTS.md` otherwise provides useful contracts, source mapping, migration rules, rollback guidance, and release constraints; retain them.
- `SMOKE_TEST.md` is still unchecked. No real Companion migration, Windows tray/autostart sign-in, input injection, launching, focus/UIPI, Explorer restart, or interrupted-hold behavior was validated in this audit. The event text `Executed ...` should distinguish OS acceptance from a completed downstream GUI task.
- The distribution and rebuilt executable are unsigned as documented. A Defender custom scan with remediation disabled failed with `0x80508023`, exit 2; it is **not** a clean antivirus result.

**Correction:** publish only the reviewed repository layout; finish documentation and packaging after code fixes, produce all advertised archives/checksums, and record actual manual verification. Populate repository/bugs/security-reporting links only when the real owner URL exists.

## Exact automated evidence

Test host: Windows amd64, reported OS build `10.0.26300`; PowerShell 7.6.9. Portable tools were downloaded into the audit workspace and checksum-verified against their publishers: Go 1.27.4, Node 22.23.4, LLVM-MinGW 20261104 (Clang 23.1.5). No new application dependencies were introduced.

| Check | Result |
| --- | --- |
| `go test -race -v ./...` | PASS, exit 0: 7 top-level tests, 20 key subtests; no reported races in exercised code |
| `go vet ./...` with Windows/amd64, CGO=1 | PASS, exit 0 |
| `go vet ./...` with Windows/amd64, CGO=0 | PASS, exit 0 |
| `go build -trimpath -ldflags='-s -w -H windowsgui' -o <audit-artifact> .` with Windows/amd64, CGO=0 | PASS, exit 0 |
| `node --test src/*.test.cjs`, Node 22.23.4 | PASS, 16/16, exit 0 |
| Same original Node tests, Node 24.18.0 | PASS, 16/16; dependency engine warning means this is not runtime support certification |
| `npm install --omit=dev --ignore-scripts --no-audit --no-fund` | PASS: 17 dependency packages installed; generated lockfile records resolution |
| Installed base `validateManifest(manifest, false)` | PASS against base 1.15.1 |
| `npm pack --dry-run --ignore-scripts --json` | PASS: 14 source/metadata/license files, no bundled dependencies; nothing published |
| `npm audit --omit=dev --json` | PASS, exit 0: zero reported advisories in resolved tree |
| `govulncheck -mode=binary <supplied-exe>` | FAIL, exit 3: 67 standard-library advisories; reachability not established |
| `govulncheck ./...` and `govulncheck -mode=binary <rebuilt-exe>` | PASS, exit 0 each: no reported vulnerabilities |
| Go audit overlay, `go test -race -overlay <overlay.json> -run TestAudit -v ./...` | 3 failed, 3 passed; HTTP guard subtests and Win64 ABI layouts pass |
| `node --test <migration-audit.test.cjs>` | 3 failed, 2 passed against real base upgrade processor |
| `node --test <transport-audit.test.cjs>` | 1 failed: aborted response stays pending |
| `py -3.13 tools/generate_icon_resource.py` | PASS: output byte-identical to checked-in resource |
| Archive/executable hashes | Five existing listed objects match; one advertised module ZIP missing |
| Read-only privacy scan | No matches in public source ZIP; historical workspace findings listed above |
| Defender custom scan, remediation disabled | FAILED to scan: exit 2, `0x80508023` |

The seven failed added assertions are **audit reproductions of existing defects**, not newly introduced regressions. Initial sandboxed Node workers failed with EPERM before running tests; rerunning with process execution allowed passed all supplied tests. A first version of the migration probe attempted to assign to a getter; that audit harness error was fixed before the recorded 3-fail/2-pass result.

The audit binary is 6,980,096 bytes, Windows x64 GUI subsystem, SHA-256 `6be113a2214c3d0bdf7e7c91e8914ca2b5e6ea3b4989686aefbf3a4c45d2dcc3`. It is an **unapproved test artifact**, not a replacement release. Its build metadata points at the local baseline with `vcs.modified=true` because dependency installation created an untracked lockfile. The supplied 5,659,648-byte EXE remains untouched. Byte-for-byte equivalence to that older-compiler EXE was not claimed.

## Reproduction, changes, and next action

Local-only audit tools/evidence live in the parent audit directory: `run-checks.ps1`, `inspect_artifacts.py`, `probes/`, `logs/`, `go-overlay.json`, `toolchains/`, and `artifacts/`. The runner sets process-local Go/compiler/Node paths and writes only audit artifacts; overlay probes do not change listener source. The transport probe is run separately with Node 22. Full logs and the generated dependency lockfile must be reviewed before any future public inclusion.

Repository changes from this session are the audit report and HANDOFF status only; a local dependency-install lockfile is untracked evidence. Root workspace HANDOFF also receives a pointer so the next contributor finds the correct checkout. No protocol, action IDs, migrations, UI, registry settings, live configuration, or application dependencies were changed. Rollback is to continue using the untouched supplied snapshots; the local audit branch can simply be left unused. There are no published changes to roll back.

**Single next safe action:** after reviewing these findings, implement and regression-test the narrowly scoped dashboard authority validation fix in A01 on the audit branch. Keep publication blocked until the remaining findings and manual release gates are resolved.
