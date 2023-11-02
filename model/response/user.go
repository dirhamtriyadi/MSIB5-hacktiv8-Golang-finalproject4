package response

import "time"

type UserRegisterResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDeleteResponse struct {
	Message string `json:"message"`
}
