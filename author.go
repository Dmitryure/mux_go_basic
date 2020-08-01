package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Author data
type Author struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

// AuthorEndpoint sends authors in json
func AuthorEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	json.NewEncoder(res).Encode(authors)
}

// AuthorByIdEndpoint returns authod by id
func AuthorByIdEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	params := mux.Vars(req)
	for _, author := range authors {
		if author.ID == params["id"] {
			json.NewEncoder(res).Encode(author)
			return
		}
	}
	json.NewEncoder(res).Encode(Author{})
}

func DeleteAuthor(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	params := mux.Vars(req)
	for i, author := range authors {
		if author.ID == params["id"] {
			authors = append(authors[:i], authors[:i+1]...)
		}
	}
}

// ChangeAuthor changes author
func ChangeAuthor(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	params := mux.Vars(req)
	var changes Author
	json.NewDecoder(req.Body).Decode(&changes)
	for i, author := range authors {
		if author.ID == params["id"] {
			if changes.Firstname != "" {
				author.Firstname = changes.Firstname
			}
			if changes.Lastname != "" {
				author.Lastname = changes.Lastname
			}
			if changes.Password != "" {
				author.Password = changes.Password
			}
			if changes.Username != "" {
				author.Username = changes.Username
			}
			authors[i] = author
			json.NewEncoder(res).Encode(author)
			return
		}
	}
	json.NewEncoder(res).Encode(Author{})
}
