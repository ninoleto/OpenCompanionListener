# Development handoff — Open Companion Listener

## 2026-11-12 FINAL: v0.1.4 published and verified; v0.1.3 retired (latest)

- **Branch/commits:** `fix/v0.1.4-module-package`; GitHub `main` received the tested history by fast-forward, including owner-PASS documentation commit `f3d6e6e50f1d5cdce2601cd1d01e9e8be65fd3fb`. Annotated `v0.1.4` points to the exact packaged source `30be74c8841594d67a4aaf49ee5cc4f6db36c827`. This final documentation-only record follows publication; application source and release bytes remain frozen.
- **Owner testing:** live Windows PASS for EXE, tray, dashboard and existing Companion buttons. The owner's final-release instruction superseded the earlier hold and expressly authorized old release/tag retirement only after public verification.
- **Published:** https://github.com/ninoleto/OpenCompanionListener/releases/tag/v0.1.4 — public, non-draft, non-prerelease, latest. Prepared notes finalized with the owner PASS, current validation link and frozen source/build-metadata explanation. All five existing files uploaded: Windows ZIP, Companion TGZ, source ZIP, BUILD-INFO.json and SHA256SUMS.txt.
- **Exact local release files:** `dist/release-v0.1.4/OpenCompanionListener-v0.1.4-Windows-x64.zip`; `community-openlistener-0.1.4.tgz`; `OpenCompanionListener-v0.1.4-Source.zip`; `BUILD-INFO.json`; `SHA256SUMS.txt` (all in that directory). Final EXE SHA-256 remains `b8f74decb94e9e60ec6b4015cb1df7a95e2cdd806add3f72de0ff5dd0d6847b5`; TGZ remains `949ab82e84ebf9cf0ccdaff3d811bc4486af41adfcf0a20c3f823928a299b725`.
- **Verification PASS:** existing five files match the prior reproduction byte-for-byte; all four recorded checksums, ZIP CRC/source manifests, EXE/TGZ identity pass. Public release page/API/body/latest state and annotated tag target verified. All five files independently downloaded without authentication, matching original bytes, sizes, GitHub digests and downloaded SHA256SUMS. No artifact was rebuilt or changed.
- **GitHub Actions:** automatic push runs `43274014863` (source/tag `30be74c`) and `43274014799` (validation `f3d6e6e`) completed successfully before retirement. No additional test/audit workflow was manually started. Existing local automated and owner results remain in the validation document.
- **v0.1.3 retirement:** only after those checks passed, deleted the v0.1.3 GitHub release and remote tag. Subsequent API checks confirm both absent; v0.1.4 release ID, asset identities/digests and annotated tag unchanged. Old release commit `0f988d4eff11b20203b98ad9fd5a35308ef55e9e` remains accessible and an ancestor of main. No history rewrite/force-push. Local old archives and local v0.1.3 tag remain as rollback records.
- **Files changed this final-release session:** HANDOFF.md, docs/VALIDATION-v0.1.4.md, GITHUB_RELEASE.txt, README.md, docs/COMPANION-INSTALL.md and CHANGELOG.md only. No code, working module/TGZ, artwork, installation, settings or startup state changed. Final evidence stays ignored under `.cache/v0.1.4/`; independent downloads are in `public-release-download-v0.1.4/` there.
- **Limitations/security:** owner PASS does not establish every OS build, preset, startup/sign-in, interruption or rollback case. No fresh antivirus scan or separate agent-driven desktop install was performed. Existing unsigned executable, unencrypted HTTP, desktop authority and replay limitations remain documented.
- **Rollback:** preserve Git history, v0.1.4 tag and published asset bytes. Use local previous EXE/module/configuration backups for deployment rollback; further public fixes require a new version.

**Single next safe action:** monitor v0.1.4 user reports and address any confirmed defect in a new patch release without replacing its tag or assets.


## 2026-11-12 final release authorized — owner live Windows PASS (historical)

- **Branch/commit:** `fix/v0.1.4-module-package`; packaged source remains `30be74c8841594d67a4aaf49ee5cc4f6db36c827`. This documentation-only commit records release authorization and will be pushed with the tested source. The v0.1.4 tag must identify the frozen artifact source; later documentation commits do not change the built files.
- **Owner PASS:** live Windows testing passed: the v0.1.4 EXE, tray, dashboard and existing Companion buttons work. This supersedes the previous publication hold and pending owner smoke confirmation. Detailed all-preset, startup, OS-matrix and interruption cases are not inferred from the report.
- **Authorized sequence:** verify existing artifacts/checksums, push tested history, publish v0.1.4 with all five existing files, verify the public release and independently downloaded bytes, then delete only the old v0.1.3 GitHub release and tag. Stop on failure; never delete v0.1.3 after failed publication/verification. Preserve all Git commit history and v0.1.4.
- **Files changed:** HANDOFF.md, docs/VALIDATION-v0.1.4.md and publication documentation only. No code/module/TGZ/logo/build output changed. Release notes distinguish frozen build-time records from current validation. No rebuild, new tests/audits or live-installation changes.
- **Artifact identity:** existing final EXE SHA-256 `b8f74decb94e9e60ec6b4015cb1df7a95e2cdd806add3f72de0ff5dd0d6847b5`; TGZ `949ab82e84ebf9cf0ccdaff3d811bc4486af41adfcf0a20c3f823928a299b725`. Full existing checksums remain in `dist/release-v0.1.4/SHA256SUMS.txt`.
- **Pending:** current checksum verification, remote push/publication/public-download verification and conditional old release/tag deletion. No success is claimed before evidence.
- **Risk/rollback:** artifacts remain unsigned and existing security limits apply. Local v0.1.3 archives and commit history remain preserved even after the authorized GitHub release/tag removal.

**Single next safe action:** verify the existing release bytes and checksums before any GitHub publication or deletion.


## 2026-11-12 v0.1.4 local release packages verified — publication held (historical)

