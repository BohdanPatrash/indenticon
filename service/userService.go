package service

import (
	"fmt"

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

func AddUser(email string) {
	db := repo.DB()
	var hash Hash = GetHash(email)
	fmt.Println("Email: ", email)
	fmt.Println("Hash: ", hash.String())
	user := dto.User{
		Email: email,
		Hash:  hash,
	}
	fmt.Println("user: ", user)
	db.Users[hash.String()] = user
}

func GetByEmail(email string) dto.User {
	return GetByHash(GetHash(email))
}

func GetByHash(hash []byte) dto.User {
	var h Hash = hash[:]
	db := repo.DB()

	return db.Users[h.String()]
}
