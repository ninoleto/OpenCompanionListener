#!/usr/bin/env python3
"""Build reproducible archives from HEAD, never from surrounding snapshots.

Requires Python 3.11+ and the Go compiler in release-toolchain.json on PATH.
Node dependencies are installed separately with npm ci on the Companion host.
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
    args = parser.parse_args()
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
    info = json_bytes({
        'status': 'RELEASE PACKAGE - OWNER AUTHORIZED' if args.release else 'LOCAL CANDIDATE - NOT APPROVED FOR PUBLICATION',
        'version': version, 'sourceCommit': commit, 'toolchain': config,
        'target': 'windows/amd64', 'goamd64': 'v1', 'cgo': False, 'goFlags': flags,
        'exeSHA256': digest(binary), 'moduleDependenciesBundled': False,
        'dependencyInstall': 'npm ci --omit=dev --ignore-scripts --no-audit --no-fund',
        'sourceFilesSHA256': {name: digest(data) for name, data in sorted(files.items())},
    })
    output.mkdir()
    prefix = f'OpenCompanionListener-v{version}' + ('' if args.release else '-candidate')
    # Windows archive includes complete source/docs so all local links resolve.
    archive(output / (prefix + '-Windows-x64.zip'), {
        **files, 'OpenCompanionListener.exe': binary, 'BUILD-INFO.json': info,
    })
    archive(output / (prefix + '-Source.zip'), {**files, 'BUILD-INFO.json': info})
    module_files = {name.removeprefix(MODULE): data for name, data in files.items() if name.startswith(MODULE)}
    module_files['INSTALL.md'] = files['docs/COMPANION-INSTALL.md']
    module_files['BUILD-INFO.json'] = info
    archive(output / (prefix + '-Companion-Module.zip'), {
        'companion-module-community-openlistener/' + name: data for name, data in module_files.items()
    })
    (output / 'BUILD-INFO.json').write_bytes(info)
    checksums = ''.join(f'{digest(path.read_bytes())}  {path.name}\n' for path in sorted(output.iterdir()))
    (output / 'SHA256SUMS.txt').write_text(checksums, encoding='utf-8', newline='\n')
    print(checksums, end='')


if __name__ == '__main__':
    main()
