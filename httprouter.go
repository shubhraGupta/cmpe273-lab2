package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"

	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

type Resp struct {
	Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))

}

func helloindex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var u User

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)
	//fmt.Fprintf(w, "%s", u)

	// log.Println(t.Test)

	// Add an Id
	s := u.Name

	// Marshal provided interface into JSON structure
	rs := Resp{}
	rs.Greeting = "Hello " + s + "!"

	uj, _ := json.Marshal(rs)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s\n", uj)

}

func main() {

	mux := httprouter.New()

	mux.GET("/hello/:name", hello)
	mux.POST("/hello", helloindex)

	server := http.Server{

		Addr: "0.0.0.0:8080",

		Handler: mux,
	}

	server.ListenAndServe()

}
