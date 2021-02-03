package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money %d", u.name, u.age, u.money)
}

func (u *User) setName(name string) {
	u.name = name
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{name: "Bob", age: 25, money: -50, avg_grades: 84.5, happiness: 0.8}
	bob.setName("Not Bob")
	fmt.Fprintf(w, bob.getAllInfo())
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
