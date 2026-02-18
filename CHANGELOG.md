# Changelog
All notable changes to Pagerr will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [2.0.0] — 2026-02-18

### Added
- **Service reorder** — tap ⇅ reorder in the category header to enter reorder mode. Tap any two cards to swap their positions. Works across categories. Tap ✓ done to exit
- **JSON export** — Settings → Data & Transfer → Export JSON. Downloads a full backup of all services, categories, credentials and settings as `pagerr-config.json`
- **JSON import** — Settings → Data & Transfer → Import JSON. Restores a previously exported config from file
- **Data & Transfer section** added to Settings panel
- Wizarr added to auto-login support

### Changed
- README fully rewritten with updated feature list, backup/restore documentation and personal project notice
- New logo

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
