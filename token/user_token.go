package token

import (
	"errors"
	"gin-api/ent"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	if password == "" {
		return nil, errors.New("password is empty")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
func CheckPassword(providedPassword string, user *ent.User) error {
	log.Println("provided", providedPassword)
	log.Println("user in db", user.Password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
