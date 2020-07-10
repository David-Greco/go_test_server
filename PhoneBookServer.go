package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//Every entry for the phonebook
type Entry struct {
	name   string
	number int
}

//Used to easily make the slice
type PHONEBOOK []Entry

var phonebook = make(PHONEBOOK, 0)

func main() {

	http.HandleFunc("/phonebook/list", list)
	http.HandleFunc("/phonebook/add/", add)
	http.HandleFunc("/phonebook/delete/", delete)
	http.HandleFunc("/phonebook/lookup/", lookup)

	http.ListenAndServe(":8080", nil)
}

//Follows format http://localhost:8080/phonebook/list
func list(w http.ResponseWriter, r *http.Request) {

	for x := range phonebook {
		user := phonebook[x]
		fmt.Fprintf(w, "Name: %s\nNumber: %d\n\n", user.name, user.number)
	}
}

//Follows format http://localhost:8080/phonebook/add/name/number
func add(w http.ResponseWriter, r *http.Request) {

	var user Entry
	substring := r.URL.Path[15:]
	name := substring[:strings.IndexByte(substring, '/')]
	namelen := len(name) + 1
	user.name = name
	user.number, _ = strconv.Atoi(substring[namelen:])

	phonebook = append(phonebook, user)

	fmt.Fprintf(w, "Name: %s\nNumber: %d\nAdded to Phonebook.", user.name, user.number)
}

//Follows format http://localhost:8080/phonebook/delete/name/number
func delete(w http.ResponseWriter, r *http.Request) {

	Name := r.URL.Path[18:]

	//temp for loop
	index := 0
	for x := range phonebook {
		if phonebook[x].name == Name {
			index = x
		}
	}

	phonebook = append(phonebook[:index], phonebook[index+1:]...)

	fmt.Fprintf(w, "%s deleted from Phonebook.", Name)
}

//Follows format http://localhost:8080/phonebook/lookup/name
func lookup(w http.ResponseWriter, r *http.Request) {

	Name := r.URL.Path[18:]

	for x := range phonebook {
		user := phonebook[x]
		if user.name == Name {
			fmt.Fprintf(w, "Name: %s\nNumber: %d\n\n", user.name, user.number)
			return
		}
	}
	fmt.Fprintf(w, "Not Found")
}
