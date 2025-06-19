const { createLogger, format, transports } = require('winston');
const net = require('net');

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
    new transports.Stream({
      stream: net.connect({ port: 5000, host: 'logstash' }),
    }),
  ],
});

module.exports = logger;
