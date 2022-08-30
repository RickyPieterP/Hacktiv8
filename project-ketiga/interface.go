package main

import (
	"fmt"
)

type User struct {
	Nama string
}

type userService struct {
	db []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
	// Print() string
}

func NewUserService(db []*User) UserIface {
	return &userService{
		db: db,
	}
}

func (u *userService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + "berhasil didaftarkan"
}

func (u *userService) GetUser() []*User {
	return u.db
}

// func (u *userService) Print() string {
// 	a := "ini return var"

// 	return a
// }

func main() {
	var db []*User
	userSvc := NewUserService(db)

	resRegister := userSvc.Register(&User{Nama: "nama1 "})
	fmt.Println(resRegister)

	resRegister = userSvc.Register(&User{Nama: "nama2 "})
	fmt.Println(resRegister)

	resGetUser := userSvc.GetUser()

	fmt.Println("-----------Hasil get user-------------")
	for _, v := range resGetUser {
		fmt.Println(v.Nama)
	}

	// tampung := userSvc.Print()
	// fmt.Println(tampung)
}
