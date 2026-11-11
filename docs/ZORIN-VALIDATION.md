# ZORIN: stage, validate, migrate and roll back safely

This is an operator runbook, not a record of a ZORIN installation. No commands
in this document have been run on the existing ZORIN host by this validation
session. The existing module directory, Companion configuration, connections,
pairing keys and devices must remain intact. Publication is separately blocked
until the owner explicitly approves it.

## 1. Record the baseline and make private backups

1. Record the installed Companion version, how it is started (desktop app,
   system service, user service or container), the configured developer/extra
   module **parent** path, and the persistent Companion configuration location.
   Read these from the installation; do not guess a service name or data path.
2. Export the full Companion configuration using that version's export facility.
   Keep it outside this source repository. Record page/button counts, connection
   IDs/names, surface assignments and module versions. Do not delete/recreate the
   Open Listener connection: its identity is used by existing buttons.
3. Check whether the export includes connection secrets and surface/device
   settings. Do not assume it does. Keep a private backup of the full persistent
   Companion data too: use an application-consistent snapshot or, during an
   agreed maintenance window, stop the confirmed Companion instance, copy its
   data, and restart the unchanged installation. Never take an unverified live
   database copy and treat it as a tested recovery backup.
4. Treat exports and data backups as secrets. Preserve permissions, keep them
   private, and do not upload them or print pairing keys in validation logs.
5. Copy the existing module folder as shown below; this does not move or replace
   it. Verify the backup archive before continuing. A module backup alone cannot
   reverse a configuration migration.

Use a new Bash terminal, replace the two example paths with verified absolute
paths, and keep this shell for the following preparation blocks. These commands
write only to a fresh review directory in your home folder.

```bash
set -euo pipefail
umask 077
ACTIVE_PARENT="/absolute/path/to/configured-extra-module-parent"
DOWNLOADS="/absolute/path/to/copied-candidate-downloads"
MODULE="companion-module-community-openlistener"
ACTIVE_PARENT=$(realpath -e -- "$ACTIVE_PARENT")
ACTIVE="$ACTIVE_PARENT/$MODULE"
test -d "$ACTIVE"
test -f "$ACTIVE/companion/manifest.json"
RUN=$(mktemp -d "$HOME/openlistener-review-XXXXXX")
mkdir "$RUN/backups" "$RUN/staging" "$RUN/logs"
cp -a -- "$ACTIVE" "$RUN/backups/$MODULE"
tar -C "$RUN/backups" -czf "$RUN/backups/module-before.tar.gz" "$MODULE"
tar -tzf "$RUN/backups/module-before.tar.gz" > "$RUN/logs/backup-contents.txt"
(cd "$RUN/backups" && sha256sum module-before.tar.gz > module-before.sha256 && sha256sum -c module-before.sha256)
printf 'Private review directory: %s\n' "$RUN"
```

Keep the original module and the full config/data backup together as the
pre-upgrade recovery set. If its restoration cannot be verified in an isolated
environment, stop before changing the production host.

## 2. Verify and extract the candidate away from the active module

Copy `OpenCompanionListener-v0.1.3-candidate-Companion-Module.zip` and its matching
`SHA256SUMS.txt` from `dist/publication-candidate/` to the download location.
Do not mix checksum files from different candidate builds. The candidate source
commit is `58b9f14821326cae0f6d92a9f1a379babff14c69`.

```bash
ZIP="OpenCompanionListener-v0.1.3-candidate-Companion-Module.zip"
SUMS="$DOWNLOADS/SHA256SUMS.txt"
test -f "$SUMS"
test -f "$DOWNLOADS/$ZIP"
awk -v file="$ZIP" '$2 == file {print; count++} END {if (count != 1) exit 1}' "$SUMS" > "$RUN/logs/module.sha256"
(cd "$DOWNLOADS" && sha256sum -c "$RUN/logs/module.sha256")
unzip -q "$DOWNLOADS/$ZIP" -d "$RUN/staging"
STAGED="$RUN/staging/$MODULE"
test -f "$STAGED/package-lock.json"
test -f "$STAGED/companion/manifest.json"
```

## 3. Install the locked dependencies and validate in staging

Use Node 22.8+ within the 22.x line; the candidate's recorded test runtime is
22.23.4. Do not replace a system Node installation used by other software merely
to run this test. Use an existing suitable runtime or a separately installed
runtime selected only for this shell. Stop if it is unavailable.

