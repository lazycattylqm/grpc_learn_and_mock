const path = require('path')
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const init = (protoRelativePath) => {
    const protoPath = path.join(__dirname, protoRelativePath);
    const packageDefinition = protoLoader.loadSync(
        protoPath,
        {
            keepCase: true,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true
        });
    return grpc.loadPackageDefinition(packageDefinition)
}

module.exports = {
    grpc,
    init
}