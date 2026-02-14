<div align="center">

# 🗂️ Pagerr

**A PIN-protected homelab dashboard built for Tailscale — and anywhere else you self-host.**

Pagerr is a single-file web app that gives you a clean, mobile-friendly landing page for all your self-hosted services. Designed primarily for [Tailscale](https://tailscale.com) users who want quick, private access to their server from any device on their tailnet — no exposure to the open internet required. It also works great behind reverse proxies like Traefik, Caddy, Nginx Proxy Manager, or any other setup you already run.

Add your services, organise them into categories, and access everything from behind a PIN or fingerprint lock. No database, no backend, no fuss.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)](docker-compose.yml)
[![Self-hosted](https://img.shields.io/badge/Self--hosted-yes-22c55e)](#)
[![Tailscale](https://img.shields.io/badge/Tailscale-friendly-4d8eff)](#)

</div>

---

## ✨ Features

- **Single HTML file** — no build step, no database, no dependencies
- **PIN lock** — works in all modern browsers; 4–6 digit PIN set on first launch
- **Fingerprint / biometric lock** — WebAuthn-based; currently supported in **Firefox** (Chrome/Safari support varies by platform)
- **Add to home screen** — install Pagerr as an icon on your phone or tablet for an app-like experience
- **Auto icon matching** — automatically fetches service icons from [dashboard-icons](https://github.com/homarr-labs/dashboard-icons) by name
- **Auto-login support** — stores credentials and auto-submits login forms for supported services
- **Categories** — organise services into labelled groups
- **Dark & light theme** — smooth, switchable theming
- **Customisable display** — adjustable icon size, text size, and toggleable service labels
- **Custom branding** — upload your own server logo and set a server name
- **Zero tracking** — all data stored locally in your browser's `localStorage`; nothing leaves your device

---

## 🌐 Access Methods

Pagerr is built with private, local-network access in mind. Here are the recommended ways to use it:

### 🔒 Tailscale (recommended)
The cleanest setup. Run Pagerr on your server and access it privately from any device on your tailnet — phone, tablet, laptop — without opening any ports or configuring DNS.

```
http://your-tailscale-ip:8484
```

No reverse proxy needed. Tailscale handles the secure tunnel.

### 🔀 Reverse Proxy
If you already run a reverse proxy, Pagerr slots right in behind it. The `docker-compose.yml` includes commented-out Traefik labels as a starting point.

| Proxy | Notes |
|---|---|
| **Traefik** | Labels included in `docker-compose.yml` — just uncomment and set your domain |
| **Caddy** | Point a `reverse_proxy` block at `localhost:8484` |
| **Nginx Proxy Manager** | Add a new proxy host pointing to your container |
| **Nginx** | Standard `proxy_pass` to the container port |

---

## 🚀 Quick Start

```bash
git clone https://github.com/mortaljinx/pagerr.git
cd pagerr
docker compose up -d
```

Then open `http://your-server-ip:8484` in your browser — or your Tailscale IP if accessing remotely.

> On first launch you'll be prompted to set a PIN, and optionally a server name and logo.

### Serve manually (no Docker)

Because Pagerr is a single HTML file you can serve it with any web server, or open it directly in a browser.

```bash
# Python
python3 -m http.server 8484

# Node
npx serve . -p 8484
```

---

## 📱 Install as a Home Screen App

Pagerr works as a Progressive Web App — you can add it to your phone or tablet home screen for an app-like experience with no browser chrome.

**iOS (Safari):** Open Pagerr → tap the Share button → **Add to Home Screen**

**Android (Chrome/Firefox):** Open Pagerr → tap the browser menu → **Add to Home Screen** or **Install App**

Once installed it launches full-screen, just like a native app.

---

## 🐳 Docker Compose

See [`docker-compose.yml`](docker-compose.yml) for the full template. Key defaults:

| Setting | Default | Description |
|---|---|---|
| Port | `8484` | Host port to access Pagerr |
| Volume | `./pagerr.html:/usr/share/nginx/html/index.html:ro` | Mounts the app file |
| Restart | `unless-stopped` | Auto-restarts with Docker |

To change the port, edit the `ports` mapping in `docker-compose.yml`.

---

## 🔐 Auto-Login Support

Pagerr can automatically submit login credentials for the following services when you tap a card:

| Service | Method |
|---|---|
| Radarr / Sonarr / Lidarr / Prowlarr / Readarr (Arr stack) | Form POST |
| SABnzbd | Form POST |
| qBittorrent | Form POST |
| Deluge | Form POST |
| Transmission | Basic Auth |
| NZBGet | Basic Auth |
| Tautulli | Form POST |
| Pi-hole | Form POST |

> Services not in this list will open normally — the browser will remember your session after the first manual login.

> **Note on CORS:** Auto-login works best when Pagerr and your services share the same domain or are accessed over Tailscale. Some services block cross-origin requests regardless of configuration.

---

## 🎨 Customisation

All settings are accessible from the ⚙️ icon in the top-right corner:

- **Theme** — Dark / Light
- **Service Labels** — toggle service name text on/off beneath icons
- **Icon Size** — Small / Medium / Large
- **Text Size** — Small / Medium / Large
- **Auto-lock Timeout** — 5 min / 10 min / 30 min / 1 hour / off
- **Categories** — add, rename, or delete categories
- **Saved Credentials** — view and remove stored auto-login credentials
- **Change PIN** — update your access PIN at any time

---

## 🔑 Browser Compatibility

| Feature | Chrome | Firefox | Safari | Edge |
|---|---|---|---|---|
| PIN lock | ✅ | ✅ | ✅ | ✅ |
| Fingerprint / biometric | ⚠️ | ✅ | ⚠️ | ⚠️ |
| Add to home screen | ✅ | ✅ | ✅ | ✅ |
| Auto-login | ✅ | ✅ | ✅ | ✅ |

> Biometric (WebAuthn) support outside Firefox depends on your OS and browser version. PIN is always available as a fallback.

---

## 📁 File Structure

```
pagerr/
├── pagerr.html          # The entire application (single file)
├── docker-compose.yml   # Docker Compose deployment template
├── nginx.conf           # Minimal nginx config used by the container
├── .gitignore
├── CHANGELOG.md
├── LICENSE
└── README.md
```

---

## 🛡️ Security Notes

- All data (services, credentials, settings) is stored in your browser's `localStorage` — nothing is sent to any server
- Credentials are stored in plaintext in `localStorage`; Pagerr's lock screen is a convenience lock, not encryption
- Serving over **HTTPS is strongly recommended**, especially if accessing outside your LAN — Tailscale provides this automatically via its encrypted tunnel
- For reverse proxy setups, use Traefik / Caddy / NPM to terminate TLS and optionally restrict access by IP

---

## 🔄 Updating

```bash
git pull
docker compose down
docker compose up -d
```

Your data lives in `localStorage` in your browser, not in the container — updating will never wipe your services or settings.

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
