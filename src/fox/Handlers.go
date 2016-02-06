package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"io/ioutil"
	"io"
	"log"

	"github.com/pborman/uuid"

	"time"
)

func sendHeaders(w http.ResponseWriter){
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")

}
func FoxShow(w http.ResponseWriter, r *http.Request){
	var fox Fox
	vars := mux.Vars(r)

	// Read the fox from storage
	fox, err := ReadFox(vars["foxId"])

    sendHeaders(w)    
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
		log.Print("Error encoding the fox")
		panic(err)
	}
}

func FoxList(w http.ResponseWriter, r *http.Request)  {
    var foxes []Fox
    
    foxes, err := GetFoxes()
    log.Printf("Found %d foxes", len(foxes))

    sendHeaders(w)    
    if err != nil{
	   w.WriteHeader(http.StatusNotFound)
	   if err := json.NewEncoder(w).Encode(Error{Code:404, Message:"Fox list not available"}); err != nil{
		  panic(err)
	   }
	   return
    }

    if err := json.NewEncoder(w).Encode(foxes); err != nil {
	log.Print("Error encoding fox list")
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

    sendHeaders(w)    
    st := string(body[:110])
    log.Print(st)

	if err := json.Unmarshal(body, &fox); err != nil{
	w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(Error{Code:422, Message:err.Error()}); err != nil{
			panic(err)
		}
	return
	}

	w.WriteHeader(status)

    s := StoreFox(fox, uuid)
    log.Print(s)
	if err:= json.NewEncoder(w).Encode(s); err != nil{
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
		log.Println("Deleting prior version of fox " + foxId)
		addFoxToStorage(w, r, http.StatusAccepted, foxId)

		if err := json.NewEncoder(w).Encode(Error{Code:http.StatusAccepted, Message:fmt.Sprint("Fox %s updated", foxId)}); err != nil{
			panic(err)
		}

	}else{
		// If the fox does not exist, go for the creation instead
		log.Println("New fox, creating with uuid " + foxId)
		AddFox(w, r)
	}


}

func DeleteFox(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

    sendHeaders(w)    
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
