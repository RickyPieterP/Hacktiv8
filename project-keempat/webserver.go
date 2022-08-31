package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var PORT = ":8080"

type User struct {
	Nama string `json:"nama"`
}

type userService struct {
	db []*User
}

type UserIface interface {
	Register(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	// Print() string
}

func NewUserService(db []*User) UserIface {
	return &userService{
		db: db,
	}
}

func (u *userService) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)

		var user User
		err := decoder.Decode(&user)

		if err != nil {
			fmt.Fprint(w, "error")
		}

		u.db = append(u.db, &User{Nama: user.Nama})

		// print := fmt.Sprintf(user.Nama + " berhasil didaftarkan")
		v, _ := json.Marshal(user)

		w.Write(v)
	}
}

func (u *userService) GetUser(w http.ResponseWriter, r *http.Request) {

	v, _ := json.Marshal(u.db)

	// fmt.Fprint(w, v)
	w.Write(v)
}

func main() {
	var db []*User
	userSvc := NewUserService(db)

	// http.HandleFunc("/", greet)
	http.HandleFunc("/register", userSvc.Register)
	http.HandleFunc("/getuser", userSvc.GetUser)

	http.ListenAndServe(PORT, nil)
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	msg := "Hello World"
// 	fmt.Fprint(w, msg)
// }
