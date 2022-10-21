package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"project-pertama/sesi-10/server/controller"
)

type Router struct {
	port string
	user *controller.UserController
}

func NewRouter(port string, user *controller.UserController) *Router {
	return &Router{
		port: port,
		user: user,
	}
}

func (r *Router) Start() {
	mux := http.NewServeMux()
	
	mux.Handle("/user/ping", middleware1(middleware2(r.user.Ping)))
	log.Println("Server is running at port", r.port)
	http.ListenAndServe(r.port, mux)
}

func middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware 1 is started...")
		ctx := r.Context()
		ctx = context.WithValue(ctx, "role", "admin")
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
		log.Println("middleware 1 is ended...")
	}
}

func middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware 2 is started...")
		if true {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "unauthorized",
			})
			return
		}
		log.Println("access is ", r.Context().Value("role"))
		next.ServeHTTP(w, r)
		log.Println("")
		log.Println("middleware 2 is ended...")
	}
}
