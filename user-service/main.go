package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "user-service/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")

    log.Println("User Service is running on port 9000")
    log.Fatal(http.ListenAndServe(":9000", r))
}