- **Branch/commit:** `fix/v0.1.4-module-package`. Packaged application source is commit `30be74c8841594d67a4aaf49ee5cc4f6db36c827`. This later evidence-only update changes HANDOFF.md and docs/VALIDATION-v0.1.4.md, not the packaged application. BUILD-INFO.json identifies the frozen source exactly.
- **Outcome/files:** approved ICO and 128px PNG copied unchanged; existing resource generator produced the SYSO used by EXE/tray. No logo redesign or dashboard layout/color change. The existing module/TGZ and application behavior remain unchanged. Packaging reuses the verified TGZ with mandatory SHA-256 verification; no module rebuild, npm install or new environment.
- **Final packages:** `dist/release-v0.1.4/` contains exactly `OpenCompanionListener-v0.1.4-Windows-x64.zip`, `community-openlistener-0.1.4.tgz`, `OpenCompanionListener-v0.1.4-Source.zip`, `BUILD-INFO.json`, `SHA256SUMS.txt`. All five files match `dist/release-v0.1.4-reproduction/` byte-for-byte. A local testable EXE is extracted at `dist/release-v0.1.4-extracted/OpenCompanionListener.exe`.
- **Identity:** final EXE SHA-256 `b8f74decb94e9e60ec6b4015cb1df7a95e2cdd806add3f72de0ff5dd0d6847b5`. TGZ SHA-256 remains `949ab82e84ebf9cf0ccdaff3d811bc4486af41adfcf0a20c3f823928a299b725`. Final archives use git-archive LF HTML; the earlier direct working-tree test EXE used CRLF HTML and is not the release binary.
- **Tests/results:** two committed-source Windows vet/build/package runs PASS, exit 0, checksums identical. All four SHA256SUMS entries, ZIP CRCs, exact source inventory/hashes versus git archive, safe paths/private-file exclusions, notices, packaged manifest/runtime/module hashes and unchanged TGZ PASS. Final extracted x64/GUI EXE: all ten approved icon images plus group 1 match; exact PNG embedded; native Windows LoadImageW loads 16/32/48/256px icons without executing listener, PASS. Branded Go race suite 15 top-level + 20 subtests PASS; CGO=1 and CGO=0 vet PASS. Mocked headless dashboard screenshot visually checked; unchanged 44px logo slot/layout/colors PASS. Working module and original TGZ file fingerprints unchanged. Prior module 28/28 test result retained without rerunning/rebuilding it. `git diff --check` PASS.
- **Tooling exception:** the first archive attempt was denied access to its temporary source directory by the sandbox; authorized local rerun built both matching sets. No tools were downloaded, no prior audits repeated. Ignored validation helpers/previews remain under `.cache/v0.1.4/`.
- **Not performed/remaining:** actual visible tray and live EXE startup/dashboard/input smoke; running production listener owns port 17993 and was not stopped. No startup registry changes, fresh antivirus scan, exhaustive Windows/preset/rollback tests, remote CI, publication, public downloads or independent installation. Native resource loading and mocked preview are accurately scoped, not full desktop testing.
- **Security/rollback:** no new security behavior; preserve existing unsigned-EXE, HMAC-without-encryption, full desktop authority and memory-only replay caveats. Live listener/settings/Companion and published v0.1.3 remain untouched. Previous artwork is backed up and in Git. Restore matching previous EXE/module/configuration for deployment rollback.
- **Publication:** explicitly on hold at the owner's instruction. No push, tag or release action was performed.

**Single next safe action:** visually smoke-test the extracted final EXE in a separate Windows test account/VM with port 17993 free, keeping the production setup and publication hold intact.


## 2026-11-12 approved logo integration — source freeze for local packaging (historical)

- **Branch/commit:** `fix/v0.1.4-module-package`, based on `e9164b8`. This entry belongs to the source-freeze commit that will be identified by the final BUILD-INFO.json. Earlier uncommitted v0.1.4 work is preserved in the same checkpoint.
- **Objective/changes:** copied the supplied approved ICO to `OpenCompanionListener.ico` and 128px PNG to `listener/ui/icon.png`; regenerated `listener/icon_windows_amd64.syso` with the existing tool. Tray code already loads group 1, so no tray/Go behavior change was needed. Updated artwork notices, changelog, release draft and verification docs. Added paired `--module-package` / `--module-sha256` packaging arguments to preserve the verified TGZ without rebuilding/installing the module; metadata explicitly says local/unpublished and records reuse. The raw local design delivery remains ignored and preserved.
- **Compatibility/security:** no changes this session to module files, TGZ, dependencies, actions, transport, migrations, listener input/auth/network/startup or dashboard HTML/styles. The already-made version labels remain 0.1.4. Original artwork is backed up under ignored `.cache/v0.1.4/branding-backup/`. All supplied asset checksums pass. No new security boundary; earlier unencrypted-HTTP, desktop authority and unsigned-binary limits remain.
- **Tests:** Go race suite PASS (15 top-level tests plus 20 subtests), both CGO=1 / Windows-amd64 CGO=0 vet checks PASS, Windows GUI build PASS, all exit 0. All ten embedded PE images and tray group 1 exactly match the approved ICO; Windows LoadImageW loads resource 1 at 16/32/48/256px without starting the EXE. Approved PNG bytes are embedded unchanged. Headless dashboard preview PASS: approved 128px image in existing 44px slot; screenshot visually inspected using mocked data, no live dashboard access. Icon images visually inspected at 16/32/48/256px. Packaging rejects missing paired arguments and mismatched SHA (3 negative probes). Complete module file/TGZ hash baseline unchanged. Prior 28/28 module tests remain applicable and were not repeated.
- **Tooling:** reused installed tools and prior test caches; no downloads. Headless browser startup initially hit sandbox EPERM; authorized local rerun passed. Evidence and preview under ignored `.cache/v0.1.4/`.
- **Pending at source freeze:** two committed-source release builds, checksum/source/PE checks and exact TGZ equality. Subsequent evidence will identify this frozen source commit; do not confuse a later handoff-only commit with packaged application source.
- **Not performed:** actual listener startup/visible tray smoke because the working listener occupies port 17993 and must remain untouched; fresh antivirus scan, exhaustive desktop/preset/startup/rollback cases, GitHub CI/download/install. Resource loading and mocked rendering are not live tray or command execution tests.
- **Rollback:** previous artwork remains in Git and the ignored backup; v0.1.3 release and current production installations/settings remain untouched. Deployment rollback uses the matching previous EXE/module/configuration.
- **Publication:** explicitly prohibited by the owner in this turn. No remote write/tag/release action is authorized or performed.

**Single next safe action:** package this committed source twice with the pinned, verified TGZ and compare all release checksums.


## 2026-11-12 v0.1.4 continuation — logo and final release set pending (historical)

