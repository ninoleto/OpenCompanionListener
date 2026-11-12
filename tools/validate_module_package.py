#!/usr/bin/env python3
"""Inspect a built Companion tgz without trusting a successful build exit code."""
import argparse
import gzip
import hashlib
import io
import json
from pathlib import Path, PurePosixPath
import tarfile


def inspect_package(data, version):
    files = {}
    seen = set()
    with tarfile.open(fileobj=io.BytesIO(data), mode='r:gz') as archive:
        for member in archive:
            name = member.name.rstrip('/')
            path = PurePosixPath(name)
            if (path.is_absolute() or '..' in path.parts or '\\' in name
                    or ':' in name or not path.parts or path.parts[0] != 'pkg'
                    or name in seen or not (member.isfile() or member.isdir())):
                raise ValueError(f'Unsafe or duplicate archive member: {name}')
            seen.add(name)
            if member.isfile():
                if member.size > 5_000_000:
                    raise ValueError(f'Unexpectedly large package file: {name}')
                files[name] = archive.extractfile(member).read()
    expected = {'pkg/main.js', 'pkg/package.json', 'pkg/companion/manifest.json',
                'pkg/companion/HELP.md', 'pkg/LICENSE', 'pkg/COMPANION-BASE-LICENSE.txt'}
    if set(files) != expected:
        raise ValueError(f'Unexpected package contents: missing={expected - files.keys()}, extra={files.keys() - expected}')
    manifest = json.loads(files['pkg/companion/manifest.json'])
    package = json.loads(files['pkg/package.json'])
    if manifest['id'] != 'community-openlistener' or manifest['version'] != version or package['version'] != version:
        raise ValueError('Package identity/version mismatch')
    runtime = manifest['runtime']
    if runtime != {'type': 'node22', 'api': 'nodejs-ipc', 'apiVersion': '1.13.6', 'entrypoint': '../main.js'}:
        raise ValueError(f'Unexpected packaged runtime: {runtime}')
    if package.get('dependencies') or package.get('type') != 'commonjs':
        raise ValueError('Package requires external dependencies or wrong module format')
    if len(files['pkg/main.js']) < 1000:
        raise ValueError('Missing or empty compiled module')
    for notice, marker in [('pkg/LICENSE', b'Open Companion Listener contributors'),
                           ('pkg/COMPANION-BASE-LICENSE.txt', b'Bitfocus AS')]:
        if marker not in files[notice] or b'Permission is hereby granted' not in files[notice]:
            raise ValueError(f'Missing license notice: {notice}')
    return files


def normalize_package(data, version):
    # Only archive metadata changes; every payload byte comes from the official
    # companion-module-build output. This makes clean builds comparable.
    files = inspect_package(data, version)
    stream = io.BytesIO()
    with tarfile.open(fileobj=stream, mode='w', format=tarfile.USTAR_FORMAT) as archive:
        for name, content in sorted(files.items()):
            entry = tarfile.TarInfo(name)
            entry.mode = 0o644
            entry.size = len(content)
            archive.addfile(entry, io.BytesIO(content))
    output = io.BytesIO()
    with gzip.GzipFile(fileobj=output, mode='wb', filename='', mtime=0) as zipped:
        zipped.write(stream.getvalue())
    normalized = output.getvalue()
    if inspect_package(normalized, version) != files:
        raise ValueError('Package payload changed during normalization')
    return normalized


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('package', type=Path)
    parser.add_argument('--version', default='0.1.4')
    args = parser.parse_args()
    data = args.package.read_bytes()
    files = inspect_package(data, args.version)
    print(f'PASS: {len(files)} package files; bundled entrypoint, manifest, help and notices; SHA256 {hashlib.sha256(data).hexdigest()}')
