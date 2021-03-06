const { grpc, messages, services } = require('proto');

const target = 'localhost:3000';

const client = new services.HelloServiceClient(target, grpc.credentials.createInsecure());

const request = new messages.HelloRequest();

request.setName('liqiming node static')

client.sayHello(request, function (err, response) {
    console.log('Greeting:', response.getMessage());
});