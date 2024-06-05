package utils

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
    "user-service/models"
)

func SendNotification(notification models.Notification) {
    url := "http://localhost:9001/order/notify"
    jsonValue, _ := json.Marshal(notification)
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        log.Printf("Failed to send notification: %v", err)
        return
    }
    defer resp.Body.Close()
    log.Println("Notification sent successfully")
}
