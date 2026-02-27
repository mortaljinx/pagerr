# Changelog
All notable changes to Pagerr will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [2.1.0] — 2026-02-27

### Added
- **Service status dots** — small green/red indicator on each card showing if the service is reachable. Uses favicon ping, no CORS issues
- **Disable lock screen** — Settings → Lock Screen → toggle off to skip PIN entirely. Lock button in header hides automatically. Toggle back on to re-enable
- **Docker Hub image** — `mortaljinx/pagerr:latest` published to Docker Hub. Deploy with a single `docker run` command, no file download needed
- **GitHub Actions** — automatic build and push to Docker Hub on every push to main, for both `amd64` and `arm64`

### Changed
- Lock and settings buttons in header increased in size for easier tapping on mobile
- Add service sheet no longer auto-opens after creating a first category
- Change PIN section hidden automatically when lock screen is disabled

### Fixed
- Export JSON button no longer shows highlight ring or requires double-tap on mobile
- Encrypt modal no longer renders visibly at the bottom of the page on load

---

## [2.0.0] — 2026-02-18

### Added
- **Service reorder** — tap ⇅ reorder in the category header to enter reorder mode. Tap any two cards to swap their positions. Works across categories. Tap ✓ done to exit
- **Encrypted JSON export** — Settings → Data & Transfer → Export JSON. Password-protected export using AES-256-GCM over HTTPS or password obfuscation over HTTP
- **Encrypted JSON import** — Settings → Data & Transfer → Import JSON. Detects encryption method automatically and prompts for password
- **Server name and logo included in export** — full config backup including branding
- **Data & Transfer section** added to Settings panel
- Wizarr added to auto-login support

### Changed
- README fully rewritten with updated feature list, backup/restore documentation, screenshots, and personal project notice
- New logo

### Removed
- QR code export/import (config payload too large for reliable QR encoding)

---

## [1.0.0] — 2025

### Added
- Initial release
- Single-file HTML dashboard with PIN and biometric (WebAuthn) lock screen
- Auto icon matching via [dashboard-icons](https://github.com/homarr-labs/dashboard-icons)
- Auto-login support for Arr stack, SABnzbd, qBittorrent, Deluge, Transmission, NZBGet, Tautulli, Pi-hole
- Dark and light theme
- Service categories with add/edit/delete
- Adjustable icon size (Small / Medium / Large)
- Adjustable text size (Small / Medium / Large)
- Toggleable service labels
- Custom server logo and name
- Auto-lock timeout (5 min / 10 min / 30 min / 1 hour / off)
- Saved credentials manager
- Docker Compose deployment with nginx
