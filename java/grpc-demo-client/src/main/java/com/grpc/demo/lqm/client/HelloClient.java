package com.grpc.demo.lqm.client;

import io.grpc.Channel;
import lqm.demo.grpc.server.hello.HelloServiceGrpc;
import lqm.demo.grpc.server.hello.HelloServiceOuterClass;

public class HelloClient {
    private HelloServiceGrpc.HelloServiceBlockingStub clientStub;

    public void init(Channel channel) {
        clientStub = HelloServiceGrpc.newBlockingStub(channel);
    }

    public void sayHello(String name) {

        final HelloServiceOuterClass.HelloRequest helloRequest =
                HelloServiceOuterClass.HelloRequest.newBuilder().setName(name).build();
        final HelloServiceOuterClass.HelloReply reply = clientStub.sayHello(helloRequest);
        System.out.printf("message from server is %s%n", reply.getMessage());
    }
}
