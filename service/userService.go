package service

import (
	"github.com/BohdanPatrash/indenticon/dto"
)

func CreateUser(email string) *dto.User {
	var hash Hash = GetHash(email)
	indenticon, err := CreateIndenticon(hash)
	if err != nil {
		return nil
	}
	user := &dto.User{
		Email:      email,
		Hash:       hash,
		Indenticon: indenticon,
	}
	return user
}
