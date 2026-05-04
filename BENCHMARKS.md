# Benchmarks

## Setup

- Go HTTP server (net/http)
- PostgreSQL 16 (Docker container)
- Load balancer: none

## Baseline — GET /{slug} redirect (1000 req, 500 concurrency)

| Metric | Value |
|---|---|
| Requests/sec | 6,900 |
| Avg latency | 59.8ms |
| p50 | 46.8ms |
| p95 | 106.9ms |
| p99 | 110.2ms |
| Fastest | 25.4ms |
| Slowest | 111.6ms |

All 1000 responses: `302 Found`

**Bottleneck:** p75 jumps to 87ms — bimodal distribution suggests DB query latency under concurrency.