Companion must support API 1.13 and node22. The upstream
[compatibility matrix](https://bitfocus.github.io/companion-module-base/) lists
Companion 4.1 as the first host supporting API/base 1.13; that alone is not
evidence that your particular host passed migration tests.

```bash
node --version
npm --version
node -e 'const [major,minor]=process.versions.node.split(".").map(Number); if(major!==22 || minor<8) process.exit(1)'
(
  cd "$STAGED"
  sha256sum package-lock.json > "$RUN/logs/lock-before.sha256"
  npm ci --omit=dev --ignore-scripts --no-audit --no-fund 2>&1 | tee "$RUN/logs/npm-ci.txt"
  sha256sum -c "$RUN/logs/lock-before.sha256"
  npm run validate 2>&1 | tee "$RUN/logs/manifest.txt"
  npm test 2>&1 | tee "$RUN/logs/tests.txt"
)
```

Expected: base **1.13.6**, module/manifest **0.1.3**, **28 passing tests**, no
failures, unchanged lockfile. Stop at any failure. Run `npm ci` only in staging:
it removes an existing node_modules tree before installing from the lockfile,
as documented by [npm](https://docs.npmjs.com/cli/v10/commands/npm-ci/).
Do not use `npm update`, `npm audit fix` or `sudo npm` for this candidate.

## 4. Test migration using an isolated Companion installation

1. Use a separate ZORIN test VM/machine and a copied Companion data/profile,
   never the production data directory. Keep its surface/device registrations
   separate. Do not attach or re-pair production control surfaces for this test.
2. Restrict its network **before** importing production configuration: imported
   connections may reconnect immediately. Permit only the isolated Windows test
   listener and required test services, with production devices unreachable.
   Keep the real ZORIN host and its paired devices running unchanged.
3. Point the test host's extra-module path to a separate parent containing the
   validated candidate folder. Do not let it discover both old and candidate
   folders with the same `community-openlistener` ID. Use the normal discovery
   path, not an in-app module package upload.
4. Restore/import the copied pre-upgrade configuration on that isolated host.
   Update only its **test copy** of the Open Listener connection to the isolated
   Windows endpoint and test pairing key. Leave the production connection intact.
5. Record pre/post page counts, button locations, connection IDs and action
   options. Let Companion run the included upgrade scripts; do not manually
   edit migration indexes in a live configuration.
6. Exercise the cases below on a scratch Windows document and test page. Export
   the migrated test configuration, restart the test host, and repeat checks to
   ensure migration persisted without duplicating or losing actions.

| Existing action | Expected result |
| --- | --- |
| `shortcut_builder` | Becomes the normal shortcut builder; modifiers and key preserved. |
| Text `hotkey`, `key_down`, `key_up` | Simple shortcuts use builder options; unusual combinations remain advanced text. Press/hold/release semantics remain intact. |
| Wrapped/expression shortcut values | Expressions remain expressions, including literal-looking numeric contents; evaluated values reach the intended action. |
| Literal `scroll`, both axes and signs | Single scroll action retains direction and magnitude (1–20). |
| Dynamic/signed legacy scroll | Advanced mode preserves the signed value/expression; the additive repair also covers actions whose old upgrade already ran. |
| Invalid legacy scroll | Remains editable and reports an error; it does not silently execute an invented value. |
| Existing `scroll_slider`, file/URL/utility actions | Saved options and behavior unchanged. |
| Existing pages/buttons/presets | No automatic presets added, no duplicate action selector, no missing connection references. |

If a prior buggy migration already discarded expression text, recover it from
the pre-upgrade export. The candidate cannot reconstruct data that is absent.
Use synthetic fixtures for missing legacy cases; do not fabricate a production
migration result from unit tests. Complete [SMOKE_TEST.md](../SMOKE_TEST.md).

## 5. Future activation on ZORIN without overwriting the old module

This phase is **not performed by the current validation session**. Use it only
when the owner chooses a maintenance window after the isolated tests pass.

1. Refresh the private configuration/data/module backups and stop changes to
   buttons during the maintenance window. Record the original extra-module path.
2. Build a separate candidate parent. If the current parent holds other custom
   modules, copy them too so changing the parent does not make them disappear:

   ```bash
   CANDIDATE_PARENT="$RUN/modules-candidate"
   test ! -e "$CANDIDATE_PARENT"
   mkdir "$CANDIDATE_PARENT"
   cp -a -- "$ACTIVE_PARENT/." "$CANDIDATE_PARENT/"
   # Move only the copied old module, never the installed original.
   test ! -e "$RUN/copied-old-module"
   mv -- "$CANDIDATE_PARENT/$MODULE" "$RUN/copied-old-module"
   cp -a -- "$STAGED" "$CANDIDATE_PARENT/$MODULE"
   ```

3. Inspect any symlinks/permissions and ensure the Companion service account can
   read the candidate parent. The private review directory may not be readable
   by a different service account. Have the operator place a copy in a new
   appropriately owned deployment directory; do not broadly relax backup
   permissions or guess ownership. Keep all backups outside discovery paths.
4. Stop the **confirmed** Companion instance using its normal launcher/service
   controls. Point that instance's supported extra-module-path setting at the
   new parent, preserving other launch settings. If the installation cannot
   safely select a separate parent, stop and tailor the deployment procedure;
   do not fall back to overwriting the installed folder.
5. Restart the same Companion instance with its existing configuration and
   device identities. Do not recreate connections or rotate production keys.
   Confirm the candidate is loaded from the new path, then test selected existing
   buttons before enabling normal operation. The old module folder stays intact.
6. After successful checks, export a separate post-upgrade recovery snapshot.
   Keep the pre-upgrade set until rollback is no longer needed.

## 6. Rollback

On the isolated host, practice rollback first. On a future production activation,
stop the confirmed Companion instance, restore its original extra-module parent
setting, and restore the **matching pre-upgrade configuration/data backup** using
the installation's supported restore method. Restore full data only while that
instance is stopped. Do not overlay a live database or import an unrelated backup.

The original module was never overwritten. Move the candidate parent out of
discovery if necessary, restart the original installation, and check original
connection IDs, device/surface assignments, page/button counts and known buttons.
Retain candidate logs privately for diagnosis. Restoring only the old module
without its matching configuration may leave migrated actions incompatible.
