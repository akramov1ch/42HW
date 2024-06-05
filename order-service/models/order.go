package models

type Notification struct {
    UserID string `json:"user_id"`
    Action string `json:"action"`
}
