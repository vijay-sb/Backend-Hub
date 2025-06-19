const express = require('express');
const app = express();
const userRoutes = require('./routes/userRoutes.js')
const requestLogger = require('./utils/requestLogger.js')
const logger = require('./config/logger.js')

app.use(express.json())
app.use(requestLogger);

app.use('/users',userRoutes);

app.get('/',(req,res)=>
{
    logger.info({message:'Server staring '});
    res.send("Hello World");

  }
);

//Global Error Handler
app.use((err, req, res, next) => {
  logger.error({
    message: err.message,
    stack: err.stack,
    url: req.originalUrl,
  });
  res.status(500).json({ error: 'Internal Server Error' });
});


const PORT = process.env.PORT || 3000;
app.listen(PORT, () => logger.info({ message: `Server running on port ${PORT}` }));
