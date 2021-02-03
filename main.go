package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money %d", u.Name, u.Age, u.Money)
}

func (u *User) setName(name string) {
	u.Name = name
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{Name: "Bob", Age: 25, Money: -50, Avg_grades: 84.5, Happiness: 0.8}
	//bob.setName("Not Bob")
	//fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page.")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about/", about_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User =
	//bob := User{name: "Bob",age: 25, money: -50, avg_grades: 84.5, happiness: 0.8}

	handleRequest()
}
