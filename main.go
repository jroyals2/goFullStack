package main

import (
	"fmt"
	"github.com/gorilla/mux"
	c "github.com/jroyals2/firstFullStackGo/controllers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items", c.GetAllCall).Methods("GET")
	r.HandleFunc("/items", c.CreateOneCall).Methods("POST")
	r.HandleFunc("items", c.UpdateOneCall).Methods("PUT")
	r.HandleFunc("/items", c.DeleteOneCall).Methods("DELETE")
	r.HandleFunc("/items/{id}", c.ShowOneCall).Methods("GET")

	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on port 3001")
}
