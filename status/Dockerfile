# stage 1: generate HTML
FROM ubuntu:24.04 AS builder

WORKDIR /build
COPY generate-status .
RUN chmod +x generate-status && ./generate-status > index.html

# stage 2: serve with nginx (unprivileged)
FROM nginxinc/nginx-unprivileged:1.28

COPY --from=builder /build/index.html /usr/share/nginx/html/index.html

EXPOSE 8080
