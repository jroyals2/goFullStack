package controllers

import (
	"fmt"
	"net/http"
)

//GetAllCall will get all the items and return it to the front end
func GetAllCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not Done Yet")
}

//ShowOneCall will get one item and show it to the front end
func ShowOneCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not Done Yet")
}

//CreateOneCall will post a new item to the database and send the new updated list of all to the front end
func CreateOneCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not done yet")
}

//UpdateOneCall will grab one Item and edit it and return the to the show page
func UpdateOneCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not done yet")
}

//DeleteOneCall will delete one item and return the all page
func DeleteOneCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Not Done Yet")
}
