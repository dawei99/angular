const { prisma } = require('./hello-world/generated/prisma-client');
const { GraphQLServer } = require('graphql-yoga');
const resolvers = {
  Query: {
    users(root, args, context){
      return context.prisma.users({where: {id: args.where.id}})
    },
    user(root, args, context){
      console.log(args.id)
      return context.prisma.user({id: args.where.id})
    },
  },
  Mutation: {
    createUser(root, args, context) {
      return context.prisma.createUser({ name: args.name });
    },
  },
  User: {

  },
};

const server = new GraphQLServer({
  typeDefs: './schema.graphql',
  resolvers,
  context: {
    prisma,
  },
});
server.start(() => console.log('Server is running on http://localhost:4000'));