- **Branch/commit:** `fix/v0.1.4-module-package`, HEAD `e9164b8`. Existing uncommitted v0.1.4 work was inspected and preserved; this session also leaves its changes uncommitted. The incoming tracked diff is saved privately in ignored `.cache/v0.1.4/incoming.patch`. No push, tag, release edit or publication occurred.
- **Objective:** finish the already-working GUI-importable module, update installation/release documentation, integrate approved local artwork and verify v0.1.4 without changing the working Windows/Companion setup. Artwork files/path have not been supplied; a question is pending. Do not invent or regenerate a logo.
- **Owner evidence:** the v0.1.4 TGZ imported through Companion's GUI on the production host, retained connection settings, connected, and most existing buttons work. Recorded as owner-reported production testing; exact installed hashes, host versions, every action/preset and rollback are unspecified. Do not repeat completed testing.
- **Changes this session:** README, CHANGELOG, CONTRIBUTING, GITHUB_RELEASE.txt, SMOKE_TEST, THIRD_PARTY_NOTICES, docs/ARCHITECTURE, COMPANION-INSTALL, RELEASING, VALIDATION-v0.1.4 and ZORIN-VALIDATION; dashboard version labels; bundled module HELP version/GUI-install instructions; CI's packaging job now selects pinned Node 22.23.4. Developer commands stay in docs/DEVELOPMENT.md. Prior packaging, lockfile, manifest/version, license and validator changes remain intact.
- **Compatibility:** no module `src/` changes; actions, connection ID, presets and migrations preserved. Rebuilt TGZ differs from the previous local package only in HELP.md; compiled `main.js` is byte-identical (SHA-256 `e4ae10c1018d1f5bd2cda8432c1fbb72974453b353ff2dac8b37b6ed9d55ac06`). Dashboard HTML differs from HEAD only in two version labels; layout/colors remain identical. Existing ICO/PNG/SYSO have not changed. Listener code still differs only in its prior API version edit.
- **Tests (exact results):** Go 1.27.4 `go test -race -count=1 -v ./...` PASS, 15 top-level tests + 20 subtests, exit 0; `go vet ./...` CGO=1 and Windows/amd64 CGO=0 PASS, exit 0 each; Windows GUI build PASS, exit 0. Node 22.23.4 `npm test` 28/28 PASS, exit 0; real base manifest validation and official tools 2.9.1 build PASS, exit 0. TGZ entrypoint/identity/runtime/notices/payload checks PASS; five malformed archives rejected; normalized archive idempotence PASS. PE x64/GUI and embedded API/dashboard 0.1.4 verified. 34 relative documentation links resolve; `git diff --check` PASS. Existing security audits were not repeated.
- **Tooling issues resolved:** sandbox denied Node worker spawning; authorized local rerun passed. Go telemetry writes were redirected inside ignored workspace storage and disabled. PowerShell's stderr handling initially stopped the package build on the known upstream base-license warning; direct subprocess capture completed with exit 0 and the supplied notice is present. A local verifier initially compared normalized HTML with raw embedded CRLF; corrected probe passes without changing HTML.
- **Verified artifact:** `dist/module-v0.1.4-verified/community-openlistener-0.1.4.tgz`, SHA-256 `949ab82e84ebf9cf0ccdaff3d811bc4486af41adfcf0a20c3f823928a299b725`; adjacent SHA256SUMS covers this TGZ only. Existing pre-help package retained at `.cache/v0.1.4/module-before-help-update.tgz`. Pre-logo EXE at `.cache/v0.1.4/OpenCompanionListener-docs.exe` is a test build, not a release artifact. Detailed checks/logs remain under `.cache/v0.1.4/`.
- **Verification not performed / blockers:** approved logo integration; final clean committed-source ZIPs/BUILD-INFO/full checksums and repeat-build comparison; final logo visibility in EXE/tray/dashboard; final release-commit CI and independent downloads/installation. The running listener keeps dashboard port occupied; no native smoke run or shutdown was attempted. Full all-preset, startup, interruption and rollback coverage remains unverified. Future v0.1.4 links are explicitly labelled forthcoming, not claimed live. No new antivirus scan is claimed.
- **Security/risks:** no new execution surface or protocol change. Existing unsigned-EXE, unencrypted-HTTP/full-desktop authority and memory-only replay limitations persist. Production listener, settings, startup registry and Companion installation untouched. Published v0.1.3 preserved.
- **Rollback:** retain the incoming diff, original artwork and earlier artifacts; revert only this session's focused edits if required. Deployment rollback remains the previous EXE plus matching module/configuration backup.

**Single next safe action:** obtain the approved local logo package/path, then integrate its supplied ICO/PNG via the existing resource pipeline before building and verifying the final release set.


## 2026-11-11 public description update — published and verified (historical)

- **Branch/commit:** documentation branch `docs/v0.1.3-description` from public main `0f988d4eff11b20203b98ad9fd5a35308ef55e9e`. Description commit `eff5b1e8ec72cd49d756aa9f6440c2acce3eec8a` was pushed to main; this handoff-only follow-up records the completed verification. Existing development/audit branches are preserved.
- **Objective/outcome:** describe the single primary purpose clearly: send keystrokes, shortcuts, mouse clicks and scrolling commands to a Windows PC using Bitfocus Companion. Published the owner's exact opening in the complete rewritten v0.1.3 release notes and README introduction; set the repository About text exactly as requested. Removed industry-specific positioning and promotional workflow claims from the release description.
- **Files changed:** README.md introduction and GITHUB_RELEASE.txt (Markdown content); HANDOFF.md for this record. The README from `## Features` onward is unchanged. GitHub CLI updated the release with an identical ignored Markdown file at `.cache/v0.1.3-release-notes.md`. No application, module, tooling, CI, version, binary or asset file changed.
- **Documentation coverage:** controls, 38 opt-in presets, listener/module communication, all five actual asset links, Windows/Companion installation, Node 22.8+ within 22.x, exact npm commands, trusted-network/security limits, MIT/source/Actions/contribution guidance, AGENTS.md, HANDOFF.md and Issues. Ordinary user instructions precede developer and AI-agent guidance.
- **Checks completed:** `git diff --check` PASS, exit 0. Exact supplied opening and unchanged README technical remainder verified. GitHub's Markdown renderer produced the expected headings, lists, table and code blocks for both documents. All 15 release-note links returned HTTP 200. Unauthenticated public release/README content matches the local documents; actual GitHub HTML contains the correctly rendered opening, sections and install commands. About text matches exactly. An initial local verifier incorrectly required two ordered lists in the unchanged README; correcting that verifier assumption yielded PASS without changing the README's technical content.
- **Release preservation:** annotated `v0.1.3` tag object remains `fde0d75e4bd9eb569f40f576571ff0655a145b37`, pointing to release commit `0f988d4eff11b20203b98ad9fd5a35308ef55e9e`. All five existing asset IDs, names, sizes, SHA-256 digests and creation/update timestamps match the pre-edit baseline. Release ID, title, tag, publication time, target and draft/prerelease state are unchanged. Only its description was edited. Main advances for documentation; the versioned source archives and binaries are preserved.
- **Verification not performed:** no application tests, rebuilds, fresh antivirus scans, binary uploads or interactive desktop checks for this documentation-only change. Existing v0.1.3 test/coverage limits remain stated accurately; no new functionality or test result is claimed.
- **Security/risks/unresolved bugs:** no new execution behavior or security boundary; pairing-key, firewall and unencrypted-HTTP guidance remains. No documentation blocker found. Private baseline/render/link evidence is ignored under `.cache/` and is not committed.
- **Rollback:** revert only the documentation commits on main and restore the earlier About/release text from the local pre-edit snapshot if needed. Never move the v0.1.3 tag or replace its assets for a description correction.

