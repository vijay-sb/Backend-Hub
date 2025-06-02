# Backend-Hub
A space to learn/try-out new backend tools

# Monitoring-Pipeline
This is a simple Node.js TODO backend with an integrated **observability pipeline** using **Prometheus** and **Grafana**. The project is fully containerized using Docker.

## 🚀 Setup

```bash
git clone https://github.com/vijay-sb/Backend-Hub.git
cd Backend-Hub/monitoring-pipeline
docker compose up --build

API: http://localhost:3000

Metrics: http://localhost:3000/metrics

Prometheus: http://localhost:9090

Grafana: http://localhost:3001
