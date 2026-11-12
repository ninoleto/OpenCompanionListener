# Developer guide

For ordinary installation, use the [README](../README.md). This guide is for building or modifying the source. Read [CONTRIBUTING.md](../CONTRIBUTING.md), [architecture](ARCHITECTURE.md) and [security](../SECURITY.md); coding agents must also read [AGENTS.md](../AGENTS.md) and [HANDOFF.md](../HANDOFF.md).

## Tools

Release pins are in `tools/release-toolchain.json`: Go 1.27.4, Node 22.23.4, Companion base 1.13.6 and official `@companion-module/tools` 2.9.1. That tooling supports API 1; tools 3.x requires API 2. Python 3.11+ is needed for archive tooling. Windows race checks require CGO and a compatible x64 MinGW compiler; the shipped EXE uses CGO=0.

## Tests and builds

PowerShell, from the repository root:

```powershell
$env:CGO_ENABLED = '1'
Push-Location listener
go test -race ./...
if ($LASTEXITCODE) { throw 'Go tests failed' }
go vet ./...
if ($LASTEXITCODE) { throw 'Go vet failed' }
$env:GOOS = 'windows'
$env:GOARCH = 'amd64'
$env:CGO_ENABLED = '0'
go vet ./...
if ($LASTEXITCODE) { throw 'Windows vet failed' }
go build -trimpath -buildvcs=false -ldflags='-s -w -H windowsgui' -o ../OpenCompanionListener.exe .
if ($LASTEXITCODE) { throw 'Windows build failed' }
Pop-Location
Push-Location companion-dev-modules/companion-module-community-openlistener
npm.cmd ci --ignore-scripts --no-audit --no-fund
if ($LASTEXITCODE) { throw 'Dependency install failed' }
npm.cmd run validate
if ($LASTEXITCODE) { throw 'Manifest invalid' }
npm.cmd test
if ($LASTEXITCODE) { throw 'Module tests failed' }
# The official builder invokes npx in a child Windows PowerShell process.
$previousPolicy = $env:PSExecutionPolicyPreference
try {
    $env:PSExecutionPolicyPreference = 'Bypass'
    npm.cmd run build
    if ($LASTEXITCODE) { throw 'Module build failed' }
} finally { $env:PSExecutionPolicyPreference = $previousPolicy }
py -3.13 ../../tools/validate_module_package.py community-openlistener-0.1.4.tgz
if ($LASTEXITCODE) { throw 'Package invalid' }
Pop-Location
```

On Bash:

```bash
set -eu
(cd listener && go test -race ./... && go vet ./...)
(cd listener && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go vet ./...)
(cd listener && GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -buildvcs=false -ldflags='-s -w -H windowsgui' -o ../OpenCompanionListener.exe .)
cd companion-dev-modules/companion-module-community-openlistener
npm ci --ignore-scripts --no-audit --no-fund
npm run validate
npm test
npm run build
python3 ../../tools/validate_module_package.py community-openlistener-0.1.4.tgz
```

The official builder produces `pkg/` and `community-openlistener-0.1.4.tgz`. The package validator checks the compressed artifact: entrypoint, manifest, bundled dependencies, help, licenses and allowed contents. A build exit code alone is insufficient. Import the actual file into a clean Companion installation using its graphical interface before distribution. See the [official packaging guide](https://companion.free/for-developers/module-development/module-lifecycle/module-packaging/).

## Source development in Companion

Only developers need a development-folder configuration. Use a separate Companion profile with no production connections or surfaces. Install locked dependencies in `companion-dev-modules/companion-module-community-openlistener`, then set the test launcher's development-module parent to `companion-dev-modules`. Do not expose both a source copy and an imported package with the same ID during package tests. Use GUI import for release validation.

The listener has no external Go dependencies. Non-Windows files are compile/test stubs, not a Linux desktop implementation. Never test unreviewed input or startup changes in a production account. See [SMOKE_TEST.md](../SMOKE_TEST.md) and [release instructions](RELEASING.md).