Public pages: https://github.com/ninoleto/OpenCompanionListener and
https://github.com/ninoleto/OpenCompanionListener/releases/tag/v0.1.3.

**Single next safe action:** review incoming GitHub Issues for documentation
corrections while keeping the v0.1.3 release artifacts unchanged.

## 2026-11-11 first public release preparation — repository created (historical)

The owner explicitly authorizes repository creation, source push, annotated tag
`v0.1.3`, and the first public release with verified assets. This supersedes the
earlier instruction to stop at a repository 404. Documented manual testing limits
remain limits, not passed checks; the owner approves publication of this exact
corrected EXE without additional application changes.

- **Branch/commit:** `fix/v0.1.3-publication`, preparation based on `f11916317b741116ee8d52aece70f1eaf00e68bc`. The release commit is the committed source containing this entry; final BUILD-INFO.json records its full SHA. Existing candidates from `58b9f14` remain preserved.
- **GitHub:** installed official GitHub CLI 2.103.3 with winget, which verified its installer hash. `gh auth status` and `gh api user` confirmed authenticated account `ninoleto`. The authenticated target lookup returned 404; created public `ninoleto/OpenCompanionListener` with the requested description and Issues enabled, no generated README/license/gitignore. Verified empty repository and connected origin. No unrelated repository was changed.
- **Objective/scope:** publish frozen corrected v0.1.3. Pending artwork/manual/Defender documentation is retained. README/changelog/install guide and release notes describe the first public release and actual testing limits. Module manifest gains only repository/bugs URLs. Packaging adds an opt-in `--release` switch for final filenames and BUILD-INFO status; build flags/content selection and default candidate mode are unchanged. No listener code, UI, action implementation, migration, transport, dependency or artwork changes.
- **Fresh automated verification:** Windows Go 1.27.4 `go test -race -count=1 -v ./...` PASS (15 top-level tests plus 20 key subtests), exit 0. `go vet ./...` with CGO=1 and Windows/amd64 CGO=0 PASS, exit 0 each. Node 22.23.4 clean `npm ci --omit=dev --ignore-scripts --no-audit --no-fund` installed 17 packages, exit 0; installed base 1.13.6 manifest validation PASS; 28/28 module tests PASS, exit 0; npm audit zero advisories, exit 0. Fresh govulncheck source and corrected-binary checks report no vulnerabilities, exit 0 each.
- **Privacy/source review:** tracked-file inventory and 96 historical blobs scanned for common credential patterns, personal Windows/Linux paths and private IPv4 examples: zero findings; private/generated paths excluded. This targeted scan is not proof against all possible secrets. `git diff --check` PASS. Listener and module `src/` diffs are empty. Raw evidence and helper scripts stay in ignored `.cache/`.
- **Existing executable evidence:** required SHA-256 is `510696b6fba8331bfef6d929013213fea15f43f3dde78f7db34c8370dac97f83`. Previous isolated smoke and completed Defender no-threat scan apply only if final packaging reproduces this exact EXE. Owner-reported functions and ZORIN installation are recorded in SMOKE_TEST.md; untested detailed cases are not marked passed.
- **Files changed for publication:** CHANGELOG.md, GITHUB_RELEASE.txt, HANDOFF.md, README.md, SMOKE_TEST.md, THIRD_PARTY_NOTICES.md, docs/AUDIT-v0.1.3.md, docs/COMPANION-INSTALL.md, docs/RELEASING.md, module companion/manifest.json and tools/package_release.py. Existing historical handoff/audit evidence is retained below.
- **Still to verify at this checkpoint:** final committed-source builds/reproduction and extracted module checks; main/tag push, public release/upload and independent downloads. No new interactive Windows/ZORIN tests or independent desktop installation have been performed. A live listener is not stopped or reconfigured for publication.
- **Risks/rollback:** unsigned EXE; HMAC is not encryption; full desktop-user authority; memory-only replay cache across process exits; detailed manual coverage remains incomplete. Preserve original deployments/candidates. Revert only publication metadata/tooling if necessary before release; after publishing use a new version rather than replacing tags/assets. Restore matching old EXE/module/config export for deployment rollback.

**Single next safe action:** build the final committed source twice, require the
specified EXE hash and archive checks, then publish the verified commit/assets
under the owner's existing authorization.

## 2026-11-11 continued publication preparation — GitHub access blocked (historical)

The owner approves the first public v0.1.3 release and confirms project-generated
artwork and partial working-function tests. This section supersedes earlier
claims of missing publication permission, unknown artwork rights and missing
successful antivirus scanning. No additional publication permission is needed.

