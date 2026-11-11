# Release workflow (maintainer and contributors)

The owner authorizes the first public v0.1.3 release, including public repository
creation, source push, annotated tag and release assets. GitHub CLI authenticated
as `ninoleto` and created the empty public `ninoleto/OpenCompanionListener`
repository with Issues enabled; `origin` points there. Earlier 404-based stops
are historical. The release uses the corrected EXE with SHA-256
`510696b6fba8331bfef6d929013213fea15f43f3dde78f7db34c8370dac97f83`.
Preserve supplied snapshots and report incomplete manual coverage honestly.
Never replace a published version's tag or assets; later fixes need a patch release.

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
and no compression so zlib versions cannot change them. Text comes from
`git archive` and follows Git attributes and configuration. Verification found
CRLF conversion in six text files without an explicit `eol=lf` attribute,
including LICENSE and go.mod, under the local `core.autocrlf=true` configuration.
The archived bytes match the candidate manifest and differ from their Git blobs
only by those line endings. Do not assume every archive entry has canonical LF
or claim cross-host byte identity. Two local builds prove repeatability on that
toolchain/configuration; cross-host equivalence still needs
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
startup registry settings. It does not replace the manual smoke checklist.
After the owner exited the existing listener through its tray, the 2026-11-11
validation run passed with exit 0 and confirmed the original EXE/settings/startup
value were unchanged. See HANDOFF.md for exact evidence and unperformed checks.

## Verification scope

Use `SMOKE_TEST.md` for continuing verification on isolated Windows 10/11 and Companion hosts. Unit tests
and a PE build cannot prove tray, sign-in startup, focus/UIPI, key/mouse holds,
all presets, real-host import/export migrations or rollback behavior.
The owner's partial working-function and ZORIN installation reports are now
recorded there, separately from detailed procedures that remain unverified.
The running listener matches the corrected EXE; manual-test-time identity is
still awaiting confirmation. Do not claim complete Windows configuration coverage.

Artwork provenance and MIT inclusion are now recorded in THIRD_PARTY_NOTICES.md
from the owner's confirmation, with matching packaged and embedded bytes.
A completed Defender scan of the corrected EXE reported no threats on
2026-11-11; versions and limitations are recorded in the audit. The package
includes Go's notice. Bundling Node dependencies later requires all corresponding
notices, including the upstream base license absent from its npm tarball.

The owner identifies this as the first public release and explicitly approves
publication with the documented testing limits. The module manifest now points
to the actual repository and Issues URLs. No additional manual cases are marked
passed by this approval. Any later behavior changes require a new version.

## Final packaging and publishing

Commit reviewed publication documentation on `fix/v0.1.3-publication`. Build
twice into fresh directories with the command below, choosing a new output
directory for the second run:

```powershell
py -3.13 tools/package_release.py --release --output dist/release-1
```

This opt-in flag changes only archive names and BUILD-INFO publication labeling;
build flags and content selection are unchanged. Compare both SHA256SUMS files
and require the EXE hash above. Validate CRCs, source manifest, licenses and the
extracted module's clean install, manifest and tests. Candidate-mode CI remains
available without the flag.

Final names omit `-candidate`. Upload Windows, Companion module and Source ZIPs,
BUILD-INFO.json and SHA256SUMS.txt together so every checksum entry has an asset.
Create `main` at the verified release commit while retaining the development
branch locally. Push main, make it the public default, and create/push annotated
tag `v0.1.3` only if unused. Publish with GITHUB_RELEASE.txt as the release body.
Download all public assets independently and compare checksums and EXE identity.
Record publication separately from any independent interactive installation that
has not been performed. Keep older tags/assets intact. Rollback restores the old
EXE and matching backed-up Companion module/configuration, not a rewritten release.
