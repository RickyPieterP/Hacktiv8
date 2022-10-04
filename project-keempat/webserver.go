package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var PORT = "8080"

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

	r := mux.NewRouter()

	// http.HandleFunc("/", greet)
	r.HandleFunc("/register", userSvc.Register)
	r.HandleFunc("/getuser", userSvc.GetUser)

	srv := &http.Server{
		Addr: "0.0.0.0:" + PORT,

		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	msg := "Hello World"
// 	fmt.Fprint(w, msg)
// }
