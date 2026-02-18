::: {align="center"}
`<img src="docs/logo.png" width="200" alt="Pagerr">`{=html}

# Pagerr

**A PIN-protected, single-file homelab dashboard.**

![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)
![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)
![Self-hosted](https://img.shields.io/badge/Self--hosted-yes-22c55e)
![Tailscale](https://img.shields.io/badge/Tailscale-friendly-4d8eff)

Pagerr is a single HTML file that provides fast, mobile-first access to
your self-hosted services.

No backend.\
No database.\
No accounts.

Everything runs locally in your browser.

Built for Tailscale. Works behind any reverse proxy.
:::

------------------------------------------------------------------------

## Screenshots

  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
                          Lock Screen                                              Dashboard (Dark)                                            Dashboard (Light)                                                  Settings
  ----------------------------------------------------------- ----------------------------------------------------------- ------------------------------------------------------------ ---------------------------------------------------------------
   `<img src="docs/screenshot-lock.png" width="180">`{=html}   `<img src="docs/screenshot-dark.png" width="180">`{=html}   `<img src="docs/screenshot-light.png" width="180">`{=html}   `<img src="docs/screenshot-settings.png" width="180">`{=html}

  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

------------------------------------------------------------------------

## What Is Pagerr?

Pagerr is designed for **speed and simplicity**.

It is not a homepage replacement.\
It is not a monitoring dashboard.

It is your **remote control** for your server.

Open. Tap. Done.

------------------------------------------------------------------------

## Core Principles

-   Single static file\
-   Mobile-first design\
-   Zero external dependencies\
-   No telemetry or tracking\
-   Works offline once loaded\
-   Fully self-contained

------------------------------------------------------------------------

## Features

-   **Single-file app** --- everything lives in `index.html`
-   **PIN lock (4--6 digits)** with optional biometric unlock
-   **Auto icon matching** from dashboard-icons
-   **Category organisation**
-   **Tap-to-swap reorder mode**
-   **Auto-login support for common homelab apps**
-   **JSON export & import**
-   **Dark / Light theme**
-   **Adjustable icon and text sizing**
-   **Custom server name and logo**
-   **Add to home screen (PWA-style)**

------------------------------------------------------------------------

## Quickstart (Docker -- 30 seconds)

``` bash
git clone https://github.com/mortaljinx/pagerr.git
cd pagerr
docker compose up -d
```

Open:

http://your-server-ip:10000

Set your PIN on first launch and start adding services.

------------------------------------------------------------------------

## Docker Compose

``` yaml
version: "3.8"

services:
  pagerr:
    image: nginx:alpine
    container_name: pagerr
    ports:
      - "10000:80"
    volumes:
      - ./:/usr/share/nginx/html:ro
    restart: unless-stopped
```

Repository structure:

    pagerr/
     ├── index.html
     └── docker-compose.yml

------------------------------------------------------------------------

## Serve Without Docker

Pagerr is just a static file.

``` bash
# Python
python3 -m http.server 10000

# Node
npx serve . -p 10000
```

Or simply open `index.html` directly in a browser.

------------------------------------------------------------------------

## Updating

``` bash
git pull
docker compose up -d
```

Your configuration lives in browser `localStorage`.\
Updating Pagerr never touches your services.

------------------------------------------------------------------------

## Backup & Restore

Settings → **Export JSON** to save your configuration.

To restore:

Settings → **Import JSON**

If you clear browser storage, your export file is the only recovery
method.

------------------------------------------------------------------------

## Auto-Login Support

Supported services include:

-   Arr stack (Radarr, Sonarr, Lidarr, etc.)
-   SABnzbd
-   qBittorrent
-   Deluge
-   Transmission
-   NZBGet
-   Tautulli
-   Pi-hole
-   Wizarr

Unsupported services simply open normally.

Auto-login works best when Pagerr and your services share the same
domain or Tailscale network.

------------------------------------------------------------------------

## Access Methods

### Tailscale (Recommended)

http://your-tailscale-ip:10000

No ports exposed. No firewall changes required.

### Reverse Proxy

Works with Traefik, Caddy, Nginx Proxy Manager, etc.\
Point the proxy to port `10000`.

------------------------------------------------------------------------

## Install as Home Screen App

Pagerr works like an app when added to your phone home screen.

**iOS:** Share → Add to Home Screen\
**Android:** Menu → Add to Home Screen / Install App

Launches full-screen.

------------------------------------------------------------------------

## Security Notes

-   All data is stored in browser `localStorage`
-   Credentials are stored locally (not encrypted)
-   The lock screen is a convenience barrier, not cryptographic
    protection
-   HTTPS is strongly recommended
-   Best used behind Tailscale or a VPN

------------------------------------------------------------------------

## Project Status

Pagerr is a personal project shared publicly.

It is stable and intentionally minimal.

If it fits your workflow, use it.\
If not, fork it.

------------------------------------------------------------------------

## License

MIT --- see LICENSE.

------------------------------------------------------------------------

*Built for homelabbers, by a homelabber.*
