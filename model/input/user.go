package input

type UserRegisterInput struct {
	FullName string `json:"full_name" valid:"required"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
	Role     string `json:"role" valid:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

type UserUpdateInput struct {
	FullName string `json:"full_name" `
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" `
	Role     string `json:"role"`
}

type UserUpdateID struct {
	ID int `uri:"id" valid:"required"`
}

type UserDeleteID struct {
	ID int `uri:"id" valid:"required"`
}