- **Branch/commit:** `fix/v0.1.3-publication`, HEAD `f11916317b741116ee8d52aece70f1eaf00e68bc`. Entry status contained the previous session's uncommitted HANDOFF.md only. This session leaves documentation edits for review; no commit, remote, tag or release was created. Existing packages identify source `58b9f14821326cae0f6d92a9f1a379babff14c69`.
- **Objective:** resolve supported artwork/manual/antivirus evidence, verify package provenance and GitHub access, and publish only if the remaining requirements are satisfied. Application behavior remains frozen.
- **Artwork:** owner confirms programmatically drawn project-specific graphics, not copied third-party artwork, and directs MIT inclusion. ICO/PNG/resource hashes match original snapshot/source ZIP, checkout and candidate archives. Dashboard PNG equals the ICO's 128px entry; all seven embedded EXE icon images and group match. Isolated resource regeneration is byte-identical. No additional tracked artwork was found. The original drawing script was not located; authorship provenance is owner-attested, not independently reconstructed.
- **Manual evidence:** owner reports working Num Lock, keyboard/mouse controls, DisplaySwitch.exe arguments 1–4, AutoHotkey launching, Voicemeeter recorder/VBAN feedback workflow, and corrected ZORIN module installation with 28 passing tests. Recorded as owner reports without checking unobserved detailed test cases. Current running listener SHA-256 matches the corrected candidate and prior smoke EXE (`510696b6fba8331bfef6d929013213fea15f43f3dde78f7db34c8370dac97f83`). The preserved original EXE differs. Historical manual-test executable identity was requested and remains unconfirmed; no live process/settings/module was changed.
- **Defender:** `MpCmdRun -Scan -ScanType 3 -File <corrected-exe> -DisableRemediation -ReturnHR` completed, reported no threats, exit/HRESULT 0; no skipped-file message, EXE unchanged. Product `4.18.26100.4`, engine `1.1.26100.4`, signatures `1.463.243.0`, updated 2026-11-11 09:41:12 UTC and not out of date. A sandbox-denied status read was successfully repeated with authorized read-only access; the successful scan was not repeated. Point-in-time EXE result only, not a safety guarantee or dependency scan.
- **Package checks:** local provenance verifier PASS, exit 0: four checksum entries; CRC/safe paths/source contents in all three ZIPs; all 62 source hashes against recorded-commit `git archive`; no private settings, dependencies or caches in checked archive paths; EXE identity, artwork and PE resource checks. The probe first failed when comparing raw Git blobs: six text files undergo CRLF conversion in Git archive under local configuration. The final checks explicitly verify those differences are only line endings and compare the actual packaging input. Corrected the documentation's canonical-LF claim; cross-host byte identity remains unverified.
- **GitHub blocker:** after the sandbox network failure, the permitted read-only Git retry reached GitHub and returned `Repository not found` (exit 128); unauthenticated repository API returned HTTP 404. The target is absent or inaccessible with available access. Per the owner's instruction, stopped before repository creation, push, tag or release/upload. Target remains `https://github.com/ninoleto/OpenCompanionListener`.
- **Files changed:** THIRD_PARTY_NOTICES.md, HANDOFF.md, docs/AUDIT-v0.1.3.md, SMOKE_TEST.md, docs/RELEASING.md, README.md and GITHUB_RELEASE.txt. Ignored verification script, extracted icon images and scan/provenance evidence remain under `.cache/`; they are not public source. No application/module/build-tool code, artwork or archives changed.
- **Documentation verification:** `git diff --check` PASS, exit 0. `git diff --name-only -- listener companion-dev-modules tools` returns no changed files, exit 0. The working tree contains only the seven documentation changes listed above, including the previous handoff entry.
- **Verification not performed:** fresh Go/Node suites or build (documentation-only session); independent ZORIN inspection; manual tests by the agent; full tray/startup, migration/rollback, all-preset, hold/interruption or Windows configuration coverage; new packages; remote CI; independent download/install or release asset comparison. Existing automated and isolated smoke results remain historical evidence for the unchanged code/EXE. Current packages omit these latest documentation edits and require final-commit repackaging before release.
- **Security risks/unresolved bugs:** existing SECURITY.md limitations remain; no new behavior or defect fix introduced. Repository access is the immediate blocker. Manual-test-time executable attribution and detailed untested cases remain explicitly unverified. No published commit, release URL or uploaded assets resulted from this session.
- **Rollback:** nothing deployed or published; original and candidate binaries/archives are intact. Documentation edits can be reverted independently; preserve the prior handoff entry when reviewing this session's diff.

**Single next safe action:** make `ninoleto/OpenCompanionListener` accessible:
the owner should create the repository if absent, or provide authenticated
access if it already exists privately. Recheck access before any publication.

## 2026-11-11 publication request — stopped at unresolved release gates (historical)

This section supersedes earlier statements that the owner withholds publication
permission. The owner explicitly authorizes the first public v0.1.3 release at
`https://github.com/ninoleto/OpenCompanionListener`, with a release freeze and an
instruction to stop and report any release-blocking discrepancy. That approval
remains valid; the stop below concerns missing release evidence, not permission.

- **Branch/commit:** `fix/v0.1.3-publication`, inspected HEAD `f11916317b741116ee8d52aece70f1eaf00e68bc`. Working tree was clean at entry. No Git remote was configured. This session adds only this uncommitted handoff entry.
- **Objective:** verify the corrected source and existing candidates before publishing the approved first public release, without functionality changes.
- **Review:** read AGENTS.md, this handoff, architecture, security, smoke checklist, release workflow, audit, third-party notices and release draft; reviewed recent commits and tracked-file inventory.
- **Blocking discrepancy:** artwork authorship and redistribution permission are still explicitly unconfirmed in THIRD_PARTY_NOTICES.md and audit finding A08. The release workflow says to keep publication blocked when those rights are unknown. The audit also still lacks the required interactive Windows/Companion results and a successful antivirus result; the earlier Defender attempt skipped the file. Publication authorization does not supply those missing facts.
- **Checks this session:** `git status --porcelain=v1` was empty at entry (exit 0). Recomputed SHA-256 for all four entries in `dist/publication-candidate/SHA256SUMS.txt`: BUILD-INFO.json and the Windows, Companion module and source ZIPs all match (exit 0). BUILD-INFO.json identifies source `58b9f14821326cae0f6d92a9f1a379babff14c69`. `git diff --name-only 58b9f14821326cae0f6d92a9f1a379babff14c69 HEAD` shows only six documentation files; no application/module changes since packaging. The existing packages therefore omit the latest validation documentation and would need repackaging from the final publication commit after the blockers are resolved.
- **Files changed:** HANDOFF.md only. No application, module, dependency, package, settings, startup entry or live installation changed.
- **Not performed:** fresh Go/Node tests, rebuild/repackaging, fresh antivirus scan, interactive Windows/Companion checks, GitHub authentication/remote inspection, source push, tag, release creation, upload, independent download/install or remote asset comparison. Earlier automated and limited executable smoke results remain historical evidence.
- **Security risks/unresolved issues:** artwork redistribution rights, manual validation and antivirus evidence remain unresolved. No new code/security behavior was introduced; existing security limitations remain as documented in SECURITY.md. No published commit, release URL or uploaded assets resulted from this session.
- **Rollback:** no release or deployment to revert; existing candidate archives and source commits are untouched. This handoff-only diff can be reverted independently.

