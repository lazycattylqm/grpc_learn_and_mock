const basePath = __dirname
const path = require('path')
const protoPath = path.join(basePath, 'proto', 'hello_service.proto');

const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const packageDefinition = protoLoader.loadSync(
    protoPath,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });

const helloServiceProto = grpc.loadPackageDefinition(packageDefinition).hello;

function sayHello(call, callback) {
    callback(null, { message: 'Hello ' + call.request.name });
}

var server = new grpc.Server();
server.addService(helloServiceProto.HelloService.service, { sayHello: sayHello });
server.bindAsync('0.0.0.0:3000', grpc.ServerCredentials.createInsecure(), () => {
    server.start();
});