# Changelog

---

## [3.0.0] - 2026-10-19

This release also includes everything from 2.2.0 and 2.2.1 (listed below), which were never published separately — including the sync codes, the PIN and import hardening, and the removal of auto-login.

### Added
- **New settings panel** — five tabs (Appearance, Dashboard, Security, Sync, About) instead of one long scroll. Controls are grouped into rows with a short description where it helps
- **Accent colour** — pick from six colours; it recolours buttons, toggles, highlights and tabs, and travels with sync codes and backups
- **Auto theme** — Light, Dark or Auto, where Auto follows your device and switches live
- **Reorder from settings** and a **Forget this device** button for trusted-device sign-in
- Rename and delete categories with icon buttons, and an About tab with a link to the project
- **Icon cache** — service icons are downloaded once, saved in your browser and reused, so unlocking and reloading show them instantly (and offline). Icons are warmed up before you unlock, icons known to be missing aren't re-requested every time, and a **Refresh icons** button clears the cache. If downloads are blocked, Pagerr falls back to loading icons the old way

### Changed
- **Emoji interface icons replaced with a consistent SVG icon set** — header, add button, settings, backup, lock screen and empty states. Service icons are unchanged
- Settings labels use sentence case, and wording is plainer (for example "Create sync code", "Require PIN")
- Visible keyboard focus, Escape closes settings, and animations are reduced for people who ask for that

---

## [2.2.1] - 2026-10-19

### Security
- **Imports are sanitised** — sync codes, JSON imports and stored data are validated before use. Only `http://` and `https://` service URLs are accepted (blocks `javascript:` URLs), IDs are restricted to safe characters, text fields are length-limited and stripped of control characters, and imported logos must be real image data URIs. Previously imported data was trusted as-is
- **PIN never stored in plain text on HTTP** — a built-in PBKDF2-HMAC-SHA256 implementation (identical output to Web Crypto, verified against known test vectors) replaces the old plain-text fallback. Existing plain-text PINs are upgraded automatically on the next successful unlock
- **Constant-time PIN comparison**, and PINs must be 4–6 digits
- **Change PIN now counts against the lockout**, so it can't be used to guess the current PIN
- **Stricter Content Security Policy** — `connect-src 'none'`, `object-src 'none'`, `base-uri 'none'`, `form-action 'none'`, `frame-src 'none'`
- **nginx hardening** — `Referrer-Policy: no-referrer`, `Permissions-Policy`, `Cross-Origin-Opener-Policy`, `server_tokens off`, and `blob:` allowed so the home-screen manifest works behind the bundled config (rebuild the image to pick this up)
- Service URLs without a scheme (e.g. `192.168.1.5:8080`) are now saved as `http://…`

### Removed
- About 100 lines of dead credential-encryption code left over from auto-login, including a derived key that was kept in `sessionStorage`. Leftover `pagerr_creds_enc` / `pagerr_cred_salt` entries are deleted on load
- Broken trust-status helpers that referenced an undefined function

### Fixed
- Changing the PIN over plain HTTP crashed (it tried to use Web Crypto, which isn't available there)
- Sync codes no longer include the server logo, which made them enormous; HTTP-generated sync codes (`PAGERR2X:`) can now be imported

---

## [2.2.0]

### Added
- **Copy-paste sync** — generate an encrypted sync code on one device and paste it on another. No files, cloud storage, or accounts required. Works on HTTP (XOR, `PAGERR2X:` prefix) and HTTPS (AES-256-GCM, `PAGERR2:` prefix)
- **Random per-install PIN salt** — a 16-byte random salt is generated on first launch and stored locally. Every install is unique, making precomputed attacks useless
- **PIN brute-force lockout** — 5 failed attempts triggers a 30-second lockout with exponential backoff, up to 30 minutes. Attempt count survives page refreshes. Remaining attempts shown on screen
- **Change PIN requires current PIN** — the existing PIN must be verified before a new one can be set. Credentials are re-encrypted automatically on PIN change
- **Trust token expiry** — trusted device tokens now expire after 30 days. Legacy tokens are treated as expired and removed
- **PWA manifest** — runtime-generated web manifest using your server name and theme colour for a proper Add to Home Screen experience on Android and iOS
- **Apple touch icon** — proper PWA icon embedded as base64, used for iOS home screen shortcuts and Android Chrome installs

### Changed
- **Auto-login removed** — service cards now open directly via `window.location.href`. The browser handles session storage after the first manual login. Removed login type selector, credential fields, and Saved Credentials settings section
- **Backup section renamed** — the old "Data & Transfer" section is now labelled "Backup" to distinguish it from the new Sync feature
- **PIN dots dynamic** — the unlock screen now renders the correct number of dots based on PIN length (4, 5, or 6) instead of always showing 4

### Fixed
- **6-digit PIN broken** — auto-submit previously fired at 4 digits regardless of PIN length, cutting off 5 and 6 digit PINs mid-entry
- **Auto-submit never worked** — the original code checked if the stored PBKDF2 hash was 4 characters long before auto-submitting. A PBKDF2 hash is never 4 characters, so it never fired
- **Blank screen on load** — `setTheme()` could crash before `initLockScreen()` ran if a DOM element was missing, leaving a blank lock screen background. Wrapped in try-catch with a safety fallback
- **Clipboard errors on HTTP** — sync and curl command generators previously threw an error toast when `navigator.clipboard` was unavailable over HTTP. Now caught silently with a visible text box fallback

### Removed
- Auto-login support (Arr stack, SABnzbd, qBittorrent, Deluge, Transmission, NZBGet, Tautulli, Pi-hole, Wizarr)
- Saved Credentials settings section
- Login type selector from add/edit service form

---

## [2.1.1]

### Fixed
- Minor stability improvements

---

## [2.1.0]

### Added
- Encrypted JSON export and import with AES-256-GCM (HTTPS) or XOR obfuscation (HTTP)
- Service status dots (green/red) showing live reachability
- Reorder mode — drag services within and across categories
- PBKDF2 PIN hashing (100,000 iterations, SHA-256)
- Device trust tokens — stay unlocked on trusted devices
- Biometric unlock via WebAuthn
- Auto-lock timeout (5 min / 10 min / 30 min / 1 hour / off)
- Icon size selector (Small / Medium / Large)
- Text size selector (Small / Medium / Large)
- Toggleable service labels

---

## [1.0.0]

### Added
- Initial release
- Single-file HTML dashboard
- PIN lock screen with 4–6 digit PIN
- Auto icon matching via dashboard-icons
- Service categories with add/edit/delete
- Dark and light theme
- Custom server logo and name
- Docker Compose deployment with nginx

[3.0.0]: https://github.com/mortaljinx/pagerr/releases/tag/v3.0.0
[2.2.1]: https://github.com/mortaljinx/pagerr/releases/tag/v2.2.1
