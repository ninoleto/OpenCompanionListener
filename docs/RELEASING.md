# Release workflow (maintainer and contributors)

Publication is on HOLD. The current work produces local review candidates only.
Do not push, tag, upload or replace published assets as part of this workflow.
Preserve the supplied v0.1.3 snapshots. If v0.1.3 has been publicly distributed,
release the fixes under a new patch version after the gates below pass.

## Toolchain and checks

Use the versions in `tools/release-toolchain.json`: Go 1.27.4, Node 22.23.4,
Companion base 1.13.6. The compiler is selected independently of the Go language
minimum in go.mod. Go's [download metadata](https://go.dev/dl/?mode=json)
listed 1.27.4 as stable at verification time. Recheck supported/security-patched
versions before eventual publication and deliberately update the pins if needed.

Windows race tests require CGO and a compatible MinGW C compiler (the local
verification used LLVM-MinGW 20261104, Clang 23.1.5). The GUI EXE requires no C
runtime toolchain at execution. Python 3.11+ is used only by the packaging tool;
local verification used Python 3.13. See README for PowerShell and Bash commands.

Required before packaging:

1. `go test -race -count=1 ./...` and `go vet ./...` in listener.
2. Windows/amd64 `CGO_ENABLED=0 go vet ./...` and GUI build.
3. In the module directory, `npm ci --omit=dev --ignore-scripts --no-audit --no-fund`,
   `npm run validate`, and `npm test`. Stop at any error. Migration regression
   tests use the installed base library, not just a mock.
4. `npm audit --omit=dev` and current `govulncheck ./...`; scan the built EXE with
   `govulncheck -mode=binary`. Record tool/database time and any findings, without
   interpreting zero advisories as proof of safety.
5. Review source for credentials, private config, deployment identifiers and local
   paths. Check `git diff --check`; commit the reviewed changes on the fix branch.

## Reproducible local candidates

The builder requires a clean working tree and reads only `git archive HEAD`.
It verifies the compiler version, uses Windows/amd64, GOAMD64=v1, CGO=0,
`-trimpath -buildvcs=false -ldflags='-s -w -H windowsgui'`, and records the source
commit separately in BUILD-INFO.json. It clears GOFLAGS/GOEXPERIMENT and disables
GOENV/GOWORK. It does not sign the EXE or include machine-specific build paths.

From the repository root, with Go on PATH (PowerShell Python launcher example):

```powershell
py -3.13 tools/package_release.py --output dist/candidate-1
if ($LASTEXITCODE) { throw 'First package build failed' }
py -3.13 tools/package_release.py --output dist/candidate-2
if ($LASTEXITCODE) { throw 'Second package build failed' }
if ((Get-Content dist/candidate-1/SHA256SUMS.txt -Raw) -cne (Get-Content dist/candidate-2/SHA256SUMS.txt -Raw)) {
    throw 'Candidate builds differ'
}
Get-Content dist/candidate-1/SHA256SUMS.txt | ForEach-Object {
    $expected, $name = $_ -split '  ', 2
    if ((Get-FileHash -Algorithm SHA256 (Join-Path dist/candidate-1 $name)).Hash.ToLowerInvariant() -ne $expected) {
        throw "Checksum mismatch: $name"
    }
}
```

Bash equivalent:

```bash
set -eu
python3 tools/package_release.py --output dist/candidate-1
python3 tools/package_release.py --output dist/candidate-2
cmp dist/candidate-1/SHA256SUMS.txt dist/candidate-2/SHA256SUMS.txt
(cd dist/candidate-1 && sha256sum --check SHA256SUMS.txt)
```

Choose fresh output directories on subsequent runs; existing candidates are
never overwritten. ZIP entries have sorted names, fixed timestamps/permissions
and no compression so zlib versions cannot change them. Text comes from Git
objects with canonical LF, independent of checkout CRLF conversion. Two local
builds prove repeatability on that toolchain; cross-host equivalence still needs
an independent builder. CI repeats the build on Linux but has not been run
remotely as part of local preparation.

Each output directory contains:

- `OpenCompanionListener-v0.1.3-candidate-Windows-x64.zip`: EXE, full source/docs,
  project license and Go redistribution notice. Complete layout keeps links valid.
- `OpenCompanionListener-v0.1.3-candidate-Companion-Module.zip`: named module folder,
  source/tests/manifest/license, lockfile and standalone INSTALL.md. No node_modules
  or Node.js bundle; run the locked install on the Companion host.
- `OpenCompanionListener-v0.1.3-candidate-Source.zip`: complete committed source.
- `BUILD-INFO.json`: commit, toolchain, flags, EXE hash and per-source-file hashes.
- `SHA256SUMS.txt`: SHA-256 for all three archives and BUILD-INFO.json.

Extract the module archive into a fresh staging directory and repeat its clean
install, manifest validation and tests. Check Windows archive CRCs, EXE metadata
and notices. Run Defender or another trusted antivirus scanner on the EXE if
available; record a scanner failure as unverified. Never modify live user settings
or a live Companion installation during this verification.

For a limited native-binary check on an isolated Windows account, run
`node tools/smoke_listener.cjs <candidate.exe>`. It refuses an occupied dashboard
port, gives the child a fresh LOCALAPPDATA fixture under .cache, and checks UI
guards, signed Node-to-Go status, settings/replay/key rotation, duplicate process
exit and shutdown. It injects no input, opens no browser/apps and writes no
startup registry settings. It does not replace the manual smoke checklist; the
local session could not run it because port 17993 was already occupied.

## Remaining release gates

Follow `SMOKE_TEST.md` on isolated Windows 10/11 and Companion hosts. Unit tests
and a PE build cannot prove tray, sign-in startup, focus/UIPI, key/mouse holds,
all presets, real-host import/export migrations or rollback behavior.

Record icon authorship and redistribution permission in THIRD_PARTY_NOTICES.md.
If unknown, keep publication blocked; do not infer a license grant. The package
includes Go's notice. Bundling Node dependencies later requires all corresponding
notices, including the upstream base license absent from its npm tarball.

Confirm previous distribution status and the actual owner/repository URL. Fill
manifest repository/bugs and security-reporting links with real destinations;
never invent them. If a patch version is required, update Go status, dashboard,
package.json, lockfile, manifest and docs together, then rerun validation/builds.

## Eventual publishing (separately authorized, not done here)

Only after the owner approves a concrete reviewed candidate and the gates above
are recorded: push the source, tag its tested commit with a new unused version,
upload the verified archives and checksums, and independently download/install
them. Do not claim a release shipped before that independent installation.
Keep older tags/assets intact. Use GITHUB_RELEASE.txt as a draft and describe
remaining limitations honestly. A release rollback restores both the old EXE
and the backed-up Companion module/configuration; do not overwrite snapshots.
