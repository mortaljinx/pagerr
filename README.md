<div align="center">

# 🗂️ Pagerr

**A self-hosted, PIN-protected bookmark dashboard for your homelab.**

Pagerr is a single-file web app that gives you a clean, mobile-friendly landing page for all your self-hosted services. Add your services, organise them into categories, and access everything behind a PIN or biometric lock — no database, no backend, no fuss.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)](docker-compose.yml)
[![Self-hosted](https://img.shields.io/badge/Self--hosted-yes-22c55e)](#)

</div>

---

## ✨ Features

- **Single HTML file** — no build step, no database, no dependencies
- **PIN & biometric lock** — protects your dashboard with a 4–6 digit PIN or device fingerprint/FaceID
- **Auto icon matching** — automatically fetches service icons from [dashboard-icons](https://github.com/homarr-labs/dashboard-icons) by name
- **Auto-login support** — stores credentials and auto-submits login forms for supported services
- **Categories** — organise services into labelled groups
- **Dark & light theme** — smooth system-aware theming
- **Customisable display** — adjustable icon size, text size, and toggleable service labels
- **Custom branding** — upload your own server logo and set a server name
- **Zero external dependencies** — everything runs from the single HTML file; data stored in `localStorage`
- **Mobile-first** — designed for phones and tablets as well as desktop browsers

---

## 🚀 Quick Start

### Docker Compose (recommended)

```bash
git clone https://github.com/yourusername/pagerr.git
cd pagerr
docker compose up -d
```

Then open [http://localhost:8080](http://localhost:8080) in your browser.

> On first launch you'll be prompted to set a PIN and optionally a server name and logo.

### Serve manually (no Docker)

Because Pagerr is a single HTML file you can serve it with any web server, or even open it directly in a browser.

```bash
# Python
python3 -m http.server 8080

# Node (npx)
npx serve . -p 8080
```

---

## 🐳 Docker Compose

See [`docker-compose.yml`](docker-compose.yml) for the full template. Key defaults:

| Variable | Default | Description |
|---|---|---|
| Port | `8080` | Host port to access Pagerr |
| Volume | `./pagerr.html:/usr/share/nginx/html/index.html:ro` | Mounts the app file |

To change the port, edit the `ports` mapping in `docker-compose.yml`.

---

## 🔐 Auto-Login Support

Pagerr can automatically submit login credentials for the following services:

| Service | Type |
|---|---|
| Radarr / Sonarr / Lidarr / Prowlarr (Arr stack) | Form POST |
| SABnzbd | Form POST |
| qBittorrent | Form POST |
| Deluge | Form POST |
| Transmission | Basic Auth |
| NZBGet | Basic Auth |
| Tautulli | Form POST |
| Pi-hole | Form POST |

> **Note:** Services not in this list will still open — Pagerr just won't auto-submit the login form. The browser will remember your session after the first manual login.

---

## 🎨 Customisation

All settings are accessible from the ⚙️ icon in the top-right corner of the dashboard:

- **Theme** — Dark / Light
- **Service Labels** — toggle service name text on/off beneath icons
- **Icon Size** — Small / Medium / Large
- **Text Size** — Small / Medium / Large
- **Auto-lock Timeout** — 5 min / 10 min / 30 min / 1 hour / off
- **Categories** — add, rename, or delete categories
- **Saved Credentials** — view and remove stored auto-login credentials
- **Change PIN** — update your access PIN at any time

---

## 📁 File Structure

```
pagerr/
├── pagerr.html          # The entire application (single file)
├── docker-compose.yml   # Docker Compose deployment template
├── nginx.conf           # Minimal nginx config used by the container
├── .gitignore
├── LICENSE
└── README.md
```

---

## 🛡️ Security Notes

- All data (services, credentials, settings) is stored in your browser's `localStorage` — nothing is sent to any server
- Credentials are stored in plaintext in `localStorage`; Pagerr's PIN lock is a convenience lock, not encryption
- It is **strongly recommended** to serve Pagerr over HTTPS, especially if exposed outside your LAN
- Use a reverse proxy (Nginx Proxy Manager, Traefik, Caddy) to add HTTPS and optional IP allowlisting

---

## 🔄 Updating

```bash
git pull
docker compose down
docker compose up -d
```

Your data lives in `localStorage` in your browser, not in the container, so pulling a new version will never wipe your services or settings.

---

## 🤝 Contributing

Contributions are welcome! Please open an issue first to discuss what you'd like to change.

1. Fork the repo
2. Create a feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -m 'Add my feature'`)
4. Push to the branch (`git push origin feature/my-feature`)
5. Open a Pull Request

---

## 📄 License

MIT — see [LICENSE](LICENSE) for details.

---

<div align="center">
  <sub>Built for homelabbers, by homelabbers.</sub>
</div>
