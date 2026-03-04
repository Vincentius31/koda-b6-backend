package models

type User struct {
	IDUser         int     `json:"id_user"`
	RolesID        *int    `json:"roles_id"`
	Fullname       string  `json:"fullname"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	Address        *string `json:"address"`
	Phone          *string `json:"phone"`
	ProfilePicture *string `json:"profile_picture"`
}

type UserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
