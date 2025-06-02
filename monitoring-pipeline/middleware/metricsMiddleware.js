const { httpRequestCounter, responseTimeHistogram} = require('../metrics/prometheus.js');
module.export = funciton metricsMiddleware(req, res, next)
{
  const end = responseTimeHistogram.startTimer({method:req.method , route:req.path});
  res.on('finish', () =>
  {
      httpRequestCounter.labels(req.method,req.path,res.statusCode).inc();
      end();
    });
  next();
};
