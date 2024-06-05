package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "user-service/models"
    "user-service/utils"
)

var users = make(map[string]string)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    
    if _, exists := users[user.ID]; exists {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("User already exists"))
        return
    }

    users[user.ID] = user.Data

    // Bildirishnoma yuborish
    notification := models.Notification{
        UserID: user.ID,
        Action: "created",
    }
    utils.SendNotification(notification)

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User created"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    if _, exists := users[id]; !exists {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("User not found"))
        return
    }

    users[id] = user.Data
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User updated"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    if _, exists := users[id]; !exists {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("User not found"))
        return
    }

    delete(users, id)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User deleted"))
}
