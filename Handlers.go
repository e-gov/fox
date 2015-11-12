package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"io/ioutil"
	"io"
	"log"

	"code.google.com/p/go-uuid/uuid"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request){
	todos := Todos{
		Todo{Name: "First thing to do"},
		Todo{Name: "Second thing to do"},
	}

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func FoxShow(w http.ResponseWriter, r *http.Request){
	var fox Fox
	vars := mux.Vars(r)

	// Read the fox from storage
	fox, err := ReadFox(vars["foxId"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Translate any storage error to a basic 404
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(Error{Code:404, Message:"Fox not found"}); err != nil{
			panic(err)
		}
		return
	}

	// So we got the fox, now attempt returning it
	if err:= json.NewEncoder(w).Encode(fox); err != nil{
		log.Print("Error")
		panic(err)
	}
}

func addFoxToStorage(w http.ResponseWriter, r *http.Request, status int, uuid string){
	var fox Fox
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil{
		panic(err)
	}

	if err := r.Body.Close(); err != nil{
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.Unmarshal(body, &fox); err != nil{
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(Error{Code:422, Message:err.Error()}); err != nil{
			panic(err)
		}
	}

	w.WriteHeader(status)

	if err:= json.NewEncoder(w).Encode(StoreFox(fox, uuid)); err != nil{
		log.Print("Error")
		panic(err)
	}

}

func AddFox(w http.ResponseWriter, r *http.Request){
	addFoxToStorage(w, r, http.StatusCreated, uuid.New())
}

func UpdateFox(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	if FoxExists(vars["foxId"]){
		DeleteFoxFromStorage(vars["foxId"])
	}

	addFoxToStorage(w, r, http.StatusAccepted, vars["foxId"])
}
