const client = require('prom-client');
const register = new client.Registry();

client.collectDefaultMetrics({register})

// Custom Metrics

const HttpRequestCounter = new client.Counter({
  name:'https_req_total',
  help: 'Total Number of Http Requests',
  labelNames: ['method','route','status']
});

const ResponseTimeHistogram = new client.Histogram({
  name:'response_time',
  help:"Duration of Http Requests",
  labelNames: ['method','route','status']
});

register.registerMetric(HttpRequestCounter);
register.registerMetric(ResponseTimeHistogram);

module.exports = {
  register,
  HttpRequestCounter,
  ResponseTimeHistogram,
};
