package seed

import (
	"project4/model/entity"

	"golang.org/x/crypto/bcrypt"
)

var passwordHash, _ = bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)

var User = entity.User{
	FullName: "superadmin",
	Email:    "superadmin@gmail.com",
	Password: string(passwordHash),
	Role:     "superadmin",
	Balance:  0,
}
