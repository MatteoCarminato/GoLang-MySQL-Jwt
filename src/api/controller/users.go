package controller

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("list users"))
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("An user"))
}
func CreateUsers(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Create users"))
}
func UpdateUsers(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Update users"))
}
func DeleteUsers(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Delete users"))
}