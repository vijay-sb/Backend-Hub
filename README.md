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
```



# Logging Pipeline
**Node.js Logging with ELK Stack (Elasticsearch + Logstash + Kibana)**

This project demonstrates a full-fledged logging system using **Winston**, **Logstash**, and **Elasticsearch**, visualized through **Kibana** — all containerized with **Docker Compose**.

## 🚀 Getting Started


### Run the Stack

Run the following command:


``` docker-compose up --build ```

Once running, access the services at:

    App: http://localhost:3000

    Kibana: http://localhost:5601

    Elasticsearch: http://localhost:9200

🔁 Logging Flow

Express (Winston) → Logstash (TCP 5000) → Elasticsearch → Kibana

    Winston transports logs in JSON format over TCP to Logstash.

    Logstash parses and forwards the logs to Elasticsearch.

    Kibana indexes and visualizes the logs.


# GO LANG
**A Space for learning and building in GO**


