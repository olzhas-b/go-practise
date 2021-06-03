package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Id      string
	Hobbies []string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "go is a super easy")
	bob := &User{Name: "Bob", Id: "1222222", Hobbies: []string{"football", "csgo", "basketball"}}
	templ, _ := template.ParseFiles("templates/home_page.html")
	templ.Execute(w, bob)
}

func contact_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it is contact page")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contact/", contact_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