**Single next safe action:** obtain and record the bundled artwork's origin and
redistribution permission before resuming release validation. Do not infer
authorship, replace artwork or change application behavior during the freeze.

## 2026-11-11 final validation — limited executable smoke PASS

This section is the latest status. The owner requires preservation of the working
Windows installation, ZORIN module, settings and paired devices, and explicitly
withholds permission to publish. The candidate application/module is unchanged.

- **Branch/commit:** `fix/v0.1.3-publication`, candidate source `58b9f14821326cae0f6d92a9f1a379babff14c69`; runbook preparation `290ba61`, followed by this results-only documentation commit. The candidate archives were not rebuilt or replaced.
- **Objective:** identify the port 17993 owner without stopping it, guide the owner through a tray exit, run the existing isolated executable smoke tool afterward, and prepare ZORIN/manual validation procedures.
- **Confirmed:** read-only inspection identified the existing installed `OpenCompanionListener.exe` as the owner of port 17993 connections. The owner then confirmed "Exited" after receiving tray Exit instructions. A fresh check found no listener process; only closing/time-wait TCP entries remained. The agent did not terminate or mutate the working listener. The candidate smoke tool successfully bound its own dashboard and API afterward.
- **Tests/checks this session:** all existing candidate SHA-256 checksums match; the executable selected for the smoke tool is byte-identical to the EXE in the candidate Windows ZIP and matches BUILD-INFO.json (`510696b6fba8331bfef6d929013213fea15f43f3dde78f7db34c8370dac97f83`). ZORIN runbook Bash examples pass `bash --noprofile --norc -n` (exit 0, syntax only, no commands executed). `git diff --check` passes. Sandbox restrictions on Windows process metadata and Git Bash startup were resolved with permitted read-only/syntax-check reruns.
- **Executable smoke:** Node 22.23.4 `node tools/smoke_listener.cjs .cache/evidence/OpenCompanionListener.exe` PASS, exit 0, through the local preservation-check wrapper. Confirmed candidate startup; dashboard GET `/` HTTP 200; missing UI header/foreign Host state reads HTTP 403; signed Node-to-Go `/v1/status` success; consumed nonce rejected HTTP 401 after settings save and in-process restart; fresh requests continue succeeding; duplicate candidate exits 0 while the first remains; test-key rotation persists, rejects old-key requests and accepts the new key; `/ui/exit` ends the candidate with exit 0. No tray initialization error was observed by the script; visible icon/menu behavior was not inspected.
- **Preservation/cleanup:** before/after SHA-256 comparisons confirm the original installed EXE and current-user settings file are unchanged; the existing HKCU startup value is unchanged. After the test there are no OpenCompanionListener processes and no listening socket on port 17993. The original application remains stopped after the owner's exit; it was not automatically relaunched. Raw smoke output is local-only at `.cache/evidence/final-executable-smoke.txt`.
- **Not performed:** ZORIN access/install/configuration import, desktop input or visual tray tests, startup registry writes, sign-in tests, real-host migrations, paired-device changes, antivirus rescan or publication. Previous Go/Node results below remain historical verification; they were not rerun for documentation-only edits.
- **Files changed:** preparation commit `290ba61` added the ZORIN runbook and expanded manual checklist. This follow-up updates `HANDOFF.md`, `docs/AUDIT-v0.1.3.md`, `SMOKE_TEST.md` and `docs/RELEASING.md` with the confirmed executable result. No application, module or test-tool code changed.
- **Isolation/rollback:** unchanged smoke tool used temporary LOCALAPPDATA, a random test key and a loopback-only ephemeral API port. The ignored preservation wrapper ran successfully. No production files, registry values or ZORIN modules changed. Existing candidate archives and rollback snapshots remain intact. The working installation can be resumed by launching its original EXE; no data restoration or re-pairing is needed for this test.
- **Remaining gates:** Windows/Companion manual checklist, artwork rights, successful antivirus verification, owner distribution/version/repository decisions, and explicit publication approval. LOCALAPPDATA isolation does not isolate HKCU startup: manual startup tests require a separate Windows account/VM.

**Single next safe action:** perform the remaining manual Windows/Companion checklist in an isolated test setup using `SMOKE_TEST.md` and `docs/ZORIN-VALIDATION.md`, preserving the installed module and production device pairings.

## 2026-11-11 publication-fix continuation — local candidates only

This section supersedes the audit-only and historical next actions below.
Publication remains **on HOLD**; the owner explicitly prohibited GitHub publishing.

