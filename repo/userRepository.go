package repo

import (
	"github.com/BohdanPatrash/indenticon/dto"
)

func SaveUser(user *dto.User) error {
	err := db.Insert(user)
	return err
}

func GetAllUsers() ([]dto.User, error) {
	var users []dto.User = make([]dto.User, 0)

	err := db.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}
