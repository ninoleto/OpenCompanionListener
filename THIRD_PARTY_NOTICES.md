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

`OpenCompanionListener.ico`, `listener/ui/icon.png`, and the derived
`listener/icon_windows_amd64.syso` came from the supplied v0.1.3 snapshot.
The resource generator is included in `tools/generate_icon_resource.py`.
Original artwork authorship and permission to redistribute remain unconfirmed.
**Publication remains on hold until the owner records the origin and rights.**
No author or license grant for the artwork is inferred from its presence here.
