
const { createLogger, format, transports } = require('winston');
const net = require('net');

let logstashStream;

try {
  logstashStream = net.connect({ host: 'logstash', port: 5000 });
  logstashStream.on('error', err => {
    console.error('⚠️ Logstash TCP stream error:', err.message);
  });
} catch (err) {
  console.error('⚠️ Failed to connect to Logstash:', err.message);
}

const logger = createLogger({
  level: 'info',
  format: format.combine(
    format.timestamp(),
    format.errors({ stack: true }),
    format.json()
  ),
  defaultMeta: { service: 'elk-logging-backend' },
  transports: [
    new transports.Console(),
    ...(logstashStream ? [new transports.Stream({ stream: logstashStream })] : [])
  ]
});

module.exports = logger;
