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

	"time"
)

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
	foxId := vars["foxId"]

	if FoxExists(foxId){
		DeleteFoxFromStorage(foxId)
	}

	addFoxToStorage(w, r, http.StatusAccepted, foxId)

	if err := json.NewEncoder(w).Encode(Error{Code:http.StatusAccepted, Message:fmt.Sprint("Fox %s updated", foxId)}); err != nil{
		panic(err)
	}

}

func DeleteFox(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	foxId := vars["foxId"]
	if !FoxExists(foxId){
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(Error{Code:http.StatusNotFound, Message:fmt.Sprint("Fox %s not found", foxId)}); err != nil{
			panic(err)
		}
		return
	}
	DeleteFoxFromStorage(foxId)
	w.WriteHeader(http.StatusOK)

}

func ShowStats(w http.ResponseWriter, r *http.Request){
	s := Statistics{
		TimeSinceLastNOK:int64(time.Since(timeOfLastNOK)/time.Millisecond),
		TimeSinceLastOK:int64(time.Since(timeOfLastOK)/time.Millisecond),
		ParallelRequestCount:parallelRequestCount,
		NodeName:nodeName}

	if err := json.NewEncoder(w).Encode(s); err != nil{
		panic(err)
	}
}