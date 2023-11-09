package seed

import (
	"project4/model/entity"
)

var User = entity.User{
	ID:       1,
	FullName: "admin",
	Email:    "admin@gmail.com",
	Password: "admin",
	Role:     "admin",
	Balance:  0,
}
