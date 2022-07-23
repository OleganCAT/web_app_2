package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name                string
	Age                 uint16
	money               int16
	avGrades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal %d", u.Name, u.Age, u.money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Alex", 25, -50, 4.2, 0.7}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInfo())
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
