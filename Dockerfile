# Stage 1: build the tiny Quick connect relay (static binary, no runtime needed)
FROM golang:1.24-alpine AS relay
WORKDIR /src
COPY relay/ ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /pagerr-relay .

# Stage 2: nginx serving the single-file app
FROM nginx:alpine
COPY --from=relay /pagerr-relay /usr/local/bin/pagerr-relay
COPY --chmod=755 docker/40-pagerr-relay.sh /docker-entrypoint.d/40-pagerr-relay.sh
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY index.html /usr/share/nginx/html/index.html
EXPOSE 80
