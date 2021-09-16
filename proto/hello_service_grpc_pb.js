// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var hello_service_pb = require('./hello_service_pb.js');

function serialize_HelloReply(arg) {
  if (!(arg instanceof hello_service_pb.HelloReply)) {
    throw new Error('Expected argument of type HelloReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_HelloReply(buffer_arg) {
  return hello_service_pb.HelloReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_HelloRequest(arg) {
  if (!(arg instanceof hello_service_pb.HelloRequest)) {
    throw new Error('Expected argument of type HelloRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_HelloRequest(buffer_arg) {
  return hello_service_pb.HelloRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var HelloServiceService = exports.HelloServiceService = {
  sayHello: {
    path: '/HelloService/SayHello',
    requestStream: false,
    responseStream: false,
    requestType: hello_service_pb.HelloRequest,
    responseType: hello_service_pb.HelloReply,
    requestSerialize: serialize_HelloRequest,
    requestDeserialize: deserialize_HelloRequest,
    responseSerialize: serialize_HelloReply,
    responseDeserialize: deserialize_HelloReply,
  },
};

exports.HelloServiceClient = grpc.makeGenericClientConstructor(HelloServiceService);
