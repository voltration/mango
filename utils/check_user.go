package utils

import (
	"errors"
	"mango/db"

	"gorm.io/gorm"
)

func CheckIfUserExistsByEmail(email string) (db.User, error) {
	var user db.User

	result := db.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("email does not exist")
		}
		return user, result.Error
	}

	return user, nil
}
