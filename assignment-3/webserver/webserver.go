package webserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
)

const port = ":8081"

type Status struct {
	Water int `json:water`
	Wind  int `json:water`
}

func Start() {
	http.HandleFunc("/", handler)

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := getData()
	userAgent := r.Header.Get("x-client")
	if userAgent != "API" {
		tpl, err := template.ParseFiles("./webserver/template2.html")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		tpl.Execute(w, map[string]interface{}{
			"Payload": data,
			"Title":   "Assignment 3",
		})
	}
}

func getData() (out string) {
	// Write JSON
	writeJSON()

	// Read JSON
	data := readJSON()

	water := data.Water
	wind := data.Wind

	if water < 5 {
		out = "aman"
	} else if water >= 6 && water <= 8 {
		out = "siaga"
	} else if water > 8 {
		out = "bahaya"
	}

	if wind < 6 {
		out = "aman"
	} else if wind >= 7 && wind <= 15 {
		out = "siaga"
	} else if wind > 15 {
		out = "bahaya"
	}

	return out
}

func readJSON() Status {
	var data Status

	file, _ := ioutil.ReadFile("./webserver/value.json")
	json.Unmarshal(file, &data)

	return data
}

func writeJSON() {
	var data Status

	water := rand.Intn(100)
	wind := rand.Intn(100)

	data.Water = water
	data.Wind = wind

	jsonString, _ := json.Marshal(data)
	ioutil.WriteFile("./webserver/value.json", jsonString, os.ModePerm)
}
