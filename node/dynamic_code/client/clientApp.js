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
const target = 'localhost:3000';
const client = new helloServiceProto.HelloService(target,
    grpc.credentials.createInsecure());

client.SayHello({ name: "liqiming node dy" }, function (err, response) {
    console.log('Greeting:', response.message);
});