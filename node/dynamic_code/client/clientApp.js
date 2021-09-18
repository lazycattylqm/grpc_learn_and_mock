const { init, grpc } = require('proto');
const helloServiceProto = init('hello_service.proto').hello;
const target = 'localhost:3000';
const client = new helloServiceProto.HelloService(target,
    grpc.credentials.createInsecure());

client.SayHello({ name: "liqiming node dy" }, function (err, response) {
    console.log('Greeting:', response.message);
});