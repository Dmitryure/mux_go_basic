package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var authors []Author = []Author{
	Author{"1", "Dima", "Utkin", "privet", "pass"},
	Author{"2", "Vasya", "Familia", "poka", "asdads"},
}

func rootEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	res.Write([]byte(`{ "message": "Hello" }`))
}

func main() {
	// data, _ := json.Marshal(authors)
	router := mux.NewRouter()
	router.HandleFunc("/", rootEndpoint).Methods("GET")
	router.HandleFunc("/authors", AuthorEndpoint).Methods("GET")
	router.HandleFunc("/author/{id}", AuthorByIdEndpoint).Methods("GET")
	router.HandleFunc("/author/{id}/delete", DeleteAuthor).Methods("GET")
	router.HandleFunc("/author/{id}", ChangeAuthor).Methods("PUT")
	http.ListenAndServe(":3000", router)
}
