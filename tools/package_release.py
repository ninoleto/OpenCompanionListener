#!/usr/bin/env python3
"""Build reproducible archives from HEAD, never from surrounding snapshots.

Requires Python 3.11+ and the Go compiler in release-toolchain.json on PATH.
Node and npm are needed only when rebuilding the module. A verified TGZ can be
reused with --module-package and --module-sha256. End users import the bundled TGZ.
"""
import argparse
import hashlib
import io
import json
import os
from pathlib import Path, PurePosixPath
import subprocess
import tempfile
import zipfile
import shutil
from validate_module_package import inspect_package, normalize_package

ROOT = Path(__file__).resolve().parents[1]
MODULE = 'companion-dev-modules/companion-module-community-openlistener/'


def run(*args, cwd=ROOT, env=None):
    return subprocess.check_output(args, cwd=cwd, env=env)


def json_bytes(value):
    return (json.dumps(value, indent=2, sort_keys=True) + '\n').encode()


def digest(data):
    return hashlib.sha256(data).hexdigest()


def archive(path, files):
    # STORE avoids zlib-version differences. Fixed metadata/order makes ZIPs
    # identical across build directories and host time zones.
    with zipfile.ZipFile(path, 'x', compression=zipfile.ZIP_STORED) as out:
        for name, data in sorted(files.items()):
            info = zipfile.ZipInfo(name, date_time=(1980, 1, 1, 0, 0, 0))
            info.create_system = 3
            info.external_attr = 0o100644 << 16
            out.writestr(info, data)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--output', required=True, help='New directory under dist/ (must not exist)')
    parser.add_argument('--release', action='store_true', help='Use final release names/metadata for an owner-authorized publication; does not upload')
    parser.add_argument('--module-package', type=Path, help='Reuse an already verified TGZ without rebuilding or modifying it')
    parser.add_argument('--module-sha256', help='Required SHA-256 of the reused TGZ')
    args = parser.parse_args()
    if bool(args.module_package) != bool(args.module_sha256):
        parser.error('--module-package and --module-sha256 must be used together')
    reused_module = None
    if args.module_package:
        reused_module = args.module_package.read_bytes()
        if digest(reused_module) != args.module_sha256.lower():
            parser.error('reused module SHA-256 mismatch')
    output = (ROOT / args.output).resolve()
    dist = (ROOT / 'dist').resolve()
    if not output.is_relative_to(dist) or output == dist:
        parser.error('output must be a new child directory of this repository\'s dist/')
    if output.exists():
        parser.error('output exists; use a fresh directory to preserve previous candidates')
    if run('git', 'status', '--porcelain', '--untracked-files=normal').strip():
        parser.error('commit reviewed source changes before packaging; working tree must be clean')
    commit = run('git', 'rev-parse', 'HEAD').decode().strip()
    with zipfile.ZipFile(io.BytesIO(run('git', 'archive', '--format=zip', 'HEAD'))) as source:
        files = {item.filename: source.read(item) for item in source.infolist() if not item.is_dir()}
    for name in files:
        path = PurePosixPath(name)
        if path.is_absolute() or '..' in path.parts or '\\' in name:
            raise RuntimeError('unsafe source path')
        if path.name.lower() in ('settings.json', '.env') or any(part in ('.git', 'node_modules', '.cache', 'dist') for part in path.parts):
            raise RuntimeError(f'private/generated file tracked: {name}')
    config = json.loads(files['tools/release-toolchain.json'])
    env = os.environ.copy()
    env.update(GOENV='off', GOTOOLCHAIN='local', GOTELEMETRY='off', GOOS='windows',
               GOARCH='amd64', GOAMD64='v1', CGO_ENABLED='0', GOFLAGS='',
               GOEXPERIMENT='', GOWORK='off')
    go_version = run('go', 'env', 'GOVERSION', env=env).decode().strip()
    if go_version != config['go']:
        raise RuntimeError(f'expected {config["go"]}, found {go_version}')
    goroot = Path(run('go', 'env', 'GOROOT', env=env).decode().strip())
    notice = (goroot / 'LICENSE').read_bytes().replace(b'\r\n', b'\n')
    if notice != files['licenses/GO-LICENSE.txt']:
        raise RuntimeError('Go toolchain license differs from the included notice')
    pkg = json.loads(files[MODULE + 'package.json'])
    manifest = json.loads(files[MODULE + 'companion/manifest.json'])
    lock = json.loads(files[MODULE + 'package-lock.json'])
    version = pkg['version']
    assert version == manifest['version'] == lock['version']
    assert pkg['dependencies']['@companion-module/base'] == config['companionBase']
    assert pkg['devDependencies']['@companion-module/tools'] == config['companionTools']
    if reused_module is None and run('node', '--version').decode().strip() != config['node']:
        raise RuntimeError('Node version does not match release-toolchain.json')
    npm = shutil.which('npm.cmd' if os.name == 'nt' else 'npm')
    if reused_module is None and not npm:
        raise RuntimeError('npm is required on the build machine')
    flags = ['-trimpath', '-buildvcs=false', '-ldflags=-s -w -H windowsgui']
    dist.mkdir(exist_ok=True)
    # A fresh source path on each invocation also checks path independence.
    with tempfile.TemporaryDirectory(prefix='build-', dir=dist) as temp:
        source_root = Path(temp)
        for name, data in files.items():
            target = source_root / name
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_bytes(data)
        executable = source_root / 'OpenCompanionListener.exe'
        run('go', 'vet', './...', cwd=source_root / 'listener', env=env)
        run('go', 'build', *flags, '-o', str(executable), '.', cwd=source_root / 'listener', env=env)
        binary = executable.read_bytes()
        module_root = source_root / MODULE
        build_env = os.environ.copy()
        # The upstream tool invokes npx through Windows PowerShell. Scope this
        # to the build process; never alter a user/machine execution policy.
        if os.name == 'nt':
            build_env['PSExecutionPolicyPreference'] = 'Bypass'
        tgz_name = f'community-openlistener-{version}.tgz'
        if reused_module is None:
            run(npm, 'ci', '--ignore-scripts', '--no-audit', '--no-fund', cwd=module_root, env=build_env)
            run(npm, 'run', 'validate', cwd=module_root, env=build_env)
            run(npm, 'test', cwd=module_root, env=build_env)
            run(npm, 'run', 'build', cwd=module_root, env=build_env)
            module_binary = normalize_package((module_root / tgz_name).read_bytes(), version)
        else:
            module_binary = reused_module
        module_files = inspect_package(module_binary, version)
    info = json_bytes({
        'status': 'LOCAL RELEASE PACKAGE - NOT PUBLISHED' if args.release else 'LOCAL CANDIDATE - NOT APPROVED FOR PUBLICATION',
        'version': version, 'sourceCommit': commit, 'toolchain': config,
        'target': 'windows/amd64', 'goamd64': 'v1', 'cgo': False, 'goFlags': flags,
        'exeSHA256': digest(binary), 'moduleDependenciesBundled': True,
        'modulePackageSHA256': digest(module_binary),
        'modulePackageOrigin': 'verified TGZ reused byte-for-byte' if reused_module is not None else 'built from source',
        'moduleFilesSHA256': {name: digest(data) for name, data in sorted(module_files.items())},
        'moduleInstall': 'Companion Modules > Import module package',
        'sourceFilesSHA256': {name: digest(data) for name, data in sorted(files.items())},
    })
    output.mkdir()
    prefix = f'OpenCompanionListener-v{version}' + ('' if args.release else '-candidate')
    # Windows archive includes complete source/docs so all local links resolve.
    archive(output / (prefix + '-Windows-x64.zip'), {
        **files, 'OpenCompanionListener.exe': binary, 'BUILD-INFO.json': info,
    })
    archive(output / (prefix + '-Source.zip'), {**files, 'BUILD-INFO.json': info})
    (output / tgz_name).write_bytes(module_binary)
    (output / 'BUILD-INFO.json').write_bytes(info)
    checksums = ''.join(f'{digest(path.read_bytes())}  {path.name}\n' for path in sorted(output.iterdir()))
    (output / 'SHA256SUMS.txt').write_text(checksums, encoding='utf-8', newline='\n')
    print(checksums, end='')


if __name__ == '__main__':
    main()
