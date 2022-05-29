package models

type Users struct {
	UserID   int    `json:"userID"`
	Username string `json:"userName"`
	Password string `json:"password"`
	UserType int    `json:"userType"`
}
