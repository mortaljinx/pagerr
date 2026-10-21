# Security policy

## Reporting a vulnerability

Please report security issues privately using GitHub's **Report a vulnerability** button on the
[Security tab](https://github.com/mortaljinx/pagerr/security/advisories/new) rather than opening a public issue.

Include what you found, how to reproduce it, and which version you tested. You'll get a reply as soon as possible,
and a fix will be released before details are published.

## Supported versions

Only the latest release receives security fixes. Pagerr is a single-file app, so updating is a file or image swap.

## Scope and security model

Pagerr is a convenience lock for a private dashboard, not a substitute for network-level protection. It runs entirely in
the browser: the PIN screen keeps casual access out, but anyone with access to a browser's developer tools on that device
can read the stored dashboard data. Run Pagerr behind Tailscale or an authenticating reverse proxy and don't expose it
directly to the internet. See the Security Notes section of the README for details.

## Quick connect relay

The official Docker image includes a small helper (`relay/`, written in Go) that Quick connect uses as a mailbox between two of your devices.

- It holds data **in memory only**, never on disk, for at most three minutes, and deletes it the moment the receiving device collects it
- Each mailbox is a one-shot slot under a random 128-bit ID that is only ever shown as a QR code on the receiving device. Slots can't be created for, or filled under, an ID that wasn't opened first
- The data is encrypted in the browser with a separate random 256-bit key (HMAC-SHA256 in counter mode with an HMAC tag). The key lives in the URL fragment (after `#`), which browsers never send to servers, so the relay and your web server logs can't read the dashboard
- Limits: 50 open slots, 512 KB per slot, 12 new slots per minute per client address, plus an nginx rate limit. It listens on `127.0.0.1` only and nginx is the only thing that reaches it, and it drops root privileges on start
- Over plain HTTP, anyone who can modify the page in transit could tamper with it. Use HTTPS or a trusted network (such as Tailscale) as with the rest of Pagerr
- If the helper isn't running (for example you serve only `index.html` from your own web server) Quick connect hides itself
