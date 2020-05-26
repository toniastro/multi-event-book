package utils

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/ichtrojan/thoth"
	"net/http"
)

type Error struct{
	Status int `json:"status"`
	Message string `json:"message,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type A_response struct{
	Status  int `json:"status"`
	Message string `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	json.NewEncoder(w).Encode(data)
}

func Message(status int, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Errors (w http.ResponseWriter, status int, message string, errors interface{} )  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	error := &Error{Status: status, Message: message, Errors: errors}
	response, _ := json.Marshal(error)
	w.Write([]byte(response))
}

func Response (w http.ResponseWriter, status int, message string, data interface{} )  {
	file, _ := thoth.Init("log")
	resp := &A_response{Status: status, Message: message, Data: data }
	response, err := json.Marshal(resp)
	if err != nil {
		file.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
