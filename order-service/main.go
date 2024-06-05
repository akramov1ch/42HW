package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "order-service/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/order/notify", handlers.NotifyOrder).Methods("POST")

    log.Println("Order Service is running on port 9001")
    log.Fatal(http.ListenAndServe(":9001", r))
}
