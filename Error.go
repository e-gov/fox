package main

type Error struct {
	Code 		int	`json:"code"`
	Message 	string  `json:"message"`
	fields		string 	`json:"fields"`
}

