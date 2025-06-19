const userService = require('../services/userService.js')
const logger = require('../config/logger.js')

exports.getUsers=(req,res) =>
{
  const users = userService.getAllUsers();
  logger.info({message: 'Fetched all users' , count:users.length});
  res.status(200).json(users)
}

exports.getUserById = (req,res,next) =>
{
    try
  {
      const user = userService.getUserById(req.params.id);
      if(!user)
  {
      logger.warn({message: 'user not found', id: req.params.id});
      res.status(404).json({error:'User Not Fouun'})
    }
      logger.info({message:'User Fetched Successfully',id:req.params.id});
      res.status(200).json(user);
  }
  catch(error)
  {
    next(error);
  }

  } ;

 exports.createUser = (req,res,next) =>
{
  try{

  const user = userService.createUser(req.body);
  logger.info({message:'User Created',user});
  res.status(201).json(user);
  }
  catch(err)
  {
    next(err); //to Global Error Handler
  }
  }

