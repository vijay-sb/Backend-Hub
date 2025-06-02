const { HttpRequestCounter, ResponseTimeHistogram } = require('../metrics/prometheus.js');

module.exports = function metricsMiddleware(req, res, next) {
  
  const end = ResponseTimeHistogram.startTimer({ method: req.method, route: req.path });

  res.on('finish', () => {
    HttpRequestCounter.labels(req.method, req.path, res.statusCode).inc();
    end();
  });

  next();
};

