package com.example.controller;

import com.example.service.GrpcClientService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class GrpcClientController {
    @Autowired
    private GrpcClientService clientService;

    @GetMapping("/client/server")
    public String callServer() {
        return clientService.sendMessage();
    }
}
