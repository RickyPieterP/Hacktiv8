package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup
	userSvc := NewUserService(db)

	names := []string{"nama1", "nama2", "nama3", "nama4", "nama5"}

	wg.Add(len(names))
	for _, n := range names {
		go func(name string) {
			res := userSvc.Register(&User{Nama: name})
			fmt.Println(res)
			wg.Done()
		}(n)
	}
	wg.Wait()

	resGetUser := userSvc.GetUser()

	fmt.Println("-----------Hasil get user-------------")
	for _, v := range resGetUser {
		fmt.Println(v.Nama)
	}

	// tampung := userSvc.Print()
	// fmt.Println(tampung)
}
