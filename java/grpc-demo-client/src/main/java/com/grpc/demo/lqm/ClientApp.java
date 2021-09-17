package com.grpc.demo.lqm;

import com.grpc.demo.lqm.client.HelloClient;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class ClientApp {
    private static final String TARGET = "localhost:3000";
    private static final String NAME = "liqiming java";

    public static void main(String[] args) {
        final ManagedChannel channel = ManagedChannelBuilder.forTarget(TARGET).usePlaintext().build();
        final HelloClient helloClient = new HelloClient();
        helloClient.init(channel);
        helloClient.sayHello(NAME);
    }
}
