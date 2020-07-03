package service

import (
	"github.com/BohdanPatrash/indenticon/dto"

	"github.com/BohdanPatrash/indenticon/repo"
)

func GetAll() []dto.User {
	var users []dto.User = make([]dto.User, 0)
	for _, val := range repo.DB().Users {
		users = append(users, val)
	}
	return users
}

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

func SaveUser(user *dto.User) {
	db := repo.DB()
	var hash Hash = user.Hash
	db.Users[hash.String()] = *user
}

func GetByEmail(email string) dto.User {
	return GetByHash(GetHash(email))
}

func GetByHash(hash []byte) dto.User {
	var h Hash = hash[:]
	db := repo.DB()

	return db.Users[h.String()]
}
