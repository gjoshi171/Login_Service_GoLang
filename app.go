package main

import (
	. "GoMysql/models"
	. "GoMysql/mongoDb"
	. "GoMysql/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePersonEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		return
	}
	var personDb []*Person
	var response Response
	var result Result
	var username string = person.GetUsername()
	//var password string = person.GetPassword()
	//personsDb = FindAll(username, password)
	personDb = FindByUserName(username)

	//CheckPasswordHash(personsDb[0].GetPassword(), person.GetPassword())

	if len(personDb) > 0 && person.GetEmail() == personDb[0].GetEmail() && person.GetUsername() == personDb[0].GetUsername() && CheckPasswordHash(personDb[0].GetPassword(), person.GetPassword()) == true {

		response.SetMessage("already existed")
		ResponseWithError(w, http.StatusBadRequest, response)
	} else {
		Insert(person)

		result.SetPerson(&person)

		response.SetResult(&result)
		response.SetMessage("success")
		ResponseWithJSON(w, http.StatusOK, "success", response)
	}

}

func findPersonEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person []*Person

	person = FindByUserName(params["userName"])

	fmt.Println("The object  is ", person)
	ResponseWithJSON(w, http.StatusOK, "success", person)
}

func login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		return
	}

	var username string = person.GetUsername()
	var personDb []*Person
	personDb = FindByUserName(username)

	if len(personDb) > 0 && personDb[0].GetUsername() == person.GetUsername() && CheckPasswordHash(personDb[0].GetPassword(), person.GetPassword()) == true {
		ResponseWithJSON(w, http.StatusOK, "successfully login", personDb[0])
	} else {
		ResponseWithError(w, http.StatusBadRequest, "login unsuccessful")
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/findPerson/{userName}", findPersonEndPoint).Methods("GET")

	r.HandleFunc("/createPerson", CreatePersonEndPoint).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	http.ListenAndServe(":3000", r)
}
