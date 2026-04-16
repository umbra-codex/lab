# Lab

A personal collection of labs, projects, scripts, and experiments used for learning, exploring, and gaining hands-on experience.

---

## Contents

| Directory | Description | Tech |
|---|---|---|
| [`backup`](./backup) | Backup utility CLI with create, restore, and list operations | Bash, Docker |
| [`bash`](./bash) | Bash scripting examples organized by topic | Bash |
| [`bash-dotfiles`](./bash-dotfiles) | Shell and editor configuration files with setup script | Bash, Starship, Vim |
| [`bookbot`](./bookbot) | Book text analysis tool — word and character frequency | Python |
| [`docker-bookbot`](./docker-bookbot) | Containerized bookbot with Python built from source | Python, Docker |
| [`docker-goserver`](./docker-goserver) | Minimal HTTP server containerized with a lean image | Go, Docker |
| [`docker-load_balancer`](./docker-load_balancer) | Round-robin load balancer using Caddy reverse proxy | Caddy, Docker |
| [`greeter`](./greeter) | Greeting service demonstrating Docker environment variables | Bash, Docker |
| [`health`](./health) | Container health check configuration example | Docker, Nginx |
| [`joke-dashboard`](./joke-dashboard) | Auto-refreshing dad joke dashboard using a public API | Bash, Docker Compose, Nginx |
| [`k8s`](./k8s) | Kubernetes manifests for a multi-component chat application | Kubernetes |
| [`myapp`](./myapp) | Minimal containerized application for testing | Bash, Docker |
| [`status`](./status) | Static status page generator using a multi-stage Docker build | Bash, Docker, Nginx |
| [`status-page`](./status-page) | Enhanced status page with styling, health checks, and OCI labels | Bash, Docker, Nginx |
| [`webflyx`](./webflyx) | Sample movie catalog dataset in Markdown and CSV formats | Markdown, CSV |

---

## Topics Covered

- **Bash scripting** — variables, conditionals, loops, functions, arrays, parameter expansion
- **Dotfile management** — shell config, editor setup, prompt customization
- **Docker** — multi-stage builds, health checks, unprivileged users, OCI labels, environment variables
- **Docker Compose** — multi-container applications with shared volumes
- **Load balancing** — reverse proxy and round-robin configuration with Caddy
- **Kubernetes** — deployments, services, configmaps, persistent volumes, ingress, horizontal pod autoscaling
- **Go & Python** — small application examples with containerization
- **Data formats** — Markdown, CSV, YAML
