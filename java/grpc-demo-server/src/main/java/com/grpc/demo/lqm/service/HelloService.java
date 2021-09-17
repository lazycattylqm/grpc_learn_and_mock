package com.grpc.demo.lqm.service;

import io.grpc.stub.StreamObserver;
import lqm.demo.grpc.server.hello.HelloServiceGrpc;
import lqm.demo.grpc.server.hello.HelloServiceOuterClass;

public class HelloService extends HelloServiceGrpc.HelloServiceImplBase {
    @Override
    public void sayHello(HelloServiceOuterClass.HelloRequest request,
                         StreamObserver<HelloServiceOuterClass.HelloReply> responseObserver) {
        final String name = request.getName();
        System.out.printf("request name is %s%n", name);
        final String helloMsg = String.format("Hello %s, this is java server", name);
        final HelloServiceOuterClass.HelloReply reply =
                HelloServiceOuterClass.HelloReply.newBuilder().setMessage(helloMsg).build();
        responseObserver.onNext(reply);
        responseObserver.onCompleted();
    }
}
