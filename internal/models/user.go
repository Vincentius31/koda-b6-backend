package models

import "time"

type User struct {
	IDUser         int       `json:"id_user"`
	RolesID        *int      `json:"roles_id"`
	Fullname       string    `json:"fullname"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Address        *string   `json:"address"`
	Phone          *string   `json:"phone"`
	ProfilePicture *string   `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
    Token  string `json:"token"`
    RoleID *int   `json:"role_id"` 
}
