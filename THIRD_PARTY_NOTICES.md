# Third-party notices

Open Companion Listener's own source is MIT licensed; see `LICENSE` and the
module's `LICENSE`. These notices do not replace those licenses.

## Windows executable

The executable statically includes the Go runtime and standard library from
Go 1.27.4. Its BSD-style redistribution notice is reproduced verbatim in
`licenses/GO-LICENSE.txt`, copied from that compiler's `LICENSE` file. The
package builder verifies it against the selected toolchain and includes it
in the Windows archive. There are no external Go module dependencies.

## Companion module

The Companion ZIP contains this project's source and `package-lock.json`;
**it does not bundle Node.js or node_modules**. Install the locked runtime
packages with `npm ci --omit=dev --ignore-scripts --no-audit --no-fund`.
The direct dependency is `@companion-module/base` 1.13.6 (MIT). The lockfile
records exact transitive versions, integrity hashes, and declared licenses.
Dependency downloads are required for this install step; normal use is local.

If a future distributor bundles node_modules, they must include each
dependency's applicable copyright/license notices, including the upstream
Companion base notice absent from its current npm tarball. A package.json
license field alone is not a substitute for the full notice. This repository's
builder deliberately does not create that kind of bundle.

## Artwork

The owner confirmed on 2026-11-11 that the application icon was created
specifically for this project during development using programmatically drawn
graphics, and was not copied from a third-party product. At the owner's
direction, the project artwork in `OpenCompanionListener.ico`,
`listener/ui/icon.png`, and the derived `listener/icon_windows_amd64.syso` is
included under the project's MIT license and contributor copyright notice in
`LICENSE`. This provenance statement is based on the owner's confirmation;
the original drawing script was not found in the reviewed source.

Verification on 2026-11-11 established that these files are byte-identical
across the original project snapshot, supplied source ZIP, current checkout,
and corrected candidate Windows/source ZIPs. The dashboard PNG is exactly the
ICO's 128-pixel image. All seven icon images and the icon group embedded in the
corrected EXE match the ICO. Running `tools/generate_icon_resource.py` in an
isolated directory reproduced the checked-in resource byte-for-byte. The
tracked artwork inventory contains these three files; no additional artwork
was identified. See `docs/AUDIT-v0.1.3.md` for package and executable evidence.

| Project artwork | SHA-256 |
| --- | --- |
| `OpenCompanionListener.ico` | `7b2b5a21fa9921180875ee6b3f9a8d8a30a2672ccce5506cfbc0a1c5d2a578dd` |
| `listener/ui/icon.png` | `fe9f9b869c00d1493010c033d0f000b622124eb3108cd7f3ebc6e8547d801c50` |
| `listener/icon_windows_amd64.syso` | `7f717b5729784f1755b12264b553525c010a45fac0969005b9dbcc6e466e68ea` |
