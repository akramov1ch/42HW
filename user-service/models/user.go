package models

type User struct {
    ID   string `json:"id"`
    Data string `json:"data"`
}

type Notification struct {
    UserID string `json:"user_id"`
    Action string `json:"action"`
}
