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
