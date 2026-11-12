# Release workflow (maintainer and contributors)

v0.1.4 is in preparation. Do not publish until final packages are verified.
Preserve the published v0.1.3 tag, binaries and assets. Follow
[VALIDATION-v0.1.4.md](VALIDATION-v0.1.4.md) for current evidence and blockers.

## Toolchain and checks

Use the existing versions in `tools/release-toolchain.json`: Go 1.27.4,
Node 22.23.4, Companion base 1.13.6 and Companion tools 2.9.1.
Python 3.11+ builds the archives. These tools are needed by developers only.
Windows race tests also require a compatible MinGW C compiler.
The release EXE uses CGO=0.

Run the commands in [DEVELOPMENT.md](DEVELOPMENT.md): Go race tests, both
vet configurations, Windows GUI build, module tests and manifest validation.
The TGZ builder needs the locked development dependencies, so use
`npm ci --ignore-scripts --no-audit --no-fund` without `--omit=dev` when
building packages. Reuse existing toolchains, caches and isolated test profiles.

Record existing security-scan evidence and its artifact/date limits; do not
describe a previous binary's scan as verification of a rebuilt binary.
Review the changed files for private data, run `git diff --check`, and commit
the reviewed source on the development branch before final packaging.

## Reproducible local packages

The builder requires a clean working tree and reads only `git archive HEAD`.
It checks Go and Node versions, dependency pins and the Go license. It builds
Windows/amd64, GOAMD64=v1, CGO=0 with
`-trimpath -buildvcs=false -ldflags='-s -w -H windowsgui'`, records the source
commit in BUILD-INFO.json and excludes private settings/caches/dependencies.

From the root, with the pinned tools on PATH:

```powershell
$moduleArgs = @('--module-package', 'dist/module-v0.1.4-verified/community-openlistener-0.1.4.tgz', '--module-sha256', '949ab82e84ebf9cf0ccdaff3d811bc4486af41adfcf0a20c3f823928a299b725')
py -3.13 tools/package_release.py --release --output dist/release-v0.1.4 @moduleArgs
if ($LASTEXITCODE) { throw 'First build failed' }
py -3.13 tools/package_release.py --release --output dist/release-v0.1.4-reproduction @moduleArgs
if ($LASTEXITCODE) { throw 'Second build failed' }
if ((Get-Content dist/release-v0.1.4/SHA256SUMS.txt -Raw) -cne (Get-Content dist/release-v0.1.4-reproduction/SHA256SUMS.txt -Raw)) {
    throw 'Builds differ'
}
Get-Content dist/release-v0.1.4/SHA256SUMS.txt | ForEach-Object {
    $expected, $name = $_ -split '  ', 2
    if ((Get-FileHash -Algorithm SHA256 (Join-Path dist/release-v0.1.4 $name)).Hash.ToLowerInvariant() -ne $expected) {
        throw "Checksum mismatch: $name"
    }
}
```

On Bash, use `python3` for the same builder commands, compare checksum files
with `cmp`, then run `sha256sum --check SHA256SUMS.txt` in the output directory.
Pass the same `--module-package` and `--module-sha256` values explicitly on Bash.
For this artwork-only completion, these paired arguments validate and copy the
previously verified TGZ unchanged; no Node/npm invocation, install or module
rebuild occurs. BUILD-INFO records its reused origin. The default builder path
still builds modules from source for future module changes and CI.
Use fresh directory names on subsequent attempts; the builder refuses overwrite.
Omit `--release` for candidate ZIP names. The flag only selects final filenames
and metadata; it does not publish or prove verification is complete.

ZIP entries have fixed ordering, timestamps and permissions. TGZ normalization
changes archive metadata only; payload bytes come from the official builder.
Text follows `git archive` attributes/configuration, which can convert line
endings. Two matching local builds establish local repeatability, not independent
cross-host reproduction.

## Exact release files

- `OpenCompanionListener-v0.1.4-Windows-x64.zip`: Windows EXE, full source, documentation and licenses.
- `community-openlistener-0.1.4.tgz`: bundled module, help and notices, imported through Companion's GUI.
- `OpenCompanionListener-v0.1.4-Source.zip`: complete committed source.
- `BUILD-INFO.json`: source commit, toolchain, flags and EXE/module/source hashes.
- `SHA256SUMS.txt`: hashes for both ZIPs, the TGZ and BUILD-INFO.json.

Validate ZIP CRCs, safe paths, source hashes, EXE GUI/x64 metadata and embedded
artwork against the approved assets. Validate the compressed module with
`python tools/validate_module_package.py <package.tgz>`.
Reject missing entrypoints even if the build exited successfully. Check
generated dependency notices and confirm no host-side npm install is needed.

## Windows and Companion evidence

Use [SMOKE_TEST.md](../SMOKE_TEST.md). Record owner reports as owner reports,
without repeating completed tests or marking unspecified cases passed. The owner
confirms the v0.1.4 TGZ imported through the GUI on the production Companion host,
retained connection settings, connected and ran most existing buttons.

Keep the working listener, settings, startup entry and Companion installation
untouched. The limited `node tools/smoke_listener.cjs <candidate.exe>` check
requires a free dashboard port and uses temporary settings; it must not stop a
running listener. It does not verify visible tray behavior or desktop input.
Record final artwork/EXE checks separately from earlier module tests.
Unobserved detailed cases and final-package identity remain explicit limits.

## GitHub publication

Only after final verification, publish the reviewed commit as a new annotated
`v0.1.4` tag and release, with `GITHUB_RELEASE.txt` as the body and all five
files above. Do not move the v0.1.3 tag or replace/delete its assets.
Verify the release commit's GitHub Actions, rendered documentation and download
links. Download the uploaded files independently and compare SHA-256 checksums.
Record publication and independent installation as separate facts.

Rollback uses the previous Windows EXE and matching Companion module/configuration
backup. Never rewrite a published release to perform a rollback.