- **Branch/commit:** `fix/v0.1.3-publication`; listener fixes `5d78b82`, Companion fixes `3904614`, packaging/documentation checkpoint `a9b6ae3`, followed by this evidence/tooling commit. Audit parent `5faf3f1`, supplied source baseline `cd2f043`. Exact final candidate commit is recorded in `dist/publication-candidate/BUILD-INFO.json`; no upstream remote, tag or release was created.
- **Objective:** retain completed security/migration/transport fixes, finish dependency and packaging work, verify clean builds and document release gates without changing live deployments or preserved snapshots.
- **Files changed:** listener config/main/dashboard and security regressions; module main/transport/upgrades/tests, pinned package/lock and real-base manifest validator; deterministic archive builder/toolchain pins, CI, Git attributes/ignores, Go license/third-party notices, README/security/architecture/install/release/smoke/changelog/release draft and audit/handoff docs.
- **Automated tests:** Windows Go 1.27.4 `go test -race -count=1 -v ./...`: 15 top-level tests plus 20 key subtests PASS, exit 0. `go vet ./...` with Windows/amd64 CGO=1 and CGO=0: PASS, exit 0 each. Node 22.23.4 clean `npm ci --omit=dev --ignore-scripts --no-audit --no-fund`: 17 dependencies installed, exit 0; `npm run validate`: real base 1.13.6 manifest PASS, exit 0; `npm test`: 28/28 PASS, exit 0. Repeated clean install/validation/all 28 tests from the extracted module archive: PASS. npm audit: zero advisories, exit 0. govulncheck v1.9.0 source and binary scans: no vulnerabilities, exit 0 each (database updated 2026-11-05). `git diff --check`: PASS.
- **Packages:** two complete builds from checkpoint `a9b6ae3` were byte-identical across separate temporary source directories. SHA-256, ZIP CRCs, required files, per-source-file hashes and Go license match PASS. The 6,990,848-byte EXE is Windows x64 GUI, Go 1.27.4, CGO=0, GOAMD64=v1; SHA-256 `510696b6fba8331bfef6d929013213fea15f43f3dde78f7db34c8370dac97f83`. Final documentation/tooling is packaged again from this commit into `dist/publication-candidate/` and compared with `dist/publication-reproduction/`; each directory's SHA256SUMS.txt is authoritative for its archives. Module ZIP contains source/lockfile and requires the documented npm install step. No dependency bundle or in-app upload format is claimed.
- **Compatibility:** wire format, pairing keys, ports, connection ID, action IDs and all 38 presets retained. Upgrade index 0 is retained with expression-loss prevention; index 1 repairs stranded scroll actions. Expressions already destroyed by an old migration need a pre-upgrade export. Dashboard uses only the canonical numeric loopback URL. Existing app launch behavior is preserved and documented as broad user-session authority.
- **Not performed:** live Companion install/import-export, Windows desktop input, tray observation, Explorer restart, sign-in autostart/registry mutation, Windows 10/11 matrix, independent cross-host reproduction/installation, signing, remote CI or publication. The new limited executable smoke tool refused to start (`EADDRINUSE`, exit 1) because the existing dashboard occupied `127.0.0.1:17993`; it did not read or mutate the live dashboard, launch the test EXE or change startup. Defender printed **file was skipped**, exit 0: no successful scan or clean result. Initial Node worker/Python temporary-directory sandbox denials were resolved with authorized local reruns; no test failure was hidden.
- **Security risks/unresolved:** HMAC is not encryption; executable arguments allow interpreters; replay cache is memory-only across full process exits; in-flight requests can finish under their captured settings. Artwork rights remain unconfirmed; the owner questions received no answer during this session. Actual repository/security URLs and prior-publication/version status remain owner decisions. If v0.1.3 was publicly distributed, fixes require a new patch version. A successful antivirus scan and actual Windows/Companion smoke evidence remain release gates. A targeted public-source scan found no personal home/project paths, private-key blocks or GitHub token patterns; this is not proof of absence of all secrets.
- **Rollback:** original v0.1.3 archives/snapshots remain untouched outside this checkout. New builds go under ignored `dist/` directories and refuse overwrite. No real settings, registry values, installed module or Companion configuration was changed. For future deployments restore both old module and matching exported configuration.

**Single next safe action:** run the final candidate through `SMOKE_TEST.md` on an isolated Windows/Companion test setup where port 17993 is free, preserving the publication hold and rollback snapshots.

## 2026-11-11 audit handoff — publication on HOLD

This section supersedes the release-readiness claims and next actions in the historical handoff below. The owner requested findings before corrective implementation and prohibited publication. The source audit is complete; the candidate is **not ready for public release**.

- **Branch/commit:** local branch `audit/v0.1.3`; supplied source baseline `cd2f0430d968794007f89facb64cb62f6f08c855`. The incoming workspace had no Git history. This checkout was extracted from the supplied GitHub source ZIP, whose SHA-256 is recorded in [the audit report](docs/AUDIT-v0.1.3.md).
- **Objective:** audit completeness, security, privacy, licensing, buildability, tests, public documentation, and release readiness while preserving working behavior.
- **Changes:** added `docs/AUDIT-v0.1.3.md` and this handoff section. Generated a local untracked module `package-lock.json` during dependency verification; it is evidence, not an approved dependency change. No application/module implementation changed. Parent audit directory contains isolated tools, logs, executable, and defect probes; workspace-root HANDOFF points here.
- **Verification:** Windows/amd64 Go 1.27.4 `go test -race -v ./...` passed (7 top-level tests, 20 key subtests; exit 0). `go vet ./...` passed with both CGO=1 and CGO=0 (exit 0 each). Required GUI Windows build passed (exit 0). Node 22.23.4 original tests passed 16/16. Dependency install, installed-base manifest validation, and npm pack dry run passed. `npm audit` reported zero advisories. Icon-resource regeneration was byte-identical. Existing source/Windows ZIP and inner checksums match; advertised module ZIP is missing.
- **Defect evidence:** added probes fail 3 Go, 3 migration, and 1 transport assertions. Confirmed foreign-Host pairing-key disclosure, replay acceptance after unchanged settings save, active-key change on persistence failure, broken legacy scroll migration/expression metadata loss, and pending transport promises after response abort. See report A01–A06 for precise evidence and fixes.
- **Security/licensing risks:** original EXE embeds Go 1.23.2; binary vulnerability scan reports 67 standard-library advisories (not 67 proven exploits). Fresh Go 1.27.4 source/binary scans report none. The file-launch authority is broader than the README's no-shell claim. Go binary redistribution notices are missing; artwork provenance remains unconfirmed. Public source scan found no credential/personal-path matches; surrounding historical workspace has private deployment material and must not be uploaded wholesale.
- **Not verified:** actual Windows input, tray, sign-in startup, launch completion, real Companion host migration, all presets, and Windows 10/11 matrix. No live user configuration or registry writes performed. Defender scan failed to run (`0x80508023`, exit 2); no clean antivirus claim. Original and rebuilt EXEs are unsigned. No historical upstream commit, release tag, public GitHub URL, or byte-identical reproduction of the original EXE was established.
- **Unresolved:** audit A01–A10, including dependency/engine drift, missing module archive, notices/provenance, public install/build instructions, and manual smoke evidence. No application fixes have been applied because findings are being reported first.
- **Rollback:** original snapshots, EXE, archives, and checksums are untouched. Audit artifacts are isolated and unapproved; no live deployment or publication to revert.

