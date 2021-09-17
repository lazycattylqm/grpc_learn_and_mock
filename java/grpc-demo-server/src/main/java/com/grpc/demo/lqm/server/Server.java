package com.grpc.demo.lqm.server;

import com.grpc.demo.lqm.service.HelloService;
import io.grpc.ServerBuilder;

import java.io.IOException;
import java.util.concurrent.TimeUnit;

public class Server {

    private static final int PORT = 3000;

    private io.grpc.Server grpcServer;

    public void startServer() throws IOException {
        grpcServer = ServerBuilder.forPort(PORT).addService(new HelloService()).build().start();
        System.out.printf("server start, listening port %d%n", PORT);
        Runtime.getRuntime().addShutdownHook(new Thread() {
            @Override
            public void run() {
                System.err.println("*** shutting down gRPC server since JVM is shutting down");
                try {
                    stopServer();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                System.err.println("*** server shut down");
            }
        });
    }

    private void stopServer() throws InterruptedException {
        if (grpcServer != null) {
            grpcServer.shutdown().awaitTermination(1, TimeUnit.MINUTES);
        }
    }

    public void blockUntilShutdown() throws InterruptedException {
        if (grpcServer != null) {
            grpcServer.awaitTermination();
        }
    }
}
