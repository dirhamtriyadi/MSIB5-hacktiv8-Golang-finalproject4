package seed

import (
	"project4/model/entity"

	"golang.org/x/crypto/bcrypt"
)

var passwordHash, _ = bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)

var User = entity.User{
	ID:       1,
	FullName: "admin",
	Email:    "admin@gmail.com",
	Password: string(passwordHash),
	Role:     "admin",
	Balance:  0,
}
