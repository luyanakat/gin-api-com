package token

import (
	"gin-api/ent"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, user *ent.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}
func CheckPassword(providedPassword string, user *ent.User) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
