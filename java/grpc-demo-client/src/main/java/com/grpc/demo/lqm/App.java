package com.grpc.demo.lqm;

import com.grpc.demo.lqm.client.HelloClient;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class App {
    private static final String TARGET = "localhost:3000";
    private static final String NAME = "liqiming";

    public static void main(String[] args) {
        final ManagedChannel channel = ManagedChannelBuilder.forTarget(TARGET).usePlaintext().build();
        final HelloClient helloClient = new HelloClient();
        helloClient.init(channel);
        helloClient.sayHello(NAME);
    }
}
