# Install the Companion module

Install the bundled v0.1.4 package using the steps below. See [release verification](VALIDATION-v0.1.4.md) for owner-confirmed live testing and its scope.

Use Companion 4.1 or later with the node22 runtime. Companion supplies the runtime and the `.tgz` contains the module's dependencies. End users do not need a terminal, Node.js installation, npm commands or a developer folder.

## First installation

1. Download `community-openlistener-0.1.4.tgz` from [GitHub Releases](https://github.com/ninoleto/OpenCompanionListener/releases). Do not extract it.
2. Open Companion's web interface on the Companion host itself. Select **Modules**, then **Import module package**.
3. Select the downloaded file and confirm the import. Check that **Open Companion Listener**, version **0.1.4**, appears in the installed modules.
4. On **Connections**, add **Open Companion Listener** under **Community**.
5. Enter the Windows PC's IPv4 address, API port (default `17992`) and the private pairing key shown in the listener's **Settings**. For Companion on the same PC, use `127.0.0.1`.
6. Save, wait for the connection to become OK, then add an action to a scratch button. Focus a harmless test document on Windows and test a key press.
7. Under **Presets**, select the connection. There are 38 optional presets across five categories. Drag only the ones you want onto your pages.

For Windows installation, network settings and firewall configuration, see the [README](../README.md). The module cannot control a macOS or Linux desktop; Companion itself can run on a different supported host.

## Updating an imported package

Export the full Companion configuration first and keep a private recovery backup. Import the newer `.tgz` through **Modules**, then select that version for the existing connection if Companion offers a version selector. Preserve the connection itself. Check the loaded version, connection status and existing buttons after a restart. Do not assume that importing another version automatically switches every connection.

## Transition from the v0.1.3 development-folder installation

The first release used a source folder rather than an importable package. Its `community-openlistener` connection ID, action IDs, options, presets and two upgrade scripts are retained in v0.1.4. The owner confirms GUI import, retained connection settings and working existing buttons on the production host; full migration and rollback coverage is not claimed. Unit tests cover legacy raw/wrapped values and expressions. See [verification](VALIDATION-v0.1.4.md) for the exact scope.

1. Export the configuration and preserve the old module folder and its dependencies. Keep a full private Companion data backup using the host's supported backup procedure. Backups may contain the pairing key.
2. Test the transition on an isolated Companion profile first. Do not delete or recreate the existing connection.
3. Stop that test instance. Move only its copied old module out of the configured development folder, or disable that folder in the test instance if it contains no other needed modules. Keep the backup outside module discovery. This one-time cleanup is for the old installation; the new package needs no development-folder setup.
4. Restart the test instance, import the `.tgz` using **Modules**, and select version **0.1.4** for the existing connection. Keep the same connection ID and label.
5. Confirm saved actions, all presets, pairing and page/button positions. Restart again and repeat representative buttons. Restore the pre-upgrade backup in isolation to check rollback before changing a working installation.

Do not proceed on a production host if either transition or rollback fails. Record sanitized Companion logs and affected action options. Already-lost legacy expression text can only be recovered from a pre-upgrade backup.

## Rollback

Stop Companion and restore the previous module together with its matching configuration/data backup using the host's restore procedure. For a previous development-folder installation, restore its original folder setting. For an imported package, select its saved prior version. Restart and verify the same connection and buttons. Restoring only module code does not reverse saved configuration changes. Keep the previous Windows EXE as well if updating the listener.

## Import and connection problems

- Use the `.tgz`, not the source ZIP or Windows ZIP. A source archive cannot be imported as a module.
- If **Import module package** is missing, verify the Companion version. Newer versions permit local imports and can restrict imports from another computer; use the local interface.
- If an old module still loads, check for a duplicate development-folder copy and the connection's selected version.
- If authentication fails, check the listener's trusted source IP, key, API port and both clocks. After rotating the key, update the connection's secret field.
- If a preset sends keys to the wrong application, focus the intended window. Presets do not select the target application.

Developers building from source should use [DEVELOPMENT.md](DEVELOPMENT.md).
