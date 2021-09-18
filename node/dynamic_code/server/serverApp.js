const {init, grpc} = require('proto')

const helloServiceProto = init('hello_service.proto').hello;

function sayHello(call, callback) {

    callback(null, { message: `Hello ${call.request.name}, this is node server` });
}

const server = new grpc.Server();
server.addService(helloServiceProto.HelloService.service, { sayHello: sayHello });
server.bindAsync('0.0.0.0:3000', grpc.ServerCredentials.createInsecure(), () => {
    server.start();
});