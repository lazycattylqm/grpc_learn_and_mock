let messages = require('./hello_service_pb');
let services = require('./hello_service_grpc_pb');
const grpc = require('@grpc/grpc-js')

module.exports = {
    grpc,
    messages,
    services
}