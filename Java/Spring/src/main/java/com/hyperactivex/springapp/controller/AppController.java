package com.hyperactivex.springapp.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.ArrayList;
import java.util.Collection;

@Controller
public class AppController {
    @RequestMapping(value="/", method = RequestMethod.GET)
    public @ResponseBody String SayHelloWorld() {
        try {
            return "Hello World!";
        } catch (Exception e) {
            ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR);
            return "Something wrong!";
        }
    }
}
