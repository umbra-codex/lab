# Lab

Hands-on work from courses, homelab builds, and one-off experiments. Each directory stands alone.
Where a directory came from a course, its name carries the source as a prefix: `bootdev-` marks
Boot.dev coursework.

---

## Contents

| Directory | Description | Tech |
|---|---|---|
| [`backup`](./backup) | Backup CLI with create, restore, and list operations | Bash, Docker |
| [`bash`](./bash) | Bash scripting examples organized by topic | Bash |
| [`bash-dotfiles`](./bash-dotfiles) | Shell and editor config files with a setup script | Bash, Starship, Vim |
| [`bootdev-bookbot`](./bootdev-bookbot) | Book text analyzer: word and character frequency | Python |
| [`bootdev-k8s`](./bootdev-k8s) | Kubernetes course manifests for a three-service crawler app: deployments, autoscalers, and Gateway API routing | Kubernetes |
| [`bootdev-learn-aws-ch10-ecs`](./bootdev-learn-aws-ch10-ecs) | Python HTTP server that pulls a value from SSM Parameter Store, packaged for ECS | Python, Docker, AWS |
| [`bootdev-megacorp`](./bootdev-megacorp) | Git 2 coursework: sample company data for branching, merge conflicts, reverts, and stashes | Git, Bash, CSV |
| [`bootdev-webflyx`](./bootdev-webflyx) | Sample movie catalog of titles, quotes, and a classics CSV | Markdown, CSV |
| [`bootdev-worldbanc`](./bootdev-worldbanc) | Terminals and Shells coursework: a mock bank filesystem of public and private trees, dated logs, transaction CSVs, and shell scripts | Bash, Go, CSV |
| [`docker-bookbot`](./docker-bookbot) | Containerized bookbot with Python built from source | Python, Docker |
| [`docker-goserver`](./docker-goserver) | Prebuilt Go HTTP server binary copied into a slim Debian image, port set by env var | Go, Docker |
| [`docker-load-balancer`](./docker-load-balancer) | Round-robin load balancer using a Caddy reverse proxy | Caddy, Docker |
| [`go-custom-package`](./go-custom-package) | Three Go modules wired together, importing a local `mystrings` package through a `replace` directive | Go |
| [`greeter`](./greeter) | Greeting service that reads its config from Docker environment variables | Bash, Docker |
| [`health`](./health) | Container health check configuration example | Docker, Nginx |
| [`joke-dashboard`](./joke-dashboard) | Auto-refreshing dad joke dashboard fed by a public API | Bash, Docker Compose, Nginx |
| [`kubecraft-mealie`](./kubecraft-mealie) | Recipe manager on Kubernetes with its own namespace, PVC, and LoadBalancer service | Kubernetes, Mealie |
| [`learn-go-with-tests`](./learn-go-with-tests) | Test-first Go exercises: integers, iteration, slices, structs, pointers and errors, maps, dependency injection | Go |
| [`monitoring`](./monitoring) | LoadBalancer service exposing Grafana from the kube-prometheus-stack release | Kubernetes, Grafana, Prometheus |
| [`myapp`](./myapp) | Minimal containerized app for testing | Bash, Docker |
| [`status`](./status) | Static status page built with a multi-stage Docker build | Bash, Docker, Nginx |
| [`status-page`](./status-page) | Status page with styling, health checks, and OCI labels | Bash, Docker, Nginx |

---

## Topics covered

- **Bash scripting**: variables, conditionals, loops, functions, arrays, parameter expansion
- **Shell and filesystem work**: navigating nested trees, searching text, running and signaling scripts
- **Dotfile management**: shell config, editor setup, prompt customization
- **Docker**: multi-stage builds, health checks, unprivileged users, OCI labels, environment variables
- **Docker Compose**: multi-container applications with shared volumes
- **Load balancing**: reverse proxy and round-robin configuration with Caddy
- **Kubernetes**: deployments, services, configmaps, persistent volumes, horizontal pod autoscaling, Gateway API routing
- **Observability**: Grafana and Prometheus on Kubernetes via kube-prometheus-stack
- **AWS**: SSM Parameter Store, container images for ECS
- **Go**: local module imports, plus a test-first pass through the language basics
- **Python**: small applications, containerized
- **Git**: branching, merge conflict resolution, rebasing, reverts, stashing, history rewriting
- **Data formats**: Markdown, CSV, YAML
