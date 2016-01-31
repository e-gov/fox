package main

type Fox struct {
	Name 		string	`json:"name"`
	Parents 	[]string `json:"parents"`
	Uuid		string 	`json:"uuid"`
}

type Foxes []Fox