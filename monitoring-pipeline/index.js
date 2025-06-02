const express = require('express')
const todosRouter = require('./routes/todos.js')
//const bodyParser = require('body-parser')
const metricsMiddleware = require('./middleware/metricsMiddleware.js')
const { register } = require('./metrics/prometheus.js')
const port = 3000;


const app = express();
app.use(express.json());
app.use(metricsMiddleware);

app.get("/",(req,res)=>res.send("Todo api is live"));
app.use('/todos',todosRouter);

app.get('/metrics',async (req,res) =>

{
    res.set('Content-type',register.contentType)
    res.end(await register.metrics());
  })


app.listen(port,()=>{
  console.log(`Server is running on port: ${port}`);
})


