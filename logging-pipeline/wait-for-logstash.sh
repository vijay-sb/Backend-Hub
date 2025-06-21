#!/bin/sh

echo "⏳ Waiting for Logstash to be available at logstash:5000..."

# Loop until Logstash is accepting connections
while ! nc -z logstash 5000; do
  sleep 1
done

echo "✅ Logstash is up. Starting Node app..."
exec node src/index.js
