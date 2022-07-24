package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                string
	Age                 uint16
	Money               int16
	AvGrades, Happiness float64
	Hobbies             []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.7, []string{"Football", "Skate", "Dance"}}
	tmpl, _ := template.ParseFiles("templates/homePage.html")
	tmpl.Execute(w, bob)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handelRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handelRequest()
}
