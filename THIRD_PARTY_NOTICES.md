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

The v0.1.4 Companion TGZ bundles the runtime dependencies into `main.js` using
official `@companion-module/tools` 2.9.1. It includes the builder's generated
license notices in `LICENSE` and the upstream base notice in
`COMPANION-BASE-LICENSE.txt` (also kept in the source module's `licenses/`).
The direct runtime dependency is `@companion-module/base` 1.13.6 (MIT).
`package-lock.json` records the locked runtime and development dependencies.
Companion supplies Node.js; no Node.js executable or `node_modules` directory
is redistributed. End users import the TGZ without npm or dependency downloads.

## Artwork

The owner supplied and approved the replacement logo on 2026-11-12 in
`branding/approved/OpenCompanionListener-Approved-Logo/`. The approved design
has upper magenta/violet signal arcs and lower cyan/teal C-shaped arcs on a dark
navy square. This owner-provided project artwork is distributed with the
project's MIT license; approval/provenance is owner-attested, not an independent
authorship investigation.

`OpenCompanionListener.ico` is an unchanged copy of the supplied ICO (10 sizes,
16–256px). `listener/ui/icon.png` is the supplied 128px PNG, unchanged, for the
existing compact dashboard slot. `listener/icon_windows_amd64.syso` is generated
from that ICO using `tools/generate_icon_resource.py`; tray code uses embedded
icon group 1. No artwork was redrawn, recolored, mirrored or generated from the
traced vector files. Dashboard HTML/styles remain unchanged by this integration.

The local delivery package, original master and unused EPS/SVG/PNG variants are
preserved outside the tracked release inventory. Its SHA256SUMS entries were
verified before integration. The approved original `APPROVED-ORIGINAL-1254.png`
has SHA-256 `f0b3d58d11c31c24a61728489ee7fe994c9c52a3b41a422a155595ba253947ac`.
The old icon's 2026-11-11 provenance/checks remain historical in
`docs/AUDIT-v0.1.3.md` and Git history; they do not describe the new artwork.

| Project artwork | SHA-256 |
| --- | --- |
| `OpenCompanionListener.ico` | `87f1ab5a77dc69ddc7e3dd5d7d5fbcb0d5d089927062f3bc6f6d90a911c71932` |
| `listener/ui/icon.png` | `792b20827d33297631b8c192384253963563cebcf44d462431fd47f54e643943` |
| `listener/icon_windows_amd64.syso` | `82997efcbfc975cb2fcddf321fa70248a63cdbca954174147b790bc12f149f51` |
