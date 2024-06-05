package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "order-service/models"
)

func NotifyOrder(w http.ResponseWriter, r *http.Request) {
    var notification models.Notification
    json.NewDecoder(r.Body).Decode(&notification)

    log.Printf("Received notification: User %s %s", notification.UserID, notification.Action)
    w.WriteHeader(http.StatusOK)
}
