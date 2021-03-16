const { prisma } = require('./hello-world/generated/prisma-client');
const http = require('http')
http.createServer(async function (request, response){
  console.log(request.url)
  console.log(request.method)
  console.log(request.headers)
  const allUsers = await prisma.user({id: 'ckm5xq2b60011073175gphufi'});
  response.end(JSON.stringify(allUsers))
}).listen(8888)

// // 一个main函数，以便我们可以使用async/await
// async function main() {
//   // 新建一个user，并新建一个post文章
//   // const newUser = await prisma.createUser({
//   //   name: 'Alice',
//   // });
//   console.log('Created new user: ${newUser.name} (ID: ${newUser.id})');
//   // 从数据库中读取所有用户user并将其打印到控制台
//   const allUsers = await prisma.users();
//   console.log(allUsers[1].name);
// }
// main().catch(e => console.error(e));
