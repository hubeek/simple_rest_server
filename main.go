package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Person struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

var people []Person

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/" {
			getPeople(w, r)
		} else {
			getPerson(w, r)
		}
	case "POST":
		addPerson(w, r)
	case "PUT":
		updatePerson(w, r)
	case "DELETE":
		deletePerson(w, r)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	for _, person := range people {
		if person.Id == id {
			json.NewEncoder(w).Encode(person)
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	var updatedPerson Person
	json.NewDecoder(r.Body).Decode(&updatedPerson)

	for i, person := range people {
		if person.Id == updatedPerson.Id {
			people[i] = updatedPerson
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	var deletedPerson Person
	json.NewDecoder(r.Body).Decode(&deletedPerson)

	newPeople := make([]Person, 0)
	for _, person := range people {
		if person.Id != deletedPerson.Id {
			newPeople = append(newPeople, person)
		}
	}
	people = newPeople

	json.NewEncoder(w).Encode(people)
}
