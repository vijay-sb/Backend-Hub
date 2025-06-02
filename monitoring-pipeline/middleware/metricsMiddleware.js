const { HttpRequestCounter, ResponseTimeHistogram } = require('../metrics/prometheus.js');

module.exports = function metricsMiddleware(req, res, next) {
  // Start timer for response time measurement with method and route labels
  const end = ResponseTimeHistogram.startTimer({ method: req.method, route: req.path });

  // Listen for the 'finish' event on the response to stop the timer and increment counters
  res.on('finish', () => {
    // Increment the request counter with labels for HTTP method, route, and status code
    HttpRequestCounter.labels(req.method, req.path, res.statusCode).inc();

    // Stop the timer and record the response duration
    end();
  });

  // Proceed to the next middleware
  next();
};

