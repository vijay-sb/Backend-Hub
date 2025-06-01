const express = require('express')
const todosRouter = require('./routes/todos.js')
const port = 3000;


const app = express();
app.use(express.json());

app.get("/",(req,res)=>res.send("Todo api is live"));
app.use("/",todosRouter);



