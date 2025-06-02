const express = require('express')
const router = express.Router();

let todos = [];
let id = 1;
 
router.get('/',(req,res)=>{res.json(todos)});
router.post('/',(req,res)=>
{
    const { title } = req.body;
    const todo = {id:id++,title,done:false}
    todos.push(todo);
    res.status(201).json(todo);
  });
router.put('/:id',(req,res) =>
{
    const todoId = parseInt(req.params.id,10);


    const todoIndex = todos.findIndex(t => t.id === todoId);
    if(todoIndex == -1)
  {
      return res.status(404).json({message: "Index Not Found"})
    }
    todos[todoIndex].done=!todo[todoIndex].done;
    return res.status(200).json(todos[todoIndex]);
  })

router.delete('/:id',(req,res) =>
{
    const todoId= parseInt(req.params.id,10);
    const todoindex = todos.findIndex(t=>t.id=== todoId);
    if (todoindex === -1) {
    return res.status(404).json({ message: 'Todo not found' });
  }
    todos.splice(todoindex,1);
    return res.status(200).json({message : 'Deleted Sucessfully'})

  })
module.exports = router;
