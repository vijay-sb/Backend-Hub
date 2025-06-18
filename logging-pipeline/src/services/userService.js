const users = [];
let currId = 1;

exports.getAllUser = () => users;
exports.getUserById = (id) => users.find(u => u.id === parseInt(id));
exports.createUser = ({ name,email } ) =>
{
  if(!name || !email)
{
    throw new Error('Name and Email are required');
  }
  const newUser={id:currId++,name,email};
  users.push(newUser);
  return newUser;
}


