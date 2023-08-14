package main

import (
    "net/http"
    "./user"
)

func main() {
    userService := &user.UserService{}
    userHandler := &user.UserHandler{
        Service: userService,
    }

    http.HandleFunc("/users/create", userHandler.CreateUserHandler)

    http.ListenAndServe(":8080", nil)
}
