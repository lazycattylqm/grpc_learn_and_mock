const { services, grpc } = require('proto')

const { sayHello } = require('./service/helloService')

const server = new grpc.Server();
server.addService(services.HelloServiceService, { sayHello: sayHello });
server.bindAsync('localhost:3000', grpc.ServerCredentials.createInsecure(), () => {
    server.start();
});