**Single next safe action:** after reviewing the audit, implement and regression-test A01 dashboard authority validation on this branch. Keep publication blocked until the other findings and manual release gates are resolved.

## Historical handoff (retained for context; not current release approval)

**Last verified source baseline:** `v0.1.3`, released package built 2026-11-10.  
**Scope:** Windows 10/11 x64 listener + Bitfocus Companion `community-openlistener` module.  
**Project purpose:** A free, independent, lightweight desktop-control system. Public contributions and forks are explicitly welcomed (MIT).

## What already works and must stay working

- Windows-native Go listener with silent system-tray startup, Open GUI / Restart Listener / Exit, local settings dashboard, opt-in **Start with Windows** (HKCU Run), and single-instance behavior anchored by local dashboard port.
- Go HTTP API: configurable trusted peer IP and port 17992; 64-character hex pairing key; signed commands, nonce replay protection and clock-skew checking. Local UI on 127.0.0.1:17993.
- Windows input: single keys including F1–F24 and numpad, shortcut press/hold/release, mouse movement/buttons/hold/release, vertical/horizontal wheel scrolling, Num Lock toggle/on/off.
- Direct Windows Shell launching: HTTP/HTTPS URLs, application executables, files/folders, Windows Settings URIs and Control Panel tools. No generic remote shell action.
- Companion API 1.x CommonJS module, checked into `companion-dev-modules/companion-module-community-openlistener/` with `@companion-module/base` dependency. 38 curated presets and GUI-oriented action builders.
- Action migrations (older scroll and keyboard IDs) via `src/upgrades.cjs`. **Never remove published migrations**; migration behavior must be verified on a backed-up real Companion configuration, not just unit-tested.

## Verification: distinguish tested from untested

- **Automated:** On 2026-11-10, Linux-host Go tests passed; Windows cross-target static checks passed; Companion Node tests passed (16 total). Go race checks were reported as passed in the build session that produced v0.1.3; rerun on any changed code.
- **User testing:** Prior Windows-native input and application-launch functions from the earlier 0.1.x releases were exercised on an actual PC. The user subsequently installed v0.1.3 and reported it looked good. That is useful feedback, **not evidence that every 0.1.3 feature passed the full smoke test**.
- **Still requires explicit release verification:** Windows sign-in autostart registry toggle, silent launch and duplicate instance prevention; real installed Companion upgrade behavior including legacy button migration; all 38 presets; corner cases for modifier hold/release and interruption.

## Architecture and files

- `listener/main.go` — UI HTTP server, authenticated API, lifecycle, key rotation, startup setting handlers.
- `listener/config.go` — local user settings and pairing key storage.
- `listener/auth.go` — signature checks, timestamp + nonce replay window.
- `listener/commands.go` — accepted command schema and bounds.
- `listener/desktop_windows.go`, `listener/keys.go` — native input and Shell launching; non-Windows stubs are NOT a Linux implementation.
- `listener/tray_windows.go`, `listener/startup_windows.go` — native Windows user session lifecycle.
- `listener/ui/index.html` — browser-rendered loopback dashboard.
- `companion-dev-modules/companion-module-community-openlistener/src/main.cjs` — Companion actions/config, status pings.
- `src/transport.cjs` — canonical request signing; `src/upgrades.cjs` — saved-action migrations; `src/presets.cjs` — 38 optional presets.

Detailed flow, command schema and threat model: `docs/ARCHITECTURE.md` and `SECURITY.md`.

## Known limitations and engineering risks

1. **HMAC is not confidentiality.** Requests use HTTP without TLS; source IP + HMAC do not make Internet exposure safe. User must firewall to the specific trusted Companion host and use LAN/private overlay.
2. Pairing key is stored in a per-user JSON file. Do not commit real config, print keys in logs, or misrepresent it as Windows credential-vault protected.
3. GUI automation follows Windows focus and integrity restrictions. Elevated/locked desktops cannot be promised as supported.
4. No Windows service, no background input when logged out. Autostart takes effect at **user sign-in** only.
5. `key_down`/`key_up` may need extra fail-safe handling for interrupted connections and unreleased modifiers; verify and design targeted tests rather than assuming automatic release.
6. Migration in `src/upgrades.cjs`: non-integer or expression-based legacy scroll values may stay unchanged; test against real Companion exported configurations and decide explicit recovery steps.
7. Existing `example.org/open-companion-listener` Go module path is a neutral placeholder; update only as part of a reviewed repository-bootstrap change and rerun builds. Module manifest `repository`/`bugs` are currently empty—populate them **after the real GitHub URL exists**.
8. **Linux:** No usable Linux input backend. Build tags allow non-Windows compilation only. Wayland/X11 require researched, consent-aware backend choices and explicit UI privilege constraints.
9. No built-in AHK builder yet. The roadmap considers optional external automation with clear permission boundaries, not unrestricted remote shell execution.

## Working agreements for future forks / AI agents

- Keep `v0.1.3` as a rollback checkpoint. Do not overwrite tags, rewrite history, or publish an untested EXE in place of a stable download.
- Change one subsystem per PR (listener core, GUI, Companion actions, migration, or docs). Keep action IDs + signed transport compatible, or create and test a migration.
- Use the exact test commands in `AGENTS.md`; gather **Windows host** and **Companion host** test logs before claiming interoperability.
- Back up installed Companion module and Companion settings before upgrades; don't run test module changes against active streaming sessions.
- For a bug report include version, OS build, Companion version, steps, expected/actual behavior, and **redacted** logs. Never share pairing key.

## Suggested next work, ordered

1. Have the maintainer complete `SMOKE_TEST.md` on v0.1.3, with special attention to autostart and action migration.
2. Create the first public GitHub repository (no public repository URL was present when this handoff was authored), populate `repository` and `bugs` manifest values with that actual URL in a follow-up patch, and publish v0.1.3 artifacts with SHA-256 checksums.
3. Triage any user-reported bugs and issue **new patch tags** (`v0.1.4`, etc.). Do not break existing Companion buttons.
4. Evaluate optional automation workflows/AHK builder and future Linux support as separately scoped follow-up milestones.

**Next safe action:** finish v0.1.3 on-device smoke tests; upload the reviewed source to an owner-created public GitHub repository and attach tested Windows/module ZIPs as release assets. If a smoke test fails, keep the public release notes accurate and fix in a new patch version.
