# Companion module installation and rollback

For final validation on an existing ZORIN installation, use the detailed
[ZORIN staging/migration/rollback runbook](ZORIN-VALIDATION.md). It keeps the
installed module intact, tests copied configuration in isolation, and describes
a future switch to a separate module parent instead of replacing the old folder.

The module targets Companion **4.1 or later with API 1.13 and the node22 runtime**.
The [official compatibility matrix](https://bitfocus.github.io/companion-module-base/)
lists 4.1 as the first host supporting base 1.13. This is an API minimum, not
completed host certification: actual Companion import/migration testing remains
a release gate. Use Node **22.8+ within the 22.x line** for installation; the
candidate was tested with Node 22.23.4 and base 1.13.6. Node 20/24 are not the
declared module runtime.

The ZIP is a developer/extra module folder, **not a packaged in-app upload**.
It contains source, tests, manifest, license and a locked dependency tree;
node_modules and Node.js are not bundled. Download access to the npm registry is
needed once for installation. The Windows listener requires neither npm nor Node.

## Prepare without touching an active installation

1. Extract the candidate ZIP to a staging location.
2. Enter `companion-module-community-openlistener` in PowerShell or Bash.
3. Check `node --version` reports 22.x, then run each command and stop on failure:

   ```text
   npm ci --omit=dev --ignore-scripts --no-audit --no-fund
   npm run validate
   npm test
   ```

4. Export the existing Companion configuration before upgrades. Prepare on an
   isolated test host first; do not test this candidate against live buttons.
5. Use the developer/extra module **parent** directory configured in Companion.
   The folder inside it must be named `companion-module-community-openlistener`.

## First installation (Linux Companion host example)

Run only after the staging checks pass. Change the two example paths; neither
is a deployment default. `set -e` makes a failed copy or directory check stop.

```bash
set -eu
SRC="/path/to/staged/companion-module-community-openlistener"
PARENT="$HOME/companion-module-active"
DEST="$PARENT/companion-module-community-openlistener"
test -f "$SRC/companion/manifest.json"
test -d "$SRC/node_modules"
mkdir -p "$PARENT"
test ! -e "$DEST"
cp -a "$SRC" "$DEST"
```

Then reload the test Companion host and add the Open Companion Listener
connection. Enter the Windows listener IP, port and private pairing key.

## Upgrade (Linux Companion host example)

Stop the test Companion host before changing its module directory. Export its
configuration first. Set `SRC` and `PARENT` as above, then run:

```bash
set -eu
SRC="/path/to/staged/companion-module-community-openlistener"
PARENT="$HOME/companion-module-active"
DEST="$PARENT/companion-module-community-openlistener"
BACKUP="$HOME/openlistener-module-backup-$(date +%Y%m%d-%H%M%S)"
test -f "$SRC/companion/manifest.json"
test -d "$SRC/node_modules"
test -d "$DEST"
test ! -e "$BACKUP"
mv "$DEST" "$BACKUP"
if ! cp -a "$SRC" "$DEST"; then
  echo "Copy failed. Keep Companion stopped; preserve the partial folder and restore $BACKUP."
  exit 1
fi
echo "Previous module saved at $BACKUP"
```

On Windows, perform the same staging checks, stop Companion, move the existing
module folder to a separate backup location, then copy the staged folder into
the configured parent. Do not merge files into an old node_modules tree.

Reload only when safe; restarting Companion interrupts connections. Confirm the
module loads, all 38 optional presets appear, and existing buttons still work.
Test raw and expression-based legacy shortcuts and both scroll directions.
The additive migration preserves old signed scroll expressions in an advanced
mode within the single scroll action. Invalid old values remain editable and
are rejected at execution; no automatic guess is made about their intent.

If the old migration already converted an expression to a literal, its original
expression cannot be recovered from the converted action. Restore it from the
pre-upgrade configuration export and test the fixed migration.

## Rollback

Stop Companion, move the candidate module aside, restore the saved module folder
and the **matching pre-upgrade configuration export**, then restart and verify
the old buttons. Configuration migrations may already have been saved, so
restoring the module alone is insufficient. Never share pairing keys in logs.
