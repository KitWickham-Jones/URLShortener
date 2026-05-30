# URL Shortener

A distributed URL shortening service built in Go. Demonstrates a production-style architecture with caching, event streaming, Kubernetes orchestration, and observability.

## Architecture

```
POST /shorten → Go API → PostgreSQL
GET /:slug   → Go API → Redis (cache) → PostgreSQL → 302 redirect
                                      ↘ Kafka (click event)
                                              ↘ Python consumer → clicks table
```

## Stack

| Component | Technology |
|-----------|-----------|
| API | Go, stdlib `net/http` |
| Primary DB | PostgreSQL 16 |
| Cache | Redis (read-through, 24hr TTL) |
| Message Queue | Kafka (KRaft mode) |
| Analytics | Python, aiokafka, asyncpg |
| Orchestration | Kubernetes (Minikube) |
| Observability | Prometheus, Grafana |

## API

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/shorten` | Shorten a URL |
| GET | `/:slug` | Redirect to original URL |
| GET | `/health` | Health check |
| GET | `/metrics` | Prometheus metrics |

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com/very/long/url"}'

# {"short_url": "http://localhost:8080/4HTNpR"}
```

## Benchmarks

Load tested with `hey` (n=10,000, c=500 concurrent):

| Configuration | p50 | p99 | req/sec |
|---------------|-----|-----|---------|
| PostgreSQL only | 47ms | 110ms | 6,900 |
| Redis + 1 pod | 16ms | 77ms | 22,649 |
| Kubernetes 3 pods + Redis | 34ms | 120ms | 11,220 |

## Local Development

```bash
# Start infrastructure
docker compose up database redis

# Run API
cd services/urlshortener
go run cmd/server/main.go
```

## Kubernetes (Minikube)

```bash
# Build and load image
eval $(minikube docker-env)
docker build -t urlshortener:latest ./services/urlshortener

# Deploy
kubectl apply -R -f deployments/kubernetes/

# Access service
minikube tunnel
curl -I http://127.0.0.1/health

# Logs
kubectl logs -f -l app=urlshortener --prefix=true | grep -v health

# Prometheus
kubectl port-forward service/prometheus-service 9090:9090
```

## Directory Structure

```
services/
  urlshortener/       ← Go API
    cmd/server/
    internal/
      api/            ← handlers, server, slug generation
      store/          ← postgres, redis
      metrics/        ← prometheus counters
      config/
  analytics/          ← Python Kafka consumer
    consumer.py
    database.py
deployments/
  kubernetes/
    core/             ← deployment, service, HPA, configmap, secret
    databases/        ← postgres, redis
    observability/    ← prometheus, grafana
    analytics/        ← kafka
schema.sql
docker-compose.yml
```
