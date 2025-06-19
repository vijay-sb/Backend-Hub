//like a middleware 

const logger = require('../config/logger.js')

const requestLogger = (req,res,next) =>
{
  logger.info({
   message: 'Incoming request',
    method: req.method,
    url: req.originalUrl,
    ip: req.ip,
    userAgent: req.get('User-Agent'),
    timestamp: new Date().toISOString(),
  });
  next();
};

module.exports = requestLogger;
