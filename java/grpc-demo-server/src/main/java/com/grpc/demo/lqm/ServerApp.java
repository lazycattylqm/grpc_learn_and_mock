package com.grpc.demo.lqm;

import com.grpc.demo.lqm.server.Server;

import java.io.IOException;

public class ServerApp {
    public static void main(String[] args) {
        final Server server = new Server();
        try {
            server.startServer();
            server.blockUntilShutdown();
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }
}
