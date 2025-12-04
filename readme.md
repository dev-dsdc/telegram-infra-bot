# Telegram Infra Bot

![CI](https://github.com/dev-dsdc/telegram-infra-bot/actions/workflows/ci.yml/badge.svg?branch=main)
![Go Version](https://img.shields.io/badge/Go-1.25-blue)
![License](https://img.shields.io/badge/license-MIT-green)

**Инфраструктурный Telegram-бот** для автоматизации  проверок и мониторинга.
Проект демонстрирует: контейнеризацию (Docker), CI/CD (GitHub Actions), безопасную работу с секретами и базовый мониторинг (healthcheck).

---

## Ключевые возможности

- Автозапуск бота в Docker-контейнере
- Deploy по релизным тэгам (`v*`) через GitHub Actions
- Healthcheck endpoint (`/health`) для Docker
- Конфигурация через переменные окружения (без секретов в репозитории)
- Пример интеграции с Docker API / мониторингом (расширяемо)

---

## Структура проекта

## Tech Stack
- Go 1.25+
- Telegram Bot API
- Docker & Docker Compose
- GitLab CI/CD
- Alpine


## Project Structure
```
.
├── cmd/
│   └── bot/                 # main.go бота
├── internal/
│   └── health/              # HTTP health server
├── Dockerfile
├── docker-compose.yml
├── .env.example
├── .github/
│   └── workflows/
│       └── deploy.yml       # GitHub Actions деплой
├── README.md
├── go.mod
└── go.sum
```

## Requirements
- Docker (локально или в WSL2)
- Go (либо собирать через Docker)
- (опционально) docker-compose

## Installation
 - git clone https://github.com/dev-dsdc/telegram-infra-bot.git
 - cd project-name
 - cp .env.example .env # заполните переменные


## Run with Docker
- docker build -t telegram-infra-bot:local .
- docker run -d --name telegram-bot \
  -e BOT_TOKEN="$(cat .env | grep BOT_TOKEN | cut -d '=' -f2)" \
  --restart unless-stopped \
  telegram-infra-bot:local 

## Run with docker-compose
- export BOT_TOKEN="123456:ABC..."
  docker compose up -d

## Deploy
Проект настроен на deploy по тегам. Чтобы задеплоить релиз:
- git tag v1.0.0
- git push origin v1.0.0
  GitHub Actions:

## GitHub Actions:
- Сборка Docker-образа и пуш в registry (ghcr.io/<owner>/<repo>:v1.0.0 и :latest)
- Deploy на self-hosted runner (или, при другой конфигурации, через SSH)
- deploy.yml настроено запускать workflow только при push тегов v*.

## Environment Variables
- BOT_TOKEN=токен бота
- GHCR_TOKEN=персональный токен доступа 

## Healthcheck

- Внутри контейнера поднят минимальный HTTP-сервер на :8080 и эндпоинт /health. Docker использует HEALTHCHECK в Dockerfile для определения состояния контейнера (healthy/unhealthy)

## CI/CD
 - Workflow установлен в .github/workflows/deploy.yml
 - Workflow триггерится по push тегам v*
 - Secrets: BOT_TOKEN, GHCR_TOKEN (подставляется автоматически)
 - Для локального автоматического деплоя рекомендуется использовать self-hosted runner

# Useful commands
### Local  build
- go build -o bot ./cmd/bot
### Run with docker
- docker build -t telegram-infra-bot:local .
- docker run -d --name telegram-bot -e BOT_TOKEN="..." telegram-infra-bot:local
### Check health
curl -sS http://localhost:8080/health

# License

- MIT © dev-dsdc